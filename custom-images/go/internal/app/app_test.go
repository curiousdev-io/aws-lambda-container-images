package app

import (
	"encoding/json"
	"testing"
)

func TestHandleRequest_Hello(t *testing.T) {
	// slog logs to stdout; test only business logic
	status, body := HandleRequest("/hello", map[string]string{"name": "Alice"})
	if status != 200 {
		t.Errorf("expected status 200, got %d", status)
	}
	var resp Response
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}
	if resp.Message != "Hello, Alice" {
		t.Errorf("expected message 'Hello, Alice', got '%s'", resp.Message)
	}
}

func TestHandleRequest_Goodbye(t *testing.T) {
	status, body := HandleRequest("/goodbye", map[string]string{"name": "Bob"})
	if status != 200 {
		t.Errorf("expected status 200, got %d", status)
	}
	var resp Response
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}
	if resp.Message != "Goodbye, Bob" {
		t.Errorf("expected message 'Goodbye, Bob', got '%s'", resp.Message)
	}
}

func TestHandleRequest_NotFound(t *testing.T) {
	status, body := HandleRequest("/unknown", nil)
	if status != 404 {
		t.Errorf("expected status 404, got %d", status)
	}
	var resp Response
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}
	if resp.Message != "Not found" {
		t.Errorf("expected message 'Not found', got '%s'", resp.Message)
	}
}
