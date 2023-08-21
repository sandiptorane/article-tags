package connection

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

const (
	LocalEndpoint = "AWS_ENDPOINT"
	LocalRegion   = "DEFAULT_REGION"
)

// GetConnection : returns db connection
func GetConnection() (*dynamodb.Client, error) {
	awsCfg, err := NewAWS()
	if err != nil {
		log.Println("new Aws error", err)
		return nil, err
	}

	return dynamodb.NewFromConfig(awsCfg), nil
}

func NewAWS() (aws.Config, error) {
	customResolver := aws.EndpointResolverWithOptionsFunc(
		func(service string, region string, options ...any) (aws.Endpoint, error) {
			return aws.Endpoint{
				PartitionID:   "aws",
				URL:           os.Getenv(LocalEndpoint),
				SigningRegion: os.Getenv(LocalRegion),
			}, nil
		},
	)

	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile("AWS_PROFILE"),
		config.WithEndpointResolverWithOptions(customResolver))

	if err != nil {
		return aws.Config{}, err
	}

	return awsCfg, nil
}
