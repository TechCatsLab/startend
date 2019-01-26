package controllers

import (
	"time"

	"github.com/appleboy/gin-jwt"

	"github.com/golang/glog"

	"github.com/TechCatsLab/startend/backend/services/mysql"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Service *mysql.TaskServiceImpl
}

func (s *Server) Create(c *gin.Context) {
	var (
		req struct {
			User_id int
			Content string    `json:"content"`
			Started time.Time `json:"started"`
			Ended   time.Time `json:"ended"`
		}
	)
	if err := c.BindJSON(&req); err != nil {
		glog.Errorf("[status] parameter error:%s", err)

		c.JSON(400, "parameter error")
	}

	claims := jwt.ExtractClaims(c)
	req.User_id = claims["id"].(int)
	err := s.Service.Create(uint32(req.User_id), req.Content, req.Started, req.Ended)
	if err != nil {
		glog.Errorf("[error] Create in task : %s", err)

		c.JSON(400, err)
	}

	c.JSON(200, "insert ok")
}

func (s *Server) QueryById(c *gin.Context) {
	var (
		req struct {
			Id int `json:"id"`
		}
	)
	if err := c.BindJSON(&req); err != nil {
		glog.Errorf("[status] parameter error:%s", err)

		c.JSON(400, "parameter error")
	}

	task, err := s.Service.QueryById(uint32(req.Id))
	if err != nil {
		glog.Errorf("[error] query by id  : %s", err)

		c.JSON(400, err)
	}

	c.JSON(200, task)
}

func (s *Server) QueryByUserId(c *gin.Context) {
	var (
		req struct {
			User_id int
			Page    int `json:"page"`
		}
	)
	if err := c.BindJSON(&req); err != nil {
		glog.Errorf("[status] parameter error:%s", err)

		c.JSON(400, "parameter error")
	}

	claims := jwt.ExtractClaims(c)
	req.User_id = claims["id"].(int)
	tasks, err := s.Service.QueryByUserId(uint32(req.User_id), req.Page)
	if err != nil {
		glog.Errorf("[error] query by Userid : %s", err)

		c.JSON(400, err)
	}

	c.JSON(200, tasks)
}

func (s *Server) QueryByStatus(c *gin.Context) {
	var (
		req struct {
			User_id int
			Status  uint8 `json:"status"`
			Page    int   `json:"page"`
		}
	)
	if err := c.BindJSON(&req); err != nil {
		glog.Errorf("[status] parameter error:%s", err)
		c.JSON(400, "parameter error")
	}
	claims := jwt.ExtractClaims(c)
	req.User_id = claims["id"].(int)
	tasks, err := s.Service.QueryByStatus(uint32(req.User_id), req.Status, req.Page)
	if err != nil {
		glog.Errorf("[error] query by status : %s", err)

		c.JSON(400, err)
	}

	c.JSON(200, tasks)
}

func (s *Server) Stop(c *gin.Context) {
	var (
		req struct {
			Id      uint32 `json:"id"`
			Comment string `json:"comment"`
		}
	)
	if err := c.BindJSON(&req); err != nil {
		glog.Errorf("[status] parameter error:%s", err)

		c.JSON(400, "parameter error")
	}

	err := s.Service.Stop(req.Comment, req.Id)
	if err != nil {
		glog.Errorf("[error] stop a task : %s", err)

		c.JSON(400, err)
	}

	c.JSON(200, "stop a task")
}

func (s *Server) Success(c *gin.Context) {
	var (
		req struct {
			Id      uint32 `json:"id"`
			Comment string `json:"comment"`
		}
	)
	if err := c.BindJSON(&req); err != nil {
		glog.Errorf("[status] parameter error:%s", err)

		c.JSON(400, "parameter error")
	}

	err := s.Service.Success(req.Comment, req.Id)
	if err != nil {
		glog.Errorf("[error] success a task : %s", err)

		c.JSON(400, err)
	}

	c.JSON(200, "success a task")
}
