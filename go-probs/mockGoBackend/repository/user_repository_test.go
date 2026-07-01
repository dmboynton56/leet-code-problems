package repository

import (
	"context"
	"errors"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"mockGoBackend/internal/apperrors"
	"mockGoBackend/models"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	if err := db.AutoMigrate(&models.User{}, &models.Message{}); err != nil {
		t.Fatalf("migrate: %v", err)
	}
	return db
}

func TestUserRepository_CRUD(t *testing.T) {
	db := setupTestDB(t)
	repo := NewUserRepository(db)
	ctx := context.Background()

	user := &models.User{Email: "test@example.com", Name: "Test"}
	if err := repo.Create(ctx, user); err != nil {
		t.Fatalf("create: %v", err)
	}

	got, err := repo.GetByID(ctx, user.ID)
	if err != nil {
		t.Fatalf("get: %v", err)
	}
	if got.Email != user.Email {
		t.Fatalf("email mismatch: %s", got.Email)
	}

	_, err = repo.GetByID(ctx, 999)
	if err == nil {
		t.Fatal("expected not found")
	}
	if !errors.Is(err, apperrors.ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}
