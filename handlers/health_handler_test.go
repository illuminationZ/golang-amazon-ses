package handlers

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"golang-aws-ses/models"

	"github.com/gofiber/fiber/v2"
)

func TestHealthHandler_HealthCheck(t *testing.T) {
	// Create a new fiber app
	app := fiber.New()
	
	// Create handler
	handler := NewHealthHandler()
	app.Get("/", handler.HealthCheck)

	// Create test request
	req := httptest.NewRequest("GET", "/", nil)
	
	// Perform the request
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to perform request: %v", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status %d, got %d", fiber.StatusOK, resp.StatusCode)
	}

	// Check response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var response models.HealthResponse
	if err := json.Unmarshal(body, &response); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	expectedMessage := "Amazon SES API Server is running"
	if response.Message != expectedMessage {
		t.Errorf("Expected message '%s', got '%s'", expectedMessage, response.Message)
	}

	expectedVersion := "1.0.0"
	if response.Version != expectedVersion {
		t.Errorf("Expected version '%s', got '%s'", expectedVersion, response.Version)
	}
}

func TestNewHealthHandler(t *testing.T) {
	handler := NewHealthHandler()
	if handler == nil {
		t.Error("NewHealthHandler() returned nil")
	}
}