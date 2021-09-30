package main

import (
	"shared/config"

	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

func SetConnectionsTable(stack constructs.Construct) {
	println("CONNECTION Table:", config.Conf.Tables[config.GRAPHQL_SUB_CONNECTION].Name)

	awsdynamodb.NewTable(stack, jsii.String(config.Conf.Tables[config.GRAPHQL_SUB_CONNECTION].Name), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		TableName:           jsii.String(config.Conf.Tables[config.GRAPHQL_SUB_CONNECTION].Name),
		BillingMode:         awsdynamodb.BillingMode_PAY_PER_REQUEST,
		TimeToLiveAttribute: jsii.String("ttl"),
	})
}

func SetEventsTable(stack constructs.Construct) {
	println("CONNECTION Table:", config.Conf.Tables[config.GRAPHQL_SUB_EVENT].Name)

	table := awsdynamodb.NewTable(stack, jsii.String(config.Conf.Tables[config.GRAPHQL_SUB_EVENT].Name), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		Stream:              awsdynamodb.StreamViewType_NEW_IMAGE,
		TableName:           jsii.String(config.Conf.Tables[config.GRAPHQL_SUB_EVENT].Name),
		BillingMode:         awsdynamodb.BillingMode_PAY_PER_REQUEST,
		TimeToLiveAttribute: jsii.String("ttl"),
	})

	SetEventTableConsumer(stack, table.TableStreamArn())

}

func SetEventTableConsumer(stack constructs.Construct, sourceArn *string) {

	lambdaConsumer := awslambda.NewFunction(stack, SetName("EventStream"), SetGoFunctionPropsGeneric("../graphql", "make build-event-steam", "eventstreamlambda"))

	awslambda.NewEventSourceMapping(stack, SetName("SourceMappingDDBEventsStream"), &awslambda.EventSourceMappingProps{
		BatchSize:        jsii.Number(1),
		Enabled:          jsii.Bool(true),
		RetryAttempts:    jsii.Number(3),
		StartingPosition: awslambda.StartingPosition_TRIM_HORIZON,
		EventSourceArn:   sourceArn,
		Target:           lambdaConsumer,
	})

}

func SetWebsocketEndpoint(stack constructs.Construct) {
	lambdaWebsocket := awslambda.NewFunction(stack, SetName("WebsocketGraphQL"), SetGoFunctionPropsGeneric("../graphql", "make build-ws", "ws"))

	integration := awsapigatewayv2integrations.NewLambdaWebSocketIntegration(&awsapigatewayv2integrations.LambdaWebSocketIntegrationProps{
		Handler: lambdaWebsocket,
	})

	websocketAPI := awsapigatewayv2.NewWebSocketApi(stack, SetName("Websocket-API-GraphQL"), &awsapigatewayv2.WebSocketApiProps{
		ApiName: SetName("ws-graphql"),
		ConnectRouteOptions: &awsapigatewayv2.WebSocketRouteOptions{
			Integration: integration,
		},
		DisconnectRouteOptions: &awsapigatewayv2.WebSocketRouteOptions{
			Integration: integration,
		},
		DefaultRouteOptions: &awsapigatewayv2.WebSocketRouteOptions{
			Integration: integration,
		},
	})

	awsapigatewayv2.NewWebSocketStage(stack, SetName("Websocket-stage"), &awsapigatewayv2.WebSocketStageProps{
		AutoDeploy:   jsii.Bool(true),
		WebSocketApi: websocketAPI,
		StageName:    jsii.String(string(stage)),
	})
}

func SetSubscriptionsTable(stack constructs.Construct) {
	println("CONNECTION Table:", config.Conf.Tables[config.GRAPHQL_SUB_SUBSCRIPTION].Name)

	subscriptionTable := awsdynamodb.NewTable(stack, jsii.String(config.Conf.Tables[config.GRAPHQL_SUB_SUBSCRIPTION].Name), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("event"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		SortKey: &awsdynamodb.Attribute{
			Name: jsii.String("connectionId"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		TableName:           jsii.String(config.Conf.Tables[config.GRAPHQL_SUB_SUBSCRIPTION].Name),
		BillingMode:         awsdynamodb.BillingMode_PAY_PER_REQUEST,
		TimeToLiveAttribute: jsii.String("ttl"),
	})

	subscriptionTable.AddGlobalSecondaryIndex(&awsdynamodb.GlobalSecondaryIndexProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("connectionId"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		SortKey: &awsdynamodb.Attribute{
			Name: jsii.String("operationId"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		ProjectionType: awsdynamodb.ProjectionType_ALL,
		IndexName:      jsii.String(config.Conf.Tables[config.GRAPHQL_SUB_SUBSCRIPTION].SecondaryIndex[config.OperationIndex]),
	})

}
