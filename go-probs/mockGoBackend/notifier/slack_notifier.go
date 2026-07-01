package notifier

import (
	"context"
	"fmt"
	"log"
)

// SlackNotifier is another Notifier — same interface, different transport.
// Services accept notifier.Notifier and can swap implementations at wiring time.
type SlackNotifier struct {
	Channel string
}

func NewSlackNotifier(channel string) *SlackNotifier {
	return &SlackNotifier{Channel: channel}
}

func (s *SlackNotifier) Send(ctx context.Context, n Notification) error {
	if s.Channel == "" {
		return fmt.Errorf("slack: missing channel")
	}
	log.Printf("[slack] channel=%s to=%s body=%q", s.Channel, n.To, n.Body)
	return nil
}
