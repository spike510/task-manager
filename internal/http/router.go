package http

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spike510/task-manager/internal/user"
)

func NewRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// Handlers
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := user.NewHandler(userService)

	// Public endpoints
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
	r.POST("/users/register", userHandler.Register)
	r.POST("/users/login", userHandler.Login)

	// Private endpoints
	auth := r.Group("/auth", AuthMiddleware())
	auth.GET("/me", func(c *gin.Context) {
		userID := c.GetInt("user_id")
		c.JSON(200, gin.H{"user_id": userID})
	})

	return r
}
