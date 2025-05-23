package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

func createBucket(bucket string) error {
	// Load the AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-west-2"))
	if err != nil {
		return fmt.Errorf("unable to load SDK config, %v", err)
	}

	// Create a new S3 client
	s3Client := s3.NewFromConfig(cfg)

	// Create the S3 bucket
	_, err = s3Client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(bucket),
		CreateBucketConfiguration: &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraintEuWest2,
		},
	})
	if err != nil {
		return fmt.Errorf("unable to create bucket, %v", err)
	}

	fmt.Printf("Bucket %s successfully created\n", bucket)
	return nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Bucket name is required as the first argument")
	}

	bucket := os.Args[1]

	err := createBucket(bucket)
	if err != nil {
		log.Fatalf("Failed to create bucket: %v", err)
	}
}

