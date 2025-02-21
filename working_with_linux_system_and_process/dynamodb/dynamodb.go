package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoDBClient struct {
	client *dynamodb.Client
}

func NewDynamoDBClient() *DynamoDBClient {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-west-3"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	client := dynamodb.NewFromConfig(cfg)
	return &DynamoDBClient{client: client}
}

func (d *DynamoDBClient) CreateTable(tableName string, attributeDefinitions []types.AttributeDefinition, keySchema []types.KeySchemaElement, iops types.ProvisionedThroughput) error {
	fmt.Println("Creating DynamoDB table...")
	_, err := d.client.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		TableName:            aws.String(tableName),
		AttributeDefinitions: attributeDefinitions,
		KeySchema:            keySchema,
		ProvisionedThroughput: &iops,
	})
	if err != nil {
		return err
	}
	fmt.Println("Table", tableName, "created successfully")
	return nil
}

func (d *DynamoDBClient) DescribeTable(tableName string) (*dynamodb.DescribeTableOutput, error) {
	fmt.Println("Describing DynamoDB table:", tableName)
	return d.client.DescribeTable(context.TODO(), &dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	})
}

func (d *DynamoDBClient) UpdateTableThroughput(tableName string, readCapacity, writeCapacity int64) error {
	fmt.Println("Updating throughput for table:", tableName)
	_, err := d.client.UpdateTable(context.TODO(), &dynamodb.UpdateTableInput{
		TableName: aws.String(tableName),
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(readCapacity),
			WriteCapacityUnits: aws.Int64(writeCapacity),
		},
	})
	return err
}

func (d *DynamoDBClient) DeleteTable(tableName string) error {
	fmt.Println("Deleting DynamoDB table:", tableName)
	_, err := d.client.DeleteTable(context.TODO(), &dynamodb.DeleteTableInput{
		TableName: aws.String(tableName),
	})
	return err
}

func main() {
	dynamodbClient := NewDynamoDBClient()

	// Create Table
	tableName := "Movies"
	attributeDefinitions := []types.AttributeDefinition{
		{
			AttributeName: aws.String("year"),
			AttributeType: types.ScalarAttributeTypeN,
		},
		{
			AttributeName: aws.String("title"),
			AttributeType: types.ScalarAttributeTypeS,
		},
	}

	keySchema := []types.KeySchemaElement{
		{
			AttributeName: aws.String("year"),
			KeyType:       types.KeyTypeHash, // Partition key
		},
		{
			AttributeName: aws.String("title"),
			KeyType:       types.KeyTypeRange, // Sort key
		},
	}

	initialIops := types.ProvisionedThroughput{
		ReadCapacityUnits:  aws.Int64(5),
		WriteCapacityUnits: aws.Int64(5),
	}

	err := dynamodbClient.CreateTable(tableName, attributeDefinitions, keySchema, initialIops)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	// Describe Table
	describeOutput, err := dynamodbClient.DescribeTable(tableName)
	if err != nil {
		log.Fatalf("Failed to describe table: %v", err)
	}
	fmt.Println("Table description:", describeOutput)

	// Update Table IOPS
	err = dynamodbClient.UpdateTableThroughput(tableName, 10, 10)
	if err != nil {
		log.Fatalf("Failed to update throughput: %v", err)
	} else {
		fmt.Println("Updated table throughput successfully")
	}

	// Delete Table
	err = dynamodbClient.DeleteTable(tableName)
	if err != nil {
		log.Fatalf("Failed to delete table: %v", err)
	} else {
		fmt.Println("Deleted table successfully")
	}
}

