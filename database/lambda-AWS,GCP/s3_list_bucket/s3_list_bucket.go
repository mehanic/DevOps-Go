package main

import (
	"bytes"
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/credentials"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

var (
	accessID  = ""  // Add your access key
	accessKey = ""  // Add your secret key
	region    = "eu-west-2"
	endpoint  = fmt.Sprintf("https://s3.%s.amazonaws.com", region)
	ns        = "http://s3.amazonaws.com/doc/2006-03-01/"
)

type ErrorResponse struct {
	XMLName xml.Name `xml:"Error"`
	Code    string   `xml:"Code"`
	Message string   `xml:"Message"`
}

func handleError(body []byte) {
	var errorResponse ErrorResponse
	err := xml.Unmarshal(body, &errorResponse)
	if err != nil {
		log.Fatalf("Failed to parse error response: %v", err)
	}
	fmt.Printf("Error Code: %s\nMessage: %s\n", errorResponse.Code, errorResponse.Message)
}

func xmlPrettyPrint(body []byte) {
	var prettyXML bytes.Buffer
	err := xml.Indent(&prettyXML, body, "", "  ")
	if err != nil {
		log.Fatalf("Failed to pretty print XML: %v", err)
	}
	fmt.Println(prettyXML.String())
}

func listBuckets() {
	// Load AWS Config
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	creds := credentials.NewStaticCredentialsProvider(accessID, accessKey, "")

	// Create the S3 client
	svc := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Credentials = creds
	})

	// Perform the ListBuckets request using AWS4Signer
	signer := v4.NewSigner()
	req, _ := http.NewRequest("GET", endpoint, nil)

	err = signer.SignHTTP(context.TODO(), creds, req, "s3", region, nil)
	if err != nil {
		log.Fatalf("Failed to sign request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response: %v", err)
	}

	if resp.StatusCode == 200 {
		xmlPrettyPrint(body)
	} else {
		handleError(body)
	}
}

func main() {
	listBuckets()
}

