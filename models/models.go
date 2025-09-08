package models

import (
	"errors"
	"regexp"
)

// TransactionRequest represents the API request for transactions
type TransactionRequest struct {
	Email       string  `json:"email,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	Description string  `json:"description,omitempty"`
}

// Validate validates the transaction request fields
func (tr *TransactionRequest) Validate() error {
	if tr.Email == "" {
		return errors.New("email is required")
	}

	// Basic email format validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(tr.Email) {
		return errors.New("invalid email format")
	}

	if tr.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	if tr.Description == "" {
		return errors.New("description is required")
	}

	return nil
}

// TransactionResponse represents the API response for transactions
type TransactionResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// HealthResponse represents the health check response
type HealthResponse struct {
	Message string `json:"message"`
	Version string `json:"version"`
}

// ErrorResponse represents error responses
type ErrorResponse struct {
	Error string `json:"error"`
}
