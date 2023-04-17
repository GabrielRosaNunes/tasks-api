package main

import "github.com/gin-gonic/gin"

func defineRoutes() {
	router := gin.Default()
	router.GET("/", getTasks)
	router.GET("/:id",getTask)
	router.POST("/",insertTask)
	router.PUT("/:id",updateTask)
	router.DELETE("/:id",deleteTask)
	router.Run("localhost:8080")
}