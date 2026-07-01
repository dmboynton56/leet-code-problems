package notifier

import (
	"context"
	"testing"
)

func TestNotifier_Send(t *testing.T) {
	tests := []struct {
		name     string
		notifier Notifier
		n        Notification
		wantErr  bool
	}{
		{
			name:     "email ok",
			notifier: NewEmailNotifier(),
			n:        Notification{To: "user@example.com", Subject: "hi"},
		},
		{
			name:     "email missing to",
			notifier: NewEmailNotifier(),
			n:        Notification{Subject: "hi"},
			wantErr:  true,
		},
		{
			name:     "slack ok",
			notifier: NewSlackNotifier("#alerts"),
			n:        Notification{To: "@user", Body: "ping"},
		},
		{
			name:     "slack missing channel",
			notifier: NewSlackNotifier(""),
			n:        Notification{To: "@user"},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.notifier.Send(context.Background(), tt.n)
			if tt.wantErr && err == nil {
				t.Fatal("expected error, got nil")
			}
			if !tt.wantErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}
