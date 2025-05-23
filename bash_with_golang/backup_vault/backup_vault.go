package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "strings"

    vault "github.com/hashicorp/vault/api"
)

var (
    failedExports []string
    outputFile    = "output.txt"
)

// Function to recursively traverse through secrets paths to retrieve secrets.
func recurseThroughSecrets(client *vault.Client, pathPrefix string, candidateKeys []string, mountPoint string) {
    for _, candidateValue := range candidateKeys {
        nextIndex := pathPrefix + candidateValue

        // If the entry ends with a '/', we know it's a folder
        if strings.HasSuffix(candidateValue, "/") {
            // List secrets in the current folder
            secretList, err := client.KVv2ListSecrets(nextIndex, mountPoint)
            if err != nil {
                failedExports = append(failedExports, nextIndex)
                continue
            }

            recurseThroughSecrets(client, nextIndex, secretList.Data["keys"].([]interface{}), mountPoint)

        } else {
            // Read the secret data
            secret, err := client.KVv2Get(nextIndex, mountPoint)
            if err != nil {
                failedExports = append(failedExports, nextIndex)
                continue
            }

            // Write the kv put command to the output file
            writeSecretsToFile(secret.Data["data"].(map[string]interface{}), mountPoint, nextIndex)
        }
    }
}

// Write secrets to the output file
func writeSecretsToFile(secrets map[string]interface{}, mountPoint, nextIndex string) {
    f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Failed to open output file: %v", err)
    }
    defer f.Close()

    // Start building the command string
    command := fmt.Sprintf("vault kv put %s/%s", mountPoint, nextIndex)
    for k, v := range secrets {
        command += fmt.Sprintf(" '%s=%v'", k, v)
    }
    command += "\n"

    // Write command to file
    if _, err := f.WriteString(command); err != nil {
        log.Fatalf("Failed to write to output file: %v", err)
    }
}

func main() {
    // Get the VAULT_TOKEN and VAULT_ADDR from environment variables
    vaultToken := os.Getenv("VAULT_TOKEN")
    if vaultToken == "" {
        log.Fatal("VAULT_TOKEN not set.. Goodbye!")
    }

    vaultAddr := os.Getenv("VAULT_ADDR")
    if vaultAddr == "" {
        log.Fatal("VAULT_ADDR not set.. Goodbye!")
    }

    // Create a new Vault client
    config := vault.DefaultConfig()
    config.Address = vaultAddr

    client, err := vault.NewClient(config)
    if err != nil {
        log.Fatalf("Error creating Vault client: %v", err)
    }

    client.SetToken(vaultToken)

    // Ensure authentication
    if client.Token() == "" {
        log.Fatal("Failed to authenticate with Vault.")
    }

    // Get mount point from environment or use 'secret' as default
    mountPoint := os.Getenv("VAULT_MOUNT_POINT")
    if mountPoint == "" {
        mountPoint = "secret"
    }

    // Get initial entries of secrets
    topLevelKeys, err := client.KVv2ListSecrets("", mountPoint)
    if err != nil {
        log.Fatalf("Failed to list secrets: %v", err)
    }

    // Start recursion through secrets
    recurseThroughSecrets(client, "", topLevelKeys.Data["keys"].([]interface{}), mountPoint)

    // Print failed exports
    if len(failedExports) > 0 {
        fmt.Println("Exporting failed on the following, most likely due to permission issues:")
        for _, failed := range failedExports {
            fmt.Println(failed)
        }
    }
}

