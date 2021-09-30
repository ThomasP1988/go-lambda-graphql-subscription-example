package main

import (
	"shared/config"

	"github.com/aws/aws-cdk-go/awscdk/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

func SetUserTable(stack constructs.Construct) {
	println("TableName:", config.Conf.Tables[config.USER].Name)

	awsdynamodb.NewTable(stack, jsii.String(config.Conf.Tables[config.USER].Name), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("id"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		TableName:           jsii.String(config.Conf.Tables[config.USER].Name),
		BillingMode:         awsdynamodb.BillingMode_PAY_PER_REQUEST,
		TimeToLiveAttribute: jsii.String("ttl"),
	})
}
