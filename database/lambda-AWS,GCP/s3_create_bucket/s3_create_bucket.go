package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const region = "eu-west-2"
const endpoint = "s3.eu-west-2.amazonaws.com"

// Create XML structure for CreateBucketConfiguration
type CreateBucketConfiguration struct {
	XMLName           xml.Name `xml:"CreateBucketConfiguration"`
	Xmlns             string   `xml:"xmlns,attr"`
	LocationConstraint string   `xml:"LocationConstraint"`
}

// Pretty print XML data
func xmlPrettyPrint(data []byte) string {
	var prettyXML bytes.Buffer
	err := xml.Indent(&prettyXML, data, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting XML: %v", err)
	}
	return prettyXML.String()
}

// Create a new S3 bucket with the specified name
func createBucket(bucketName string, accessKeyID, secretAccessKey string) {
	// Initialize a session using AWS credentials
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}

	svc := s3.New(sess)

	// Build CreateBucketConfiguration XML
	xmlData := CreateBucketConfiguration{
		Xmlns:              "http://s3.amazonaws.com/doc/2006-03-01/",
		LocationConstraint: region,
	}
	data, err := xml.Marshal(xmlData)
	if err != nil {
		log.Fatalf("Failed to marshal XML: %v", err)
	}

	// Pretty print XML for debugging
	fmt.Println("Request XML:")
	fmt.Println(xmlPrettyPrint(data))

	// Send the HTTP PUT request to create the bucket
	url := fmt.Sprintf("https://%s.%s", bucketName, endpoint)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
	if err != nil {
		log.Fatalf("Failed to create request: %v", err)
	}

	// Add necessary headers
	req.Header.Set("Content-Type", "application/xml")
	req.SetBasicAuth(accessKeyID, secretAccessKey)

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to execute request: %v", err)
	}
	defer resp.Body.Close()

	// Handle response
	if resp.StatusCode == http.StatusOK {
		fmt.Printf("Created bucket %s OK\n", bucketName)
	} else {
		fmt.Printf("Failed to create bucket %s: %s\n", bucketName, resp.Status)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Must supply a bucket name as argument")
	}

	// Set your AWS Access Key and Secret Key
	accessID := ""      // Replace with your AWS Access Key ID
	accessKey := ""     // Replace with your AWS Secret Access Key
	bucket := os.Args[1]

	createBucket(bucket, accessID, accessKey)
}

