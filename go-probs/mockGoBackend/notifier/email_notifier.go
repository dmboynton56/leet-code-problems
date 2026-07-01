package notifier

import (
	"context"
	"fmt"
	"log"
)

// EmailNotifier simulates sending email. Satisfies Notifier without declaring it.
type EmailNotifier struct{}

func NewEmailNotifier() *EmailNotifier {
	return &EmailNotifier{}
}

func (e *EmailNotifier) Send(ctx context.Context, n Notification) error {
	if n.To == "" {
		return fmt.Errorf("email: missing recipient")
	}
	log.Printf("[email] to=%s subject=%q", n.To, n.Subject)
	return nil
}
