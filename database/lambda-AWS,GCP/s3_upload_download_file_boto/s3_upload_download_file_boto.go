package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/credentials"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

var (
	accessID  = "" // Add your access key
	accessKey = "" // Add your secret key
	region    = "eu-west-2"
)

// Upload a file to the S3 bucket
func uploadFile(bucketName, filename string) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create S3 client
	svc := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Credentials = credentials.NewStaticCredentialsProvider(accessID, accessKey, "")
	})

	// Open the file for uploading
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Unable to open file %q, %v", filename, err)
	}
	defer file.Close()

	// Upload the file to the bucket
	_, err = svc.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		log.Fatalf("Unable to upload %q to %q, %v", filename, bucketName, err)
	}

	fmt.Printf("Successfully uploaded %q to %q\n", filename, bucketName)
}

// Download a file from the S3 bucket
func downloadFile(bucketName, s3Name string) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create S3 client
	svc := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.Credentials = credentials.NewStaticCredentialsProvider(accessID, accessKey, "")
	})

	// Create the file for downloading
	file, err := os.Create("Python_download.png")
	if err != nil {
		log.Fatalf("Unable to create file %v", err)
	}
	defer file.Close()

	// Download the file from the bucket
	resp, err := svc.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(s3Name),
	})

	if err != nil {
		var noSuchKey *types.NoSuchKey
		if err != nil && resp.StatusCode == 404 {
			fmt.Println("The object does not exist.")
		} else if err != nil && resp.StatusCode != 404 {
			log.Fatalf("Unable to download item %q from %q, %v", s3Name, bucketName, err)
		}
		return
	}

	// Write the content of the object to the file
	_, err = file.ReadFrom(resp.Body)
	if err != nil {
		log.Fatalf("Failed to write the file %v", err)
	}
	fmt.Println("Downloaded image Python_download.png")
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <bucket_name> <filename>")
		return
	}

	bucket := os.Args[1]
	filename := os.Args[2]

	uploadFile(bucket, filename)
	downloadFile(bucket, filename)
}

