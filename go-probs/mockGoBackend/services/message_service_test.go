package services

import (
	"context"
	"testing"

	"mockGoBackend/internal/apperrors"
	"mockGoBackend/models"
)

type mockMessageRepo struct {
	byUser map[uint][]models.Message
}

func newMockMessageRepo() *mockMessageRepo {
	return &mockMessageRepo{byUser: make(map[uint][]models.Message)}
}

func (m *mockMessageRepo) Create(_ context.Context, msg *models.Message) error {
	m.byUser[msg.UserID] = append(m.byUser[msg.UserID], *msg)
	return nil
}

func (m *mockMessageRepo) ListByUserID(_ context.Context, userID uint) ([]models.Message, error) {
	return m.byUser[userID], nil
}

func TestMessageService_GetMessagesForUsers(t *testing.T) {
	userRepo := newMockUserRepo()
	userRepo.users[1] = models.User{ID: 1, Email: "a@b.com", Name: "A"}
	userRepo.users[2] = models.User{ID: 2, Email: "b@b.com", Name: "B"}

	msgRepo := newMockMessageRepo()
	msgRepo.byUser[1] = []models.Message{{UserID: 1, Body: "hello"}}
	msgRepo.byUser[2] = []models.Message{{UserID: 2, Body: "world"}}

	svc := NewMessageService(msgRepo, userRepo)

	tests := []struct {
		name      string
		userIDs   []uint
		wantCount int
		wantErr   error
	}{
		{name: "fetch two users", userIDs: []uint{1, 2}, wantCount: 2},
		{name: "empty ids", userIDs: nil, wantErr: apperrors.ErrInvalidInput},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := svc.GetMessagesForUsers(context.Background(), tt.userIDs)

			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("expected error %v", tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(got) != tt.wantCount {
				t.Fatalf("got %d results, want %d", len(got), tt.wantCount)
			}
		})
	}
}
