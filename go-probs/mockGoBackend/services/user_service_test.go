package services

import (
	"context"
	"errors"
	"testing"

	"mockGoBackend/internal/apperrors"
	"mockGoBackend/models"
	"mockGoBackend/notifier"
)

type mockUserRepo struct {
	users map[uint]models.User
	next  uint
}

func newMockUserRepo() *mockUserRepo {
	return &mockUserRepo{users: make(map[uint]models.User), next: 1}
}

func (m *mockUserRepo) Create(_ context.Context, user *models.User) error {
	user.ID = m.next
	m.next++
	m.users[user.ID] = *user
	return nil
}

func (m *mockUserRepo) GetByID(_ context.Context, id uint) (*models.User, error) {
	u, ok := m.users[id]
	if !ok {
		return nil, apperrors.ErrNotFound
	}
	return &u, nil
}

func (m *mockUserRepo) List(_ context.Context) ([]models.User, error) {
	out := make([]models.User, 0, len(m.users))
	for _, u := range m.users {
		out = append(out, u)
	}
	return out, nil
}

func TestUserService_CreateUser(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		userName string
		wantErr error
	}{
		{name: "valid user", email: "a@b.com", userName: "Alice"},
		{name: "missing email", email: "", userName: "Alice", wantErr: apperrors.ErrInvalidInput},
		{name: "missing name", email: "a@b.com", userName: "", wantErr: apperrors.ErrInvalidInput},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := NewUserService(newMockUserRepo(), notifier.NewEmailNotifier())
			_, err := svc.CreateUser(context.Background(), tt.email, tt.userName)

			if tt.wantErr != nil {
				if err == nil {
					t.Fatalf("expected error %v, got nil", tt.wantErr)
				}
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("expected %v, got %v", tt.wantErr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})
	}
}

