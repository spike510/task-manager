package http

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spike510/task-manager/internal/task"
	"github.com/spike510/task-manager/internal/user"

	_ "github.com/spike510/task-manager/docs" // swagger docs

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// Users
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	// Tasks
	taskRepository := task.NewRepository(db)
	taskService := task.NewService(taskRepository)
	taskHandler := task.NewHandler(taskService)

	// Public
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	r.POST("/users/register", userHandler.Register)
	r.POST("/users/login", userHandler.Login)

	// Authenticated
	auth := r.Group("/", AuthMiddleware())
	// Users
	auth.GET("/me", func(c *gin.Context) {
		userID := c.GetInt("user_id")
		c.JSON(200, gin.H{"user_id": userID})
	})

	// Tasks
	auth.GET("/tasks", taskHandler.GetTasks)
	auth.POST("/tasks", taskHandler.CreateTask)
	auth.PUT("/tasks/:id", taskHandler.UpdateTask)
	auth.DELETE("/tasks/:id", taskHandler.DeleteTask)

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
