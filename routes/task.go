package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/smartcraze/daily-dashboard/controllers"
	"gorm.io/gorm"
)

func TaskRoutes(app *gin.Engine, db *gorm.DB) {
	task := app.Group("/tasks")

	{
		task.GET("/", controllers.GetTask(db))
		task.POST("/", controllers.CreateTask(db))
		task.GET("/:id", controllers.GetTaskByID(db))
		task.PUT("/:id", controllers.UpdateTask(db))
		task.DELETE("/:id", controllers.DeleteTask(db))
	}
}
