module infra

go 1.17

require (
	github.com/aws/aws-cdk-go/awscdk v1.121.0-devpreview
	github.com/aws/constructs-go/constructs/v3 v3.3.99
	github.com/aws/jsii-runtime-go v1.34.0
	shared v0.0.1
)

require github.com/Masterminds/semver/v3 v3.1.1 // indirect

replace shared v0.0.1 => ./../shared
