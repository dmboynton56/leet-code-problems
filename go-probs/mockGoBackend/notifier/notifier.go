package notifier

import "context"

// Notification is the payload passed to any notifier implementation.
type Notification struct {
	To      string
	Subject string
	Body    string
}

// Notifier is a small interface — one method. In Go, interfaces are often defined
// by the consumer (here) not the implementer. EmailNotifier and SlackNotifier
// satisfy Notifier implicitly; no inheritance or registration step.
type Notifier interface {
	Send(ctx context.Context, n Notification) error
}
