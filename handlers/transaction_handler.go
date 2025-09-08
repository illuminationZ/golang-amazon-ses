package handlers

import (
	"context"
	"log"

	"golang-aws-ses/config"
	"golang-aws-ses/models"
	"golang-aws-ses/queue"

	"github.com/gofiber/fiber/v2"
)

// TransactionHandler handles transaction-related requests
type TransactionHandler struct {
	queueService *queue.QueueService
	config       *config.Config
}

// NewTransactionHandler creates a new TransactionHandler instance
func NewTransactionHandler(queueService *queue.QueueService, cfg *config.Config) *TransactionHandler {
	return &TransactionHandler{
		queueService: queueService,
		config:       cfg,
	}
}

// ProcessTransaction handles POST /transactions
func (h *TransactionHandler) ProcessTransaction(c *fiber.Ctx) error {
	var req models.TransactionRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.TransactionResponse{
			Status:  "error",
			Message: "Invalid request body",
		})
	}

	// Validate required fields
	if err := req.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.TransactionResponse{
			Status:  "error",
			Message: err.Error(),
		})
	}

	// Enqueue email notification task
	ctx := context.Background()
	subject := "Transaction Notification"

	if err := h.queueService.EnqueueEmailTask(ctx, h.config.SenderEmail, req.Email, subject, req.Amount, req.Description); err != nil {
		log.Printf("Error enqueuing email task: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.TransactionResponse{
			Status:  "error",
			Message: "Transaction processed but email notification failed to queue",
		})
	}

	return c.JSON(models.TransactionResponse{
		Status:  "success",
		Message: "Transaction processed successfully and email notification queued",
	})
}
