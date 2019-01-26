package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	"github.com/TechCatsLab/startend/backend/middlewares"
	"github.com/TechCatsLab/startend/backend/routes"
	"github.com/TechCatsLab/startend/backend/services"
)

func main() {
	router := gin.Default()

	services.Load(databaseConn)

	middlewares.Install(router)
	routes.Install(router)

	s := &http.Server{
		Addr:    configuration.Address,
		Handler: router,
	}

	if err := s.ListenAndServe(); err != nil {
		glog.Fatalf("server error %s", err)
	}
}
