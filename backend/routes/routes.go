package routes

import (
	"github.com/TechCatsLab/startend/backend/controllers"
	"github.com/TechCatsLab/startend/backend/middlewares"
	"github.com/gin-gonic/gin"
)

// Install all available routes.
func Install(router *gin.Engine) {
	server := &controllers.Server{}

	router.POST("/api/v1/user/login", middlewares.JWTMiddleware.LoginHandler)

	router.POST("/api/v1/insights/status", server.Status)
}
