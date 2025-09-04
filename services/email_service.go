package services

import (
	"context"
	"fmt"

	"golang-aws-ses/models"
	"golang-aws-ses/sesclient"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

const (
	// CharSet is the character encoding for the email
	CharSet = "UTF-8"
)

// EmailService handles email operations
type EmailService struct {
	sesClient *ses.Client
}

// NewEmailService creates a new EmailService instance
func NewEmailService() *EmailService {
	return &EmailService{
		sesClient: sesclient.NewSESClient(),
	}
}

// SendEmail sends an email message to an Amazon SES recipient
func (e *EmailService) SendEmail(ctx context.Context, sender, recipient, subject, htmlBody, textBody string) error {
	_, err := e.sesClient.SendEmail(ctx, &ses.SendEmailInput{
		Destination: &types.Destination{
			CcAddresses: []string{},
			ToAddresses: []string{recipient},
		},
		Message: &types.Message{
			Body: &types.Body{
				Html: &types.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(htmlBody),
				},
				Text: &types.Content{
					Charset: aws.String(CharSet),
					Data:    aws.String(textBody),
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

// SendTransactionEmail sends a transaction notification email
func (e *EmailService) SendTransactionEmail(ctx context.Context, sender, recipient string, transaction models.TransactionRequest) error {
	subject := "Transaction Notification"

	htmlBody := fmt.Sprintf(`
		<h1>Transaction Successful</h1>
		<p>Your transaction has been processed successfully.</p>
		<p><strong>Details:</strong></p>
		<ul>
			<li>Amount: $%.2f</li>
			<li>Description: %s</li>
		</ul>
		<p>Thank you for using our service!</p>
	`, transaction.Amount, transaction.Description)

	textBody := fmt.Sprintf("Transaction Successful! Amount: $%.2f, Description: %s",
		transaction.Amount, transaction.Description)

	return e.SendEmail(ctx, sender, recipient, subject, htmlBody, textBody)
}
