package main

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
	"github.com/aws/jsii-runtime-go"
)

func SetName(name string) *string {
	return jsii.String(string(stage) + "-" + name)
}

func SetGoFunctionPropsGeneric(path string, buildCommand string, handler string) *awslambda.FunctionProps {
	var environment map[string]*string = map[string]*string{
		"CGO_ENABLED": jsii.String("0"),
		"GOOS":        jsii.String("linux"),
		"GOARCH":      jsii.String("amd64"),
	}

	return &awslambda.FunctionProps{
		Code: awslambda.NewAssetCode(jsii.String(path), &awss3assets.AssetOptions{
			Bundling: &awscdk.BundlingOptions{
				Image:       awslambda.Runtime_GO_1_X().BundlingDockerImage(),
				User:        jsii.String("root"),
				Environment: &environment,
				Command: &[]*string{
					jsii.String("bash"),
					jsii.String("-c"),
					jsii.String(buildCommand),
				},
			},
		}),
		Handler:       jsii.String(handler),
		Timeout:       awscdk.Duration_Seconds(jsii.Number(300)),
		Runtime:       awslambda.Runtime_GO_1_X(),
		Environment:   lambdaEnv,
		InitialPolicy: &[]awsiam.PolicyStatement{rights},
	}
}
