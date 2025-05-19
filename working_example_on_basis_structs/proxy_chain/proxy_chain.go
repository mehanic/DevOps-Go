package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// Constants for API key and SSH key name
const KEY_NAME = "proxy"
const API_KEY = ""

type Droplet struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type DropletNetworks struct {
	IPv4 []struct {
		IPAddress string `json:"ip_address"`
	} `json:"v4"`
}

type DropletResp struct {
	Droplet struct {
		ID       int             `json:"id"`
		Name     string          `json:"name"`
		Status   string          `json:"status"`
		Networks DropletNetworks `json:"networks"`
	} `json:"droplet"`
}

// Send API request to DigitalOcean
func send(method, endpoint string, data map[string]interface{}) ([]byte, error) {
	url := "https://api.digitalocean.com/v2/" + endpoint
	client := &http.Client{}

	jsonData, _ := json.Marshal(data)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+API_KEY)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		fmt.Printf("[-] Request Error: %s\n", body)
		return nil, fmt.Errorf("request error")
	}

	if resp.StatusCode >= 500 {
		fmt.Printf("[-] Server Error: %s\n", body)
		return nil, fmt.Errorf("server error")
	}

	return body, nil
}

// Generate random string of specified size
func randomStr(size int) string {
	chars := "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, size)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

// Fetch images available in DigitalOcean
func getImages() ([]string, error) {
	fmt.Println("[*] Getting image list.")
	data, err := send("GET", "images", map[string]interface{}{"type": "distribution"})
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	var images []string
	for _, img := range result["images"].([]interface{}) {
		image := img.(map[string]interface{})
		images = append(images, image["slug"].(string))
	}

	return images, nil
}

// Fetch SSH Key fingerprint by name
func getSSHKey(keyName string) (string, error) {
	fmt.Printf("[*] Getting SSH key %s\n", keyName)
	data, err := send("GET", "account/keys", nil)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return "", err
	}

	for _, s := range result["ssh_keys"].([]interface{}) {
		key := s.(map[string]interface{})
		if key["name"] == keyName {
			return key["fingerprint"].(string), nil
		}
	}

	return "", fmt.Errorf("ssh key not found")
}

// Create a new DigitalOcean droplet
func createDroplet(images []string, fingerprint string) (int, string, error) {
	fmt.Println("[*] Creating new droplet.")

	image := images[rand.Intn(len(images))]
	droplet := map[string]interface{}{
		"name":              randomStr(8),
		"region":            "nyc1", // You can adjust this to a random region from the image if needed
		"size":              "512mb",
		"image":             image,
		"ssh_keys":          []string{fingerprint},
		"backups":           false,
		"ipv6":              false,
		"user_data":         nil,
		"private_networking": nil,
	}

	respData, err := send("POST", "droplets", droplet)
	if err != nil {
		return 0, "", err
	}

	var resp DropletResp
	if err := json.Unmarshal(respData, &resp); err != nil {
		return 0, "", err
	}

	fmt.Printf("[+] Successfully created droplet %d.\n", resp.Droplet.ID)

	// Wait for the droplet to become active
	ip, err := getDropletIP(resp.Droplet.ID)
	if err != nil {
		return 0, "", err
	}

	return resp.Droplet.ID, ip, nil
}

// Get the IP address of a specific droplet
func getDropletIP(dropletID int) (string, error) {
	for {
		fmt.Println("[*] Getting droplet IP")
		respData, err := send("GET", fmt.Sprintf("droplets/%d", dropletID), nil)
		if err != nil {
			return "", err
		}

		var resp DropletResp
		if err := json.Unmarshal(respData, &resp); err != nil {
			return "", err
		}

		if resp.Droplet.Status == "active" {
			return resp.Droplet.Networks.IPv4[0].IPAddress, nil
		}

		fmt.Println("[*] Waiting for droplet to become active.")
		time.Sleep(10 * time.Second)
	}
}

// Delete a DigitalOcean droplet
func deleteDroplet(dropletID int) error {
	fmt.Println("[*] Deleting droplet.")
	_, err := send("DELETE", fmt.Sprintf("droplets/%d", dropletID), nil)
	if err == nil {
		fmt.Printf("[+] Successfully deleted image %d.\n", dropletID)
	}
	return err
}

// Usage information
func usage() {
	fmt.Println(`
	Usage:
		./do_proxy_chains create number droplet_id_file
		./do_proxy_chains delete droplet_id_file
	`)
	os.Exit(1)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 4 && os.Args[1] == "create" {
		numDroplets, _ := strconv.Atoi(os.Args[2])
		dropletFile := os.Args[3]

		images, err := getImages()
		if err != nil || len(images) == 0 {
			fmt.Println("[-] No images found.")
			return
		}

		fingerprint, err := getSSHKey(KEY_NAME)
		if err != nil {
			fmt.Println(err)
			return
		}

		var dropletIDs []string
		for i := 0; i < numDroplets; i++ {
			dropletID, ip, err := createDroplet(images, fingerprint)
			if err != nil {
				fmt.Println(err)
				continue
			}

			dropletIDs = append(dropletIDs, strconv.Itoa(dropletID))
			fmt.Printf("[+] Created droplet with IP: %s\n", ip)
		}

		// Save droplet IDs to a file
		if err := ioutil.WriteFile(dropletFile, []byte(strings.Join(dropletIDs, "\n")), 0644); err != nil {
			fmt.Println("[-] Unable to write droplet IDs:", err)
		}

	} else if len(os.Args) == 3 && os.Args[1] == "delete" {
		dropletFile := os.Args[2]
		fileData, err := ioutil.ReadFile(dropletFile)
		if err != nil {
			fmt.Println("[-] File not found:", err)
			return
		}

		for _, line := range strings.Split(string(fileData), "\n") {
			if id, err := strconv.Atoi(line); err == nil {
				deleteDroplet(id)
			}
		}

		os.Remove(dropletFile)

	} else {
		usage()
	}
}

