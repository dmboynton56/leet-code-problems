package services

import (
	"context"
	"fmt"

	"mockGoBackend/internal/apperrors"
	"mockGoBackend/models"
	"mockGoBackend/notifier"
	"mockGoBackend/repository"
)

type UserService struct {
	users    repository.UserRepository
	notifier notifier.Notifier
}

func NewUserService(users repository.UserRepository, n notifier.Notifier) *UserService {
	return &UserService{users: users, notifier: n}
}

func (s *UserService) CreateUser(ctx context.Context, email, name string) (*models.User, error) {
	if email == "" || name == "" {
		return nil, apperrors.ErrInvalidInput
	}

	user := &models.User{Email: email, Name: name}
	if err := s.users.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("user service create: %w", err)
	}

	// Fire-and-forget welcome notification — errors bubble up so the handler can decide policy.
	if err := s.notifier.Send(ctx, notifier.Notification{
		To:      email,
		Subject: "Welcome",
		Body:    fmt.Sprintf("Hi %s, welcome to the platform.", name),
	}); err != nil {
		return user, fmt.Errorf("user service welcome notify: %w", err)
	}

	return user, nil
}

func (s *UserService) GetUser(ctx context.Context, id uint) (*models.User, error) {
	user, err := s.users.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("user service get: %w", err)
	}
	return user, nil
}

func (s *UserService) ListUsers(ctx context.Context) ([]models.User, error) {
	users, err := s.users.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("user service list: %w", err)
	}
	return users, nil
}
