package repository

import (
	"context"

	"mockGoBackend/models"
)

// UserRepository is an interface so services can depend on behavior, not GORM.
// Go interfaces are satisfied implicitly — any type with matching methods works,
// no "implements" keyword required.
type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id uint) (*models.User, error)
	List(ctx context.Context) ([]models.User, error)
}

// MessageRepository handles message persistence.
type MessageRepository interface {
	Create(ctx context.Context, msg *models.Message) error
	ListByUserID(ctx context.Context, userID uint) ([]models.Message, error)
}
