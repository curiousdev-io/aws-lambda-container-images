package lambda

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestLambdaHandler_Hello(t *testing.T) {
	ctx := context.Background()
	event := events.APIGatewayV2HTTPRequest{
		RawPath: "/hello",
		QueryStringParameters: map[string]string{
			"name": "Chris",
		},
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: "GET",
				Path:   "/hello",
			},
		},
	}

	response, err := LambdaHandler(ctx, event)

	if err != nil {
		t.Fatalf("Handler returned error: %v", err)
	}

	if response.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", response.StatusCode)
	}

	var body map[string]interface{}
	if err := json.Unmarshal([]byte(response.Body), &body); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if body["message"] != "Hello, Chris" {
		t.Errorf("Expected message 'Hello, Chris', got '%v'", body["message"])
	}
}

func TestLambdaHandler_Goodbye(t *testing.T) {
	ctx := context.Background()
	event := events.APIGatewayV2HTTPRequest{
		RawPath: "/goodbye",
		QueryStringParameters: map[string]string{
			"name": "Chris",
		},
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: "GET",
				Path:   "/goodbye",
			},
		},
	}

	response, err := LambdaHandler(ctx, event)

	if err != nil {
		t.Fatalf("Handler returned error: %v", err)
	}

	if response.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", response.StatusCode)
	}

	var body map[string]interface{}
	if err := json.Unmarshal([]byte(response.Body), &body); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if body["message"] != "Goodbye, Chris" {
		t.Errorf("Expected message 'Goodbye, Chris', got '%v'", body["message"])
	}
}

func TestLambdaHandler_HelloNoName(t *testing.T) {
	ctx := context.Background()
	event := events.APIGatewayV2HTTPRequest{
		RawPath: "/hello",
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: "GET",
				Path:   "/hello",
			},
		},
	}

	response, err := LambdaHandler(ctx, event)

	if err != nil {
		t.Fatalf("Handler returned error: %v", err)
	}

	if response.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", response.StatusCode)
	}

	var body map[string]interface{}
	if err := json.Unmarshal([]byte(response.Body), &body); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if body["message"] != "Hello, World" {
		t.Errorf("Expected message 'Hello, World', got '%v'", body["message"])
	}
}

func TestLambdaHandler_GoodbyeNoName(t *testing.T) {
	ctx := context.Background()
	event := events.APIGatewayV2HTTPRequest{
		RawPath: "/goodbye",
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: "GET",
				Path:   "/goodbye",
			},
		},
	}

	response, err := LambdaHandler(ctx, event)

	if err != nil {
		t.Fatalf("Handler returned error: %v", err)
	}

	if response.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", response.StatusCode)
	}

	var body map[string]interface{}
	if err := json.Unmarshal([]byte(response.Body), &body); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if body["message"] != "Goodbye, World" {
		t.Errorf("Expected message 'Goodbye, World', got '%v'", body["message"])
	}
}

func TestLambdaHandler_UnknownPath(t *testing.T) {
	ctx := context.Background()
	event := events.APIGatewayV2HTTPRequest{
		RawPath: "/unknown",
		RequestContext: events.APIGatewayV2HTTPRequestContext{
			HTTP: events.APIGatewayV2HTTPRequestContextHTTPDescription{
				Method: "GET",
				Path:   "/unknown",
			},
		},
	}

	response, err := LambdaHandler(ctx, event)

	if err != nil {
		t.Fatalf("Handler returned error: %v", err)
	}

	if response.StatusCode != 404 {
		t.Errorf("Expected status code 404, got %d", response.StatusCode)
	}

	var body map[string]interface{}
	if err := json.Unmarshal([]byte(response.Body), &body); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if body["message"] != "Not found" {
		t.Errorf("Expected message 'Not found', got '%v'", body["message"])
	}
}
