# Amazon SES Transaction API with Queue

A Go-based REST API using Fiber framework that processes transactions and sends email notifications via Amazon SES using an asynchronous queue system powered by Asynq and Redis.

## Features

- RESTful API with Fiber framework
- Transaction processing endpoint
- **Asynchronous email notifications** via Asynq queue
- Redis-backed task queue for reliability and scalability
- Separate server and worker binaries for production deployment
- CORS and logging middleware
- Environment variable configuration
- Graceful shutdown handling

## Prerequisites

- Go 1.24.2 or later
- Redis server (local installation or Docker)
- AWS account with SES configured
- Verified sender and recipient email addresses in SES

## Environment Variables

Create a `.env` file with the following variables:

```env
AWS_REGION=your-aws-region
AWS_ACCESS_KEY_ID=your-access-key
AWS_SECRET_ACCESS_KEY=your-secret-key
SENDER_EMAIL=your-verified-sender@example.com
PORT=8080

# Redis Configuration for Queue
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0
```

## Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Start Redis server:
   ```bash
   # Using Docker
   docker run -d -p 6379:6379 redis:alpine
   
   # Or using local installation
   redis-server
   ```
4. Set up your environment variables
5. Run the application:
   
   **Option A: Single process (development)**
   ```bash
   go run main.go
   ```
   
   **Option B: Separate processes (production)**
   ```bash
   # Terminal 1: Start API server
   go run cmd/server/main.go
   
   # Terminal 2: Start worker
   go run cmd/worker/main.go
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
Process a transaction and **queue** an email notification for asynchronous processing.

**Request Body:**
```json
{
  "email": "recipient@example.com",
  "amount": 99.99,
  "description": "Payment for services"
}
```

**Response (Success):**
```json
{
  "status": "success",
  "message": "Transaction processed successfully and email notification queued"
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
  -d '{"email": "user@example.com", "amount": 99.99, "description": "Test transaction"}'
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
- Redis connection failures
- SES email sending failures
- Task queue processing errors
- Server configuration issues

## Deployment

### Development
Use the single process mode with embedded worker:
```bash
go run main.go
```

### Production
Deploy API server and worker as separate processes:

```bash
# Build binaries
go build -o bin/server cmd/server/main.go
go build -o bin/worker cmd/worker/main.go

# Run API server
./bin/server

# Run worker (can run multiple instances)
./bin/worker
```

### Docker Deployment
```dockerfile
# Example Dockerfile for API server
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o server cmd/server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/server .
CMD ["./server"]
```

## Monitoring

### Queue Monitoring
Asynq provides built-in monitoring:
- Queue size and processing rates
- Task failure rates and retry counts
- Worker performance metrics

### Health Checks
- **API Health**: `GET /` endpoint
- **Redis Connection**: Automatic connection health checking
- **Worker Status**: Graceful shutdown signals

## Project Structure

```
├── main.go                    # Application entry point with embedded worker
├── cmd/
│   ├── server/
│   │   └── main.go           # Standalone API server binary
│   └── worker/
│       └── main.go           # Standalone worker binary
├── config/
│   └── config.go             # Configuration management and environment variables
├── models/
│   └── models.go             # Data models and structs
├── handlers/
│   ├── health_handler.go     # Health check endpoint handler
│   └── transaction_handler.go # Transaction processing handler
├── services/
│   └── email_service.go      # Email service with SES integration
├── queue/
│   ├── tasks.go              # Task definitions and creators
│   ├── queue_service.go      # Queue client for enqueuing tasks
│   └── worker_service.go     # Worker server for processing tasks
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

- **main.go**: Application bootstrap with embedded worker (development mode)
- **cmd/server/**: Standalone API server binary for production
- **cmd/worker/**: Standalone worker binary for production
- **config/**: Environment variable loading and application configuration
- **models/**: Data transfer objects and API models
- **handlers/**: HTTP request handlers for different endpoints
- **services/**: Business logic and external service integrations
- **queue/**: Asynq task definitions, queue service, and worker handlers
- **sesclient/**: AWS SES client configuration and setup

## Queue System

The application uses **Asynq** (Redis-based queue) for asynchronous email processing:

- **Queue Service**: Enqueues email tasks when transactions are processed
- **Worker Service**: Processes queued email tasks asynchronously
- **Redis**: Stores and manages task queue
- **Graceful Shutdown**: Ensures all tasks complete before stopping

### Benefits of Queue System:
- **Fast API Response**: Transactions return immediately without waiting for email
- **Reliability**: Failed email tasks can be retried automatically
- **Scalability**: Multiple workers can process emails concurrently
- **Monitoring**: Asynq provides built-in monitoring and metrics
