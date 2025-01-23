package main

import (
	"task-management/controllers"
	"task-management/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDatabase()
	r := gin.Default()

	r.POST("/tasks", controllers.CreateTask)
	r.GET("/tasks/:id", controllers.GetTask)
	r.PUT("/tasks/:id", controllers.UpdateTask)
	r.DELETE("/tasks/:id", controllers.DeleteTask)
	r.GET("/tasks", controllers.ListTasks)

	r.Run(":8080")
}
