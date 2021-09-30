package main

import (
	"fmt"
	"graphql/resolvers/user"
	"shared/config"
	"time"

	"github.com/ThomasP1988/go-lambda-graphql-subscription/dynamodb"
	"github.com/ThomasP1988/go-lambda-graphql-subscription/manager"
	"github.com/aws/aws-lambda-go/events"

	"github.com/graphql-go/graphql"
)

func GetQuery() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"getUser": user.GetGQL,
		},
	})
}

func GetSubscription() *graphql.Object {

	type MessageChat struct {
		ID   string `graphql:"id"`
		Text string `graphql:"text"`
		Type string `graphql:"type"`
	}

	var MessageChatType = graphql.NewObject(graphql.ObjectConfig{
		Name: "MessageChatType",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"text": &graphql.Field{
				Type: graphql.String,
			},
			"type": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootSubscription",
		Fields: graphql.Fields{
			"messageFeed": &graphql.Field{
				Type: MessageChatType,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					lambdaContext := p.Context.Value(manager.WSContextKey).(manager.WSlambdaContext)
					fmt.Printf("resolve lambdaContext: %v\n", lambdaContext)

					return p.Source, nil
				},
				Subscribe: func(p graphql.ResolveParams) (interface{}, error) {
					lambdaContext := p.Context.Value(manager.WSContextKey).(manager.WSlambdaContext)
					fmt.Printf("subscribe lambdaContext: %v\n", lambdaContext)

					return manager.Sub([]string{"NEW_MESSAGE"}, p.Context)
				},
			},
		},
	})
}

func GetMutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createUser": user.CreateGQL,
		},
	})
}

func SetLibForGraphQLSubscriptions(schema *graphql.Schema) error {

	connectionManager, err := dynamodb.NewDynamoDBConnectionManager(&dynamodb.DynamoDBConnectionManagerArgs{
		Table: config.Conf.Tables[config.GRAPHQL_SUB_CONNECTION].Name,
		Ttl:   time.Hour * 6,
	})

	if err != nil {
		return err
	}

	subscriptionManager, err := dynamodb.NewDynamoDBSubscriptionManager(&dynamodb.DynamoDBSubscriptionManagerArgs{
		Table:             config.Conf.Tables[config.GRAPHQL_SUB_SUBSCRIPTION].Name,
		Ttl:               time.Hour * 6,
		IndexConnectionID: config.Conf.Tables[config.GRAPHQL_SUB_SUBSCRIPTION].SecondaryIndex[config.OperationIndex],
	})

	if err != nil {
		return err
	}

	eventManager, err := dynamodb.NewDynamoDBEventManager(&dynamodb.DynamoDBEventManagerArgs{
		Table: config.Conf.Tables[config.GRAPHQL_SUB_EVENT].Name,
	})

	if err != nil {
		return err
	}

	OnWebsocketConnect := func(event *events.APIGatewayWebsocketProxyRequest) interface{} {
		fmt.Printf("OnWebsocketConnect event: %v\n", event)

		return map[string]string{
			"OnWebsocketConnect": "OnWebsocketConnect",
		}
	}

	OnConnect := func(event *events.APIGatewayWebsocketProxyRequest) interface{} {
		fmt.Printf("event: %v\n", event)
		return map[string]string{
			"OnConnect": "OnConnect",
		}
	}

	OnDisconnect := func(connection *manager.Connection) {
		fmt.Printf("disconnect connection: %v\n", connection)
	}

	println("SetLibForGraphQLSubscriptions 4")
	manager.SetManager(&manager.SetManagerArgs{
		Schema:             schema,
		Connection:         connectionManager,
		Subscription:       subscriptionManager,
		Event:              eventManager,
		OnWebsocketConnect: &OnWebsocketConnect,
		OnConnect:          &OnConnect,
		OnDisconnect:       &OnDisconnect,
	})
	return nil
}

func GetGraphQLSchema() (*graphql.Schema, error) {
	rootQuery := GetQuery()
	rootMutation := GetMutation()
	rootSubscription := GetSubscription()

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:        rootQuery,
		Mutation:     rootMutation,
		Subscription: rootSubscription,
	})

	if err != nil {
		println(err)
		return nil, err
	}

	return &schema, nil
}
