//go:build wslambda
// +build wslambda

package main

import (
	"context"
	"shared/config"

	"github.com/ThomasP1988/go-lambda-graphql-subscription/manager"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, req events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	return manager.HandleWS(ctx, req)
}

func main() {
	config.GetConfig(nil)

	schema, err := GetGraphQLSchema()
	if err != nil {
		println(err)
	}

	err = SetLibForGraphQLSubscriptions(schema)
	if err != nil {
		println(err)
	}
	lambda.Start(HandleRequest)
}
