package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Function to list all S3 buckets
func listBuckets() {
	// Initialize a session in the desired AWS region (e.g., "us-west-2")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2"), // Specify the region here
	})
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}

	// Create an S3 service client
	svc := s3.New(sess)

	// Call S3 to list current buckets
	result, err := svc.ListBuckets(nil)
	if err != nil {
		log.Fatalf("Unable to list buckets: %v", err)
	}

	// Get and print the bucket names
	fmt.Println("Bucket List:")
	for _, bucket := range result.Buckets {
		fmt.Printf("* %s\n", aws.StringValue(bucket.Name))
	}
}

func main() {
	listBuckets()
}

