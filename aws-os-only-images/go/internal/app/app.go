package app

import (
	"encoding/json"
	"log/slog"
	"os"
	"strings"
	"time"
)

type Response struct {
	Timestamp string `json:"timestamp"`
	Status    int    `json:"status"`
	Message   string `json:"message"`
}

func HandleRequest(path string, query map[string]string) (int, []byte) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	name := "World"
	if v, ok := query["name"]; ok {
		name = v
	}
	now := time.Now().UTC().Format(time.RFC3339)
	status := 200
	var message string
	if strings.HasPrefix(path, "/hello") {
		message = "Hello, " + name
	} else if strings.HasPrefix(path, "/goodbye") {
		message = "Goodbye, " + name
	} else {
		status = 404
		message = "Not found"
	}
	resp := Response{
		Timestamp: now,
		Status:    status,
		Message:   message,
	}
	logger.Info("request processed", "path", path, "query", query, "status", status, "message", message)
	body, _ := json.Marshal(resp)
	return status, body
}
