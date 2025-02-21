package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
)

const RDS_SECURITY_GROUP_NAME = "my-rds-public-sg"
const RDS_DB_SUBNET_NAME = "my-rds-subnet-group"

// DynamoDB client
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
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(iops.ReadCapacityUnits),
			WriteCapacityUnits: aws.Int64(iops.WriteCapacityUnits),
		},
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

// EC2 client
type EC2Client struct {
	client *ec2.Client
}

func NewEC2Client() *EC2Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-west-3"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	client := ec2.NewFromConfig(cfg)
	return &EC2Client{client: client}
}

func (e *EC2Client) CreateSecurityGroup() (string, error) {
	fmt.Println("Creating RDS Security Group with name", RDS_SECURITY_GROUP_NAME)
	output, err := e.client.CreateSecurityGroup(context.TODO(), &ec2.CreateSecurityGroupInput{
		GroupName:   aws.String(RDS_SECURITY_GROUP_NAME),
		Description: aws.String("RDS security group for public access"),
		VpcId:       aws.String("vpc-1732737e"),
	})
	if err != nil {
		return "", err
	}
	return *output.GroupId, nil
}

func (e *EC2Client) AddInboundRuleToSecurityGroup(securityGroupID string) error {
	fmt.Println("Adding inbound access rule to security group", securityGroupID)
	_, err := e.client.AuthorizeSecurityGroupIngress(context.TODO(), &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: aws.String(securityGroupID),
		IpPermissions: []ec2.IpPermission{
			{
				IpProtocol: aws.String("tcp"),
				FromPort:   aws.Int32(5432),
				ToPort:     aws.Int32(5432),
				IpRanges: []ec2.IpRange{
					{
						CidrIp: aws.String("0.0.0.0/0"),
					},
				},
			},
		},
	})
	return err
}

// RDS client
type RDSClient struct {
	client *rds.Client
}

func NewRDSClient() *RDSClient {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-west-3"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	client := rds.NewFromConfig(cfg)
	return &RDSClient{client: client}
}

func (r *RDSClient) CreatePostgreSQLInstance() error {
	ec2Client := NewEC2Client()

	// Create security group and add inbound rules
	securityGroupID, err := ec2Client.CreateSecurityGroup()
	if err != nil {
		return err
	}
	err = ec2Client.AddInboundRuleToSecurityGroup(securityGroupID)
	if err != nil {
		return err
	}

	fmt.Println("Creating Amazon RDS PostgreSQL DB Instance...")
	_, err = r.client.CreateDBInstance(context.TODO(), &rds.CreateDBInstanceInput{
		DBName:             aws.String("MyPostgreSQLDB"),
		DBInstanceIdentifier: aws.String("mypostgresdb"),
		DBInstanceClass:    aws.String("db.t2.micro"),
		Engine:             aws.String("postgres"),
		MasterUsername:     aws.String("postgres"),
		MasterUserPassword: aws.String("mypostgrepassword"),
		AllocatedStorage:   aws.Int32(20),
		PubliclyAccessible: aws.Bool(true),
		VpcSecurityGroupIds: []string{
			securityGroupID,
		},
	})
	if err != nil {
		return err
	}
	fmt.Println("PostgreSQL instance created successfully")
	return nil
}

func (r *RDSClient) DescribeInstances() (*rds.DescribeDBInstancesOutput, error) {
	fmt.Println("Describing all RDS instances...")
	return r.client.DescribeDBInstances(context.TODO(), &rds.DescribeDBInstancesInput{})
}

func (r *RDSClient) DeleteDBInstance(dbIdentifier string) error {
	fmt.Println("Deleting RDS instance with name", dbIdentifier)
	_, err := r.client.DeleteDBInstance(context.TODO(), &rds.DeleteDBInstanceInput{
		DBInstanceIdentifier: aws.String(dbIdentifier),
		SkipFinalSnapshot:    aws.Bool(true),
	})
	return err
}

func main() {
	// DynamoDB Example
	dynamoDBClient := NewDynamoDBClient()
	attributeDefinitions := []types.AttributeDefinition{
		{AttributeName: aws.String("year"), AttributeType: types.ScalarAttributeTypeN},
		{AttributeName: aws.String("title"), AttributeType: types.ScalarAttributeTypeS},
	}
	keySchema := []types.KeySchemaElement{
		{AttributeName: aws.String("year"), KeyType: types.KeyTypeHash},
		{AttributeName: aws.String("title"), KeyType: types.KeyTypeRange},
	}
	initialIops := types.ProvisionedThroughput{
		ReadCapacityUnits: 5,
		WriteCapacityUnits: 5,
	}

	err := dynamoDBClient.CreateTable("Movies", attributeDefinitions, keySchema, initialIops)
	if err != nil {
		log.Fatalf("Failed to create DynamoDB table: %v", err)
	}

	// RDS Example
	rdsClient := NewRDSClient()
	err = rdsClient.CreatePostgreSQLInstance()
	if err != nil {
		log.Fatalf("Failed to create RDS instance: %v", err)
	}

	// Describe Instances
	rdsInstances, err := rdsClient.DescribeInstances()
	if err != nil {
		log.Fatalf("Failed to describe RDS instances: %v", err)
	}
	fmt.Println("RDS Instances:", rdsInstances)

	// Delete an RDS instance
	err = rdsClient.DeleteDBInstance("mypostgresdb")
	if err != nil {
		log.Fatalf("Failed to delete RDS instance: %v", err)
	}
}

