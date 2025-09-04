package queue

import (
	"context"
	"log"

	"github.com/hibiken/asynq"
)

// QueueService handles queue operations
type QueueService struct {
	client *asynq.Client
}

// NewQueueService creates a new QueueService instance
func NewQueueService(redisAddr, redisPassword string, redisDB int) *QueueService {
	redisOpt := asynq.RedisClientOpt{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	}

	client := asynq.NewClient(redisOpt)

	return &QueueService{
		client: client,
	}
}

// EnqueueEmailTask enqueues an email notification task
func (q *QueueService) EnqueueEmailTask(ctx context.Context, sender, recipient, subject string, amount float64, description string) error {
	task, err := NewEmailTask(sender, recipient, subject, amount, description)
	if err != nil {
		return err
	}

	info, err := q.client.Enqueue(task)
	if err != nil {
		return err
	}

	log.Printf("Email task enqueued: ID=%s, Queue=%s", info.ID, info.Queue)
	return nil
}

// Close closes the queue client
func (q *QueueService) Close() error {
	return q.client.Close()
}
