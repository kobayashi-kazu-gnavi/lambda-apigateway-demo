package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"os"
)

func main() {
	lambda.Start(handler)
}

// Handler api gateway http api handler.
func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// get lambda version
	lambdaVersion := os.Getenv("AWS_LAMBDA_FUNCTION_VERSION")
	if lambdaVersion == "" {
		lambdaVersion = "Unknown"
	}

	// create response
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: fmt.Sprintf("{\"lambdaVersion\": \"%s\"}", lambdaVersion),
	}

	return response, nil
}
