package queue

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

// Task type constants
const (
	TypeEmailNotification = "email:notification"
)

// EmailTask represents the payload for email tasks
type EmailTask struct {
	Sender      string  `json:"sender"`
	Recipient   string  `json:"recipient"`
	Subject     string  `json:"subject"`
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

// NewEmailTask creates a new email task
func NewEmailTask(sender, recipient, subject string, amount float64, description string) (*asynq.Task, error) {
	payload := EmailTask{
		Sender:      sender,
		Recipient:   recipient,
		Subject:     subject,
		Amount:      amount,
		Description: description,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TypeEmailNotification, data), nil
}
