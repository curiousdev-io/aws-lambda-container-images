package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	lambda_handler "github.com/curiousdev-io/aws-lambda-container-images/aws-os-only-images/go/internal/lambda"
)

func main() {
	// Example: Initialize AWS SDK v2 config (for any AWS service)
	ctx := context.Background()
	_, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		panic("unable to load AWS SDK config: " + err.Error())
	}

	if os.Getenv("AWS_LAMBDA_RUNTIME_API") != "" {
		lambda.Start(lambda_handler.LambdaHandler)
	} else {
		// TODO: Add HTTP handler for local/Fargate testing if needed
		println("No HTTP handler implemented. Lambda mode only.")
	}
}
