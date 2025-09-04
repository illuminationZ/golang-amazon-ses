# Amazon SES Transaction API

A Go-based REST API using Fiber framework that sends email notifications via Amazon SES when transactions are processed.

## Features

- RESTful API with Fiber framework
- Transaction processing endpoint
- Email notifications via Amazon SES
- CORS and logging middleware
- Environment variable configuration

## Prerequisites

- Go 1.24.2 or later
- AWS account with SES configured
- Verified sender and recipient email addresses in SES

## Environment Variables

Create a `.env` file with the following variables:

```env
AWS_REGION=your-aws-region
AWS_ACCESS_KEY_ID=your-access-key
AWS_SECRET_ACCESS_KEY=your-secret-key
SENDER_EMAIL=your-verified-sender@example.com
RECIPIENT_EMAIL=recipient@example.com
PORT=8080
```

## Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Set up your environment variables
4. Run the application:
   ```bash
   go run main.go
   ```

## API Endpoints

### GET /
Health check endpoint that returns server status.

**Response:**
```json
{
  "message": "Amazon SES API Server is running",
  "version": "1.0.0"
}
```

### POST /transactions
Process a transaction and send email notification.

**Request Body:**
```json
{
  "amount": 99.99,
  "description": "Payment for services"
}
```

**Response (Success):**
```json
{
  "status": "success",
  "message": "Transaction processed successfully and email notification sent"
}
```

**Response (Error):**
```json
{
  "status": "error",
  "message": "Error description"
}
```

## Example Usage

```bash
# Test server status
curl http://localhost:8080/

# Process a transaction
curl -X POST http://localhost:8080/transactions \
  -H "Content-Type: application/json" \
  -d '{"amount": 99.99, "description": "Test transaction"}'
```

## Email Template

When a transaction is processed successfully, an email is sent with:
- Transaction amount
- Transaction description
- Success confirmation message

The email is sent in both HTML and plain text formats.

## Error Handling

The API includes comprehensive error handling for:
- Invalid request bodies
- Missing environment variables
- SES email sending failures
- Server configuration issues

## Project Structure

```
├── main.go                    # Application entry point and server setup
├── config/
│   └── config.go             # Configuration management and environment variables
├── models/
│   └── models.go             # Data models and structs
├── handlers/
│   ├── health_handler.go     # Health check endpoint handler
│   └── transaction_handler.go # Transaction processing handler
├── services/
│   └── email_service.go      # Email service with SES integration
├── sesclient/
│   └── sesclient.go          # SES client configuration
├── go.mod                    # Go module dependencies
├── go.sum                    # Dependency checksums
├── .env                      # Environment variables (not in repo)
├── .gitignore               # Git ignore rules
└── README.md                # This file
```

## Architecture

The application follows a clean architecture pattern with clear separation of concerns:

- **main.go**: Application bootstrap, dependency injection, and server configuration
- **config/**: Environment variable loading and application configuration
- **models/**: Data transfer objects and API models
- **handlers/**: HTTP request handlers for different endpoints
- **services/**: Business logic and external service integrations
- **sesclient/**: AWS SES client configuration and setup
