package models

// TransactionRequest represents the API request for transactions
type TransactionRequest struct {
	Email       string  `json:"email,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	Description string  `json:"description,omitempty"`
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
