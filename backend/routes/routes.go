package routes

import (
	"database/sql"

	"github.com/TechCatsLab/startend/backend/controllers"
	"github.com/TechCatsLab/startend/backend/middlewares"
	"github.com/TechCatsLab/startend/backend/services/mysql"
	"github.com/gin-gonic/gin"
)

// Install all available routes.
func Install(router *gin.Engine, db *sql.DB) {
	server := &controllers.Server{
		Service: &mysql.TaskServiceImpl{
			DB: db,
		},
	}

	router.POST("/api/v1/user/login", middlewares.JWTMiddleware.LoginHandler)

	auth := router.Group("/api/v1/task")
	auth.Use(middlewares.JWTMiddleware.MiddlewareFunc())
	{
		auth.POST("/create", server.Create)
		auth.POST("/query/id", server.QueryById)
		auth.POST("/query/user", server.QueryByUserId)
		auth.POST("/query/user/status", server.QueryByStatus)
		auth.POST("/stop", server.Stop)
		auth.POST("/success", server.Success)
	}
}
