// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: MIT-0
package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"golang-aws-ses/sesclient"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/joho/godotenv"
)

const (
	// Subject is the subject line for the email
	Subject = "Amazon SES Test (AWS SDK for Go)"

	// HTMLBody is the HTML body for the email
	HTMLBody = "<h1>Amazon SES Test Email (AWS SDK for Go)</h1><p>This email was sent with " +
		"<a href='https://aws.amazon.com/ses/'>Amazon SES</a> using the " +
		"<a href='https://aws.amazon.com/sdk-for-go/'>AWS SDK for Go</a>.</p>"

	// TextBody is the email body for recipients with non-HTML email clients
	TextBody = "This email was sent with Amazon SES using the AWS SDK for Go."

	// CharSet is the character encoding for the email
	CharSet = "UTF-8"
)

// SendMsg sends an email message to an Amazon SES recipient
// Inputs:
//
//	svc is the Amazon SES service client
//	sender is the email address in the From field
//	recipient is the email address in the To field
//
// Output:
//
//	If success, nil
//	Otherwise, an error from the call to SendEmail
//
// SendMsg sends an email message to an Amazon SES recipient
func SendMsg(ctx context.Context, svc *ses.Client, sender, recipient, subject string) error {
	_, err := svc.SendEmail(ctx, &ses.SendEmailInput{
		Destination: &types.Destination{
			CcAddresses: []string{},
			ToAddresses: []string{recipient},
		},
		Message: &types.Message{
			Body: &types.Body{
				Html: &types.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(HTMLBody),
				},
				Text: &types.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(TextBody),
				},
			},
			Subject: &types.Content{
				Charset: aws.String(CharSet),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(sender),
	})
	return err
}

func main() {
	_ = godotenv.Load()

	senderEmail := os.Getenv("SENDER_EMAIL")
	if senderEmail == "" {
		fmt.Println("You must set the SENDER_EMAIL environment variable to your verified sender email address.")
		return
	}
	recipientEmail := os.Getenv("RECIPIENT_EMAIL")
	if recipientEmail == "" {
		fmt.Println("You must set the RECIPIENT_EMAIL environment variable to your verified recipient email address.")
		return
	}
	sender := flag.String("f", senderEmail, "The email address for the 'From' field")
	recipient := flag.String("t", recipientEmail, "The email address for the 'To' field")
	subject := flag.String("s", "Amazon SES Test (AWS SDK for Go)", "The text for the 'Subject' field")
	flag.Parse()

	if *sender == "" || *recipient == "" || *subject == "" {
		fmt.Println("You must supply an email address for the sender and recipient, and a subject")
		fmt.Println("-f SENDER -t RECIPIENT -s SUBJECT")
		return
	}

	// initialize context and SES client helper
	ctx := context.Background()
	svc := sesclient.NewSESClient()
	// send the email
	if err := SendMsg(ctx, svc, *sender, *recipient, *subject); err != nil {
		fmt.Println("Got an error sending message:")
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Email sent to address: " + *recipient)
}
