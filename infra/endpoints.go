package main

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2"
	"github.com/aws/aws-cdk-go/awscdk/awsapigatewayv2integrations"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
)

var lambdaEnv *map[string]*string = &map[string]*string{
	"stage": jsii.String(string(stage)),
}

func SetEndpoints(stack constructs.Construct) {

	lambdaGraphQL := awslambda.NewFunction(stack, SetName("lambda-GraphQL"), SetGoFunctionPropsGeneric("../graphql", "make build-http", "http"))

	integration := awsapigatewayv2integrations.NewLambdaProxyIntegration(&awsapigatewayv2integrations.LambdaProxyIntegrationProps{
		Handler:              lambdaGraphQL,
		PayloadFormatVersion: awsapigatewayv2.PayloadFormatVersion_VERSION_2_0(),
	})

	api := awsapigatewayv2.NewHttpApi(stack, SetName("HTTP-GraphQL"), &awsapigatewayv2.HttpApiProps{
		ApiName: SetName("http-graphql"),
		CorsPreflight: &awsapigatewayv2.CorsPreflightOptions{
			AllowOrigins: jsii.Strings("*"),
			AllowMethods: &[]awsapigatewayv2.CorsHttpMethod{
				awsapigatewayv2.CorsHttpMethod_ANY,
			},
			AllowHeaders: jsii.Strings("*"),
			// AllowCredentials: jsii.Bool(true),
			ExposeHeaders: jsii.Strings("*"),
		},
	})

	api.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path:        jsii.String("/graphql/{operation}"),
		Integration: integration,
		AuthorizationScopes: &[]*string{
			jsii.String("AWS_IAM"),
		},
	})

	api.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path:        jsii.String("/graphqltest"),
		Integration: integration,
	})

	api.AddRoutes(&awsapigatewayv2.AddRoutesOptions{
		Path:        jsii.String("/graphql"),
		Integration: integration,
		AuthorizationScopes: &[]*string{
			jsii.String("AWS_IAM"),
		},
		Methods: &[]awsapigatewayv2.HttpMethod{
			awsapigatewayv2.HttpMethod_POST,
		},
	})

	// path := jsii.String("/graphql")

	// graphQLRoute := awsapigatewayv2.NewHttpRoute(stack, SetName("GraphQL-protected"), &awsapigatewayv2.HttpRouteProps{
	// 	HttpApi:     api,
	// 	RouteKey:    awsapigatewayv2.HttpRouteKey_With(path, awsapigatewayv2.HttpMethod_ANY),
	// 	Integration: integration,
	// })

	// client := client.GetClient()

	// client.CastAndSetToPtr()

	// cfnRoute := graphQLRoute.Node().DefaultChild()

	// // cfnRoute := graphQLRoute.Node().DefaultChild().(awsapigatewayv2.CfnRoute)
	// cfnRoute.(awsapigatewayv2.CfnRoute).SetAuthorizationType(jsii.String("AWS_IAM"))
	// println("cfnRoute", cfnRoute)
	// fmt.Printf("%+v\n", cfnRoute)

	awscdk.NewCfnOutput(stack, jsii.String("GraphQL/Endpoint"), &awscdk.CfnOutputProps{
		Value: api.ApiEndpoint(),
	})

}

type AWSIAMAuthorizer struct {
	awsapigatewayv2.IHttpRouteAuthorizer
}

func (auth AWSIAMAuthorizer) Bind(options *awsapigatewayv2.HttpRouteAuthorizerBindOptions) *awsapigatewayv2.HttpRouteAuthorizerConfig {
	return &awsapigatewayv2.HttpRouteAuthorizerConfig{
		AuthorizationType: jsii.String("AWS_IAM"),
	}
}
