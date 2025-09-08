// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang-aws-ses/config"
	"golang-aws-ses/handlers"
	"golang-aws-ses/queue"
	"golang-aws-ses/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize services
	emailService := services.NewEmailService()
	queueService := queue.NewQueueService(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)
	defer queueService.Close()

	// Start worker in a separate goroutine (for development/single instance deployment)
	workerService := queue.NewWorkerService(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB, emailService)
	go func() {
		log.Println("Starting embedded worker...")
		if err := workerService.Start(); err != nil {
			log.Printf("Worker error: %v", err)
		}
	}()

	// Initialize handlers
	healthHandler := handlers.NewHealthHandler()
	transactionHandler := handlers.NewTransactionHandler(queueService, cfg)

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return ctx.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Routes
	app.Get("/", healthHandler.HealthCheck)
	app.Post("/transactions", transactionHandler.ProcessTransaction)

	// Handle graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Received shutdown signal, shutting down gracefully...")

		// Shutdown HTTP server
		log.Println("Shutting down HTTP server...")
		if err := app.Shutdown(); err != nil {
			log.Printf("Error shutting down server: %v", err)
		}

		// Shutdown worker
		log.Println("Shutting down worker...")
		workerService.Stop()
		log.Println("Worker shut down gracefully")
	}()

	// Start server
	log.Printf("Server starting on port %s", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Printf("Server stopped: %v", err)
	}
}
