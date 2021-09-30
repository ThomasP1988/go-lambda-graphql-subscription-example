//go:build eventstreamlambda
// +build eventstreamlambda

package main

import (
	"context"
	"os"
	"shared/config"
	common "shared/repositories"

	"github.com/ThomasP1988/go-lambda-graphql-subscription/manager"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var conf *config.Config = config.GetConfig(&config.DEV)

var _ error = common.SetAWSConfig(&conf.Region, &conf.LocalProfile)

func HandleRequest(ctx context.Context, req events.DynamoDBEvent) error {
	println("stage", os.Getenv("stage"))

	schema, err := GetGraphQLSchema()
	if err != nil {
		println(err)
	}

	err = SetLibForGraphQLSubscriptions(schema)
	if err != nil {
		println(err)
	}

	return manager.DynamoDBStream(ctx, req)
}

func main() {
	lambda.Start(HandleRequest)
}
