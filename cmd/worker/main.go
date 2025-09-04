// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"golang-aws-ses/config"
	"golang-aws-ses/queue"
	"golang-aws-ses/services"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize email service
	emailService := services.NewEmailService()

	// Initialize worker service
	workerService := queue.NewWorkerService(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB, emailService)

	// Handle graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		log.Println("Received shutdown signal")
		workerService.Stop()
	}()

	// Start the worker
	log.Println("Starting worker...")
	if err := workerService.Start(); err != nil {
		log.Fatalf("Failed to start worker: %v", err)
	}
}
