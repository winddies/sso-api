package main

import (
	"net/http"
	"time"
	"winddies/sso-api/global"
	"winddies/sso-api/middlewares"
	"winddies/sso-api/models"
	"winddies/sso-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	global.Init("config.yml")
	models.Init()

	gin.SetMode(getGinMode())
	app := gin.New()
	app.Use(middlewares.Logger(), gin.Recovery())
	routes.InitRoutes(app)

	s := &http.Server{
		Addr:           global.Conf.Port,
		Handler:        app,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func getGinMode() string {
	switch global.Conf.Mode {
	case global.DevMode:
		return gin.DebugMode
	case global.TestMode:
		return gin.TestMode
	case global.ProdMode:
		return gin.ReleaseMode
	default:
		return gin.DebugMode
	}
}
