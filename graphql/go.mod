module graphql

go 1.17

require (
	github.com/ThomasP1988/go-lambda-graphql-subscription v0.0.1
	github.com/aws/aws-lambda-go v1.26.0
	github.com/graphql-go/graphql v0.8.0
	shared v0.0.1
)

require (
	github.com/aws/aws-sdk-go-v2 v1.9.0 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.8.0 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.4.0 // indirect
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue v1.2.0 // indirect
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression v1.2.2 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.5.0 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.0.4 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.2.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi v1.4.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.5.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.4.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.3.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.1.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.3.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.4.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.7.0 // indirect
	github.com/aws/smithy-go v1.8.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
)

replace (
	// github.com/ThomasP1988/go-lambda-graphql-subscription v0.0.1 => ./../go-lambda-graphql-subscription
	shared v0.0.1 => ./../shared
)
