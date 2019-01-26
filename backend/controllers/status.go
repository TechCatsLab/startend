package controllers

import (
	"github.com/fengyfei/go-box/system"
	"github.com/gin-gonic/gin"
)

type Server struct{}

func (s *Server) Status(c *gin.Context) {
	c.JSON(200, system.Info())
}
