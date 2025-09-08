package models

import (
	"testing"
)

func TestTransactionRequest_Validation(t *testing.T) {
	tests := []struct {
		name    string
		request TransactionRequest
		wantErr bool
	}{
		{
			name: "valid transaction request",
			request: TransactionRequest{
				Email:       "test@example.com",
				Amount:      99.99,
				Description: "Test transaction",
			},
			wantErr: false,
		},
		{
			name: "empty email",
			request: TransactionRequest{
				Email:       "",
				Amount:      99.99,
				Description: "Test transaction",
			},
			wantErr: true,
		},
		{
			name: "invalid email format",
			request: TransactionRequest{
				Email:       "invalid-email",
				Amount:      99.99,
				Description: "Test transaction",
			},
			wantErr: true,
		},
		{
			name: "zero amount",
			request: TransactionRequest{
				Email:       "test@example.com",
				Amount:      0,
				Description: "Test transaction",
			},
			wantErr: true,
		},
		{
			name: "negative amount",
			request: TransactionRequest{
				Email:       "test@example.com",
				Amount:      -10.0,
				Description: "Test transaction",
			},
			wantErr: true,
		},
		{
			name: "empty description",
			request: TransactionRequest{
				Email:       "test@example.com",
				Amount:      99.99,
				Description: "",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.request.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("TransactionRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestHealthResponse_Structure(t *testing.T) {
	response := HealthResponse{
		Message: "Test message",
		Version: "1.0.0",
	}

	if response.Message != "Test message" {
		t.Errorf("Expected message 'Test message', got '%s'", response.Message)
	}

	if response.Version != "1.0.0" {
		t.Errorf("Expected version '1.0.0', got '%s'", response.Version)
	}
}