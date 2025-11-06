package lambda

import (
	"context"
	"log/slog"

	"github.com/aws/aws-lambda-go/events"
	"github.com/curiousdev-io/aws-lambda-container-images/aws-os-only-images/go/internal/app"
)

// LambdaHandler handles API Gateway HTTP API events
func LambdaHandler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	slog.Info("Lambda handler invoked", "path", request.RawPath, "method", request.RequestContext.HTTP.Method)

	// Extract path and query parameters
	path := request.RawPath
	if path == "" {
		path = request.RequestContext.HTTP.Path
	}

	// Call business logic - HandleRequest returns (int, []byte)
	status, body := app.HandleRequest(path, request.QueryStringParameters)

	// Convert body bytes to string for API Gateway response
	var respBody string
	if body != nil {
		respBody = string(body)
	}

	// Return API Gateway response
	return events.APIGatewayV2HTTPResponse{
		StatusCode: status,
		Body:       respBody,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}
