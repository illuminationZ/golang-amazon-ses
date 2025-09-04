// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package sesclient

import (
	"context"
	"log"
	"os"

	// load .env for local development
	"github.com/joho/godotenv"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/ses"
)

// NewSESClient creates and returns an Amazon SES client configured using
// environment variables (AWS_REGION, AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY) or
// default shared config (~/.aws/config).
func NewSESClient() *ses.Client {
	// load variables from .env (optional)
	_ = godotenv.Load()
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = "ap-southeast-1"
	}

	ctx := context.Background()
	// Prepare default config options with region
	opts := []func(*config.LoadOptions) error{
		config.WithRegion(region),
	}
	// If environment credentials are set, use them explicitly
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	token := os.Getenv("AWS_SESSION_TOKEN")
	if accessKey != "" && secretKey != "" {
		opts = append(opts, config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKey, secretKey, token),
		))
	}
	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		log.Fatalf("unable to load AWS SDK config, %v", err)
	}

	return ses.NewFromConfig(cfg)
}
