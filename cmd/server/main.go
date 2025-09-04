// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0
package main

import (
	"log"

	"golang-aws-ses/config"
	"golang-aws-ses/handlers"
	"golang-aws-ses/queue"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize queue service
	queueService := queue.NewQueueService(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)
	defer queueService.Close()

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

	// Start server
	log.Printf("API server starting on port %s", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
