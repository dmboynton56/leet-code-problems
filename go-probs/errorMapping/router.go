package main

import "github.com/gin-gonic/gin"

func SetupRouter(store Store) *gin.Engine {
	router := gin.Default()
	handler := NewHandler(store)

	router.POST("/devices", handler.CreateDevice)

	return router
}
