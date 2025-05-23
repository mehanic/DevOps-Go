package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// vaultCall handles login and secret retrieval from Vault
func vaultCall(vserv, vact string) (interface{}, error) {
	baseURL := "https://" + vserv + ":8200/v1/"
	var vaultURL string
	var req *http.Request
	var err error
	client := &http.Client{}

	// Login action
	if vact == "login" {
		vaultURL = baseURL + "auth/approle/login"
		roleID := os.Getenv("ROLE_ID")
		secID := os.Getenv("SEC_ID")
		loginData := map[string]string{
			"role_id":   roleID,
			"secret_id": secID,
		}

		jsonData, err := json.Marshal(loginData)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal login data: %v", err)
		}

		req, err = http.NewRequest("POST", vaultURL, bytes.NewBuffer(jsonData))
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %v", err)
		}

		req.Header.Set("Content-Type", "application/json")

	} else if vact == "get" {
		vaultURL = baseURL + "secrets/data/myapp"
		clusterToken := os.Getenv("VAULT_TOKEN") // Replace with how you store cluster_token
		req, err = http.NewRequest("GET", vaultURL, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %v", err)
		}

		req.Header.Set("X-Vault-Token", clusterToken)
	} else {
		return nil, fmt.Errorf("invalid action option")
	}

	// Perform the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("non-200 status code: %d", resp.StatusCode)
	}

	// Handle response based on action
	var respData map[string]interface{}
	err = json.Unmarshal(body, &respData)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response data: %v", err)
	}

	if vact == "login" {
		return respData["auth"].(map[string]interface{})["client_token"], nil
	} else if vact == "get" {
		return respData["data"].(map[string]interface{})["data"], nil
	} else {
		return respData, nil
	}
}

func main() {
	vserv := "your_vault_server"
	vact := "login" // or "get"
	result, err := vaultCall(vserv, vact)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
