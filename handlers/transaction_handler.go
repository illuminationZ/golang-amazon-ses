package handlers

import (
	"context"
	"log"

	"golang-aws-ses/config"
	"golang-aws-ses/models"
	"golang-aws-ses/services"

	"github.com/gofiber/fiber/v2"
)

// TransactionHandler handles transaction-related requests
type TransactionHandler struct {
	emailService *services.EmailService
	config       *config.Config
}

// NewTransactionHandler creates a new TransactionHandler instance
func NewTransactionHandler(emailService *services.EmailService, cfg *config.Config) *TransactionHandler {
	return &TransactionHandler{
		emailService: emailService,
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

	// Send email notification
	ctx := context.Background()
	if err := h.emailService.SendTransactionEmail(ctx, h.config.SenderEmail, h.config.RecipientEmail, req); err != nil {
		log.Printf("Error sending email: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(models.TransactionResponse{
			Status:  "error",
			Message: "Transaction processed but email notification failed",
		})
	}

	return c.JSON(models.TransactionResponse{
		Status:  "success",
		Message: "Transaction processed successfully and email notification sent",
	})
}
