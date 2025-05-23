package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/credentials"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/config"
)

var (
	accessID  = "" // Add your AWS Access Key
	accessKey = "" // Add your AWS Secret Key
	region    = "eu-west-2"
	endpoint  = fmt.Sprintf("s3.%s.amazonaws.com", region)
)

func signAndSendRequest(req *http.Request, body []byte, method string) (*http.Response, error) {
	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Initialize the signer
	creds := credentials.NewStaticCredentialsProvider(accessID, accessKey, "")
	signer := v4.NewSigner(creds)

	// Sign the request
	err = signer.SignHTTP(context.TODO(), req, bytes.NewReader(body), "s3", region, aws.TimeNow())
	if err != nil {
		return nil, fmt.Errorf("failed to sign request: %w", err)
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	return resp, nil
}

func uploadFile(bucket, localPath string) {
	fileData, err := os.ReadFile(localPath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Create the URL for uploading the file to S3
	url := fmt.Sprintf("http://%s/%s/%s", endpoint, bucket, localPath)
	fmt.Println("Uploading file to " + url)

	// Prepare the PUT request for uploading the file
	req, err := http.NewRequest("PUT", url, bytes.NewReader(fileData))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	resp, err := signAndSendRequest(req, fileData, "PUT")
	if err != nil {
		log.Fatalf("Upload request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("Uploaded %s OK\n", localPath)
	} else {
		fmt.Printf("Failed to upload file. Status code: %d\n", resp.StatusCode)
	}
}

func downloadFile(bucket, s3Name string) {
	// Create the URL for downloading the file from S3
	url := fmt.Sprintf("http://%s/%s/%s", endpoint, bucket, s3Name)
	fmt.Println("Downloading file from " + url)

	// Prepare the GET request for downloading the file
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	resp, err := signAndSendRequest(req, nil, "GET")
	if err != nil {
		log.Fatalf("Download request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		file, err := os.Create(s3Name)
		if err != nil {
			log.Fatalf("Failed to create file: %v", err)
		}
		defer file.Close()

		_, err = io.Copy(file, resp.Body)
		if err != nil {
			log.Fatalf("Failed to save file: %v", err)
		}
		fmt.Printf("Downloaded %s OK\n", s3Name)
	} else {
		fmt.Printf("Failed to download file. Status code: %d\n", resp.StatusCode)
	}
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <bucket_name> <file_name>")
		return
	}

	bucket := os.Args[1]
	filename := os.Args[2]

	uploadFile(bucket, filename)
	downloadFile(bucket, filename)
}

