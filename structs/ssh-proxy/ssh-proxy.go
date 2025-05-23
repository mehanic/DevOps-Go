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

// Constants for the API key
const apiKey = ""

// Droplet structure to capture the response
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

	var jsonData []byte
	if data != nil {
		jsonData, _ = json.Marshal(data)
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

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
func getImages() ([][2]string, error) {
	fmt.Println("[*] Getting image list.")
	data, err := send("GET", "images", map[string]interface{}{"type": "distribution"})
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	var images [][2]string
	for _, img := range result["images"].([]interface{}) {
		image := img.(map[string]interface{})
		regions := image["regions"].([]interface{})
		regionsStr := make([]string, len(regions))
		for i, region := range regions {
			regionsStr[i] = region.(string)
		}
		images = append(images, [2]string{image["slug"].(string), strings.Join(regionsStr, ",")})
	}

	return images, nil
}

// Fetch SSH Keys
func getSSHKeys() ([][2]string, error) {
	fmt.Println("[*] Getting SSH keys")
	data, err := send("GET", "account/keys", nil)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	var keys [][2]string
	for _, s := range result["ssh_keys"].([]interface{}) {
		key := s.(map[string]interface{})
		keys = append(keys, [2]string{key["name"].(string), key["fingerprint"].(string)})
	}

	return keys, nil
}

// Get Droplet IP
func getDropletIP(dropletID int) (string, error) {
	fmt.Println("[*] Getting droplet IP")
	for {
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

// Create Droplet
func createDroplet(images [][2]string, keys [][2]string) (int, string, string, error) {
	fmt.Println("[*] Creating new droplet.")

	image := images[rand.Intn(len(images))]
	key := keys[rand.Intn(len(keys))]

	regions := strings.Split(image[1], ",")
	droplet := map[string]interface{}{
		"name":              randomStr(8),
		"region":            regions[rand.Intn(len(regions))],
		"size":              "512mb",
		"image":             image[0],
		"ssh_keys":          []string{key[1]},
		"backups":           false,
		"ipv6":              false,
		"user_data":         nil,
		"private_networking": nil,
	}

	respData, err := send("POST", "droplets", droplet)
	if err != nil {
		return 0, "", "", err
	}

	var resp DropletResp
	if err := json.Unmarshal(respData, &resp); err != nil {
		return 0, "", "", err
	}

	fmt.Printf("[+] Successfully created droplet %d.\n", resp.Droplet.ID)
	ip, err := getDropletIP(resp.Droplet.ID)
	if err != nil {
		return 0, "", "", err
	}
	return resp.Droplet.ID, key[0], ip, nil
}

// Delete Droplet
func deleteDroplet(dropletID int) error {
	fmt.Println("[*] Deleting droplet.")
	_, err := send("DELETE", fmt.Sprintf("droplets/%d", dropletID), nil)
	if err == nil {
		fmt.Printf("[+] Successfully deleted droplet %d.\n", dropletID)
	}
	return err
}

// Usage information
func usage() {
	fmt.Println(`
	Usage:
		./do_ssh_proxy create
		./do_ssh_proxy delete droplet_id
	`)
	os.Exit(1)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	if len(os.Args) == 2 && os.Args[1] == "create" {
		images, err := getImages()
		if err != nil || len(images) == 0 {
			fmt.Println("[-] No images found.")
			return
		}

		keys, err := getSSHKeys()
		if err != nil || len(keys) == 0 {
			fmt.Println("[-] No SSH keys found.")
			return
		}

		dropletID, key, ip, err := createDroplet(images, keys)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf(`
The SSH server is ready. You can create an SSH SOCKS proxy using the following
command:

	ssh -Nf -i ~/.ssh/private_key -D 1080 root@%s

where the private key should correspond to the public key used when creating
the DigitalOcean SSH key named %s. Once you have successfully connected to
the SSH server, configure your web browser to use the SOCKS proxy 127.0.0.1
on port 1080.

When you are finished with the proxy server, you can delete the droplet using
the following command:

	./do_ssh_proxy delete %d
`, ip, key, dropletID)

	} else if len(os.Args) == 3 && os.Args[1] == "delete" {
		dropletID, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("[-] Invalid droplet ID.")
			return
		}

		err = deleteDroplet(dropletID)
		if err != nil {
			fmt.Println(err)
		}

	} else {
		usage()
	}
}

