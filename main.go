package main

import (
	"rebac/db"
	"rebac/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	r := gin.Default()

	// Model routes
	model := r.Group("/models")
	{
		model.POST("", handlers.CreateModel)
		model.GET("", handlers.GetModels)
		model.GET(":id", handlers.GetModel)
		model.PUT(":id", handlers.UpdateModel)
		model.DELETE(":id", handlers.DeleteModel)
	}

	r.Run() // default on :8080
}
