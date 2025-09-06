package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"golang-aws-ses/services"

	"github.com/hibiken/asynq"
)

// WorkerService handles task processing
type WorkerService struct {
	server       *asynq.Server
	mux          *asynq.ServeMux
	emailService *services.EmailService
}

// NewWorkerService creates a new WorkerService instance
func NewWorkerService(redisAddr, redisPassword string, redisDB int, emailService *services.EmailService) *WorkerService {
	redisOpt := asynq.RedisClientOpt{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	}

	server := asynq.NewServer(redisOpt, asynq.Config{
		Concurrency: 10,
		Queues: map[string]int{
			"critical": 1000,
			"default":  100,
			"low":      50,
		},
	})

	mux := asynq.NewServeMux()

	worker := &WorkerService{
		server:       server,
		mux:          mux,
		emailService: emailService,
	}

	// Register task handlers
	worker.registerHandlers()

	return worker
}

// registerHandlers registers all task handlers
func (w *WorkerService) registerHandlers() {
	w.mux.HandleFunc(TypeEmailNotification, w.handleEmailNotification)
}

// handleEmailNotification processes email notification tasks
func (w *WorkerService) handleEmailNotification(ctx context.Context, t *asynq.Task) error {
	var payload EmailTask
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	log.Printf("Processing email task: %+v", payload)

	// Generate email content
	htmlBody := fmt.Sprintf(`
		<h1>Transaction Successful</h1>
		<p>Your transaction has been processed successfully.</p>
		<p><strong>Details:</strong></p>
		<ul>
			<li>Amount: $%.2f</li>
			<li>Description: %s</li>
		</ul>
		<p>Thank you for using our service!</p>
	`, payload.Amount, payload.Description)

	textBody := fmt.Sprintf("Transaction Successful! Amount: $%.2f, Description: %s",
		payload.Amount, payload.Description)

	// Send email
	if err := w.emailService.SendEmail(ctx, payload.Sender, payload.Recipient, payload.Subject, htmlBody, textBody); err != nil {
		log.Printf("Failed to send email: %v", err)
		return err
	}

	log.Printf("Email sent successfully to %s", payload.Recipient)
	return nil
}

// Start starts the worker server
func (w *WorkerService) Start() error {
	log.Println("Starting worker server...")
	return w.server.Run(w.mux)
}

// Stop stops the worker server
func (w *WorkerService) Stop() {
	log.Println("Stopping worker server...")
	w.server.Shutdown()
}
