package main

import "github.com/gin-gonic/gin"

// SetupRouter registers routes and returns the engine.
//
// Pulling router setup out of main() is a common Go pattern: production main() and
// tests both call SetupRouter(deps) without duplicating route registration.
func SetupRouter(store Store) *gin.Engine {
	router := gin.Default()
	handler := NewHandler(store)

	router.POST("/devices", handler.CreateDevice)

	return router
}
