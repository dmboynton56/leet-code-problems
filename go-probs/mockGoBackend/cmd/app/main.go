// Entry point for the mock Go backend.
//
// Run from this directory:
//   go run ./cmd/app
//
// Dependency management: go.mod pins module path and versions; go.sum records checksums.
// Add deps with `go get`; tidy with `go mod tidy`.
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"mockGoBackend/handlers"
	"mockGoBackend/internal/config"
	"mockGoBackend/models"
	"mockGoBackend/notifier"
	"mockGoBackend/repository"
	"mockGoBackend/services"
)

func main() {
	cfg := config.Load()

	// SQLite in-memory keeps the demo runnable without external Postgres/MySQL.
	// GORM maps structs to tables and generates SQL — repository layer stays thin.
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatalf("database: %v", err)
	}

	if err := db.AutoMigrate(&models.User{}, &models.Message{}); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	// Wire dependencies manually ("poor man's DI"). Larger projects often use wire or fx.
	userRepo := repository.NewUserRepository(db)
	msgRepo := repository.NewMessageRepository(db)

	// Notifier is an interface — swap EmailNotifier for SlackNotifier without changing UserService.
	emailNotifier := notifier.NewEmailNotifier()

	userSvc := services.NewUserService(userRepo, emailNotifier)
	msgSvc := services.NewMessageService(msgRepo, userRepo)

	userHandler := handlers.NewUserHandler(userSvc)
	msgHandler := handlers.NewMessageHandler(msgSvc)

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/users", userHandler.Create)
	r.GET("/users", userHandler.List)
	r.GET("/users/:id", userHandler.Get)

	r.POST("/users/:user_id/messages", msgHandler.Create)
	r.POST("/messages/batch", msgHandler.BatchGet)

	// Gin listens on a goroutine per connection internally; your handlers stay synchronous
	// unless you explicitly spawn work (see MessageService.GetMessagesForUsers).
	log.Printf("listening on :%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatalf("server: %v", err)
	}
}
