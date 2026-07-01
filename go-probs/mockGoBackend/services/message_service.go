package services

import (
	"context"
	"fmt"
	"sync"

	"mockGoBackend/internal/apperrors"
	"mockGoBackend/models"
	"mockGoBackend/repository"
)

type MessageService struct {
	messages repository.MessageRepository
	users    repository.UserRepository
}

func NewMessageService(messages repository.MessageRepository, users repository.UserRepository) *MessageService {
	return &MessageService{messages: messages, users: users}
}

func (s *MessageService) CreateMessage(ctx context.Context, userID uint, body string) (*models.Message, error) {
	if body == "" {
		return nil, apperrors.ErrInvalidInput
	}

	if _, err := s.users.GetByID(ctx, userID); err != nil {
		return nil, fmt.Errorf("message service create: %w", err)
	}

	msg := &models.Message{UserID: userID, Body: body}
	if err := s.messages.Create(ctx, msg); err != nil {
		return nil, fmt.Errorf("message service create: %w", err)
	}
	return msg, nil
}

// UserMessages bundles messages for one user (used by concurrent fetch).
type UserMessages struct {
	UserID   uint
	Messages []models.Message
}

// fetchResult carries either data or an error back from a goroutine via a channel.
type fetchResult struct {
	data UserMessages
	err  error
}

// GetMessagesForUsers fetches messages for multiple users concurrently.
//
// Goroutines: each user ID runs in its own lightweight thread scheduled by the Go runtime.
// Channels: goroutines send fetchResult on a buffered channel; the caller collects results.
// Errors from any goroutine are aggregated and returned up the service → handler chain.
//
// When goroutines finish, their stack frames and channel payloads become unreachable.
// Go's garbage collector reclaims that memory automatically — under load this avoids
// manual free/pool management and keeps latency predictable compared to GC-less heaps
// that defer cleanup until memory pressure spikes.
func (s *MessageService) GetMessagesForUsers(ctx context.Context, userIDs []uint) ([]UserMessages, error) {
	if len(userIDs) == 0 {
		return nil, apperrors.ErrInvalidInput
	}

	results := make(chan fetchResult, len(userIDs))
	var wg sync.WaitGroup

	for _, id := range userIDs {
		wg.Add(1)
		go func(userID uint) {
			defer wg.Done()

			messages, err := s.messages.ListByUserID(ctx, userID)
			results <- fetchResult{
				data: UserMessages{UserID: userID, Messages: messages},
				err:  err,
			}
		}(id)
	}

	// Close results only after all goroutines complete — avoids sending on closed channel.
	go func() {
		wg.Wait()
		close(results)
	}()

	var collected []UserMessages
	var errs []error

	for res := range results {
		if res.err != nil {
			errs = append(errs, fmt.Errorf("user %d: %w", res.data.UserID, res.err))
			continue
		}
		collected = append(collected, res.data)
	}

	if len(errs) > 0 {
		return collected, fmt.Errorf("message service concurrent fetch: %v", errs)
	}
	return collected, nil
}
