package repository

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"mockGoBackend/models"
)

type gormMessageRepo struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) MessageRepository {
	return &gormMessageRepo{db: db}
}

func (r *gormMessageRepo) Create(ctx context.Context, msg *models.Message) error {
	if err := r.db.WithContext(ctx).Create(msg).Error; err != nil {
		return fmt.Errorf("create message: %w", err)
	}
	return nil
}

func (r *gormMessageRepo) ListByUserID(ctx context.Context, userID uint) ([]models.Message, error) {
	var messages []models.Message
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&messages).Error
	if err != nil {
		return nil, fmt.Errorf("list messages for user %d: %w", userID, err)
	}
	return messages, nil
}
