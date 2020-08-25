package routes

import (
	"winddies/sso-api/controllers/auth"
	"winddies/sso-api/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRoutes(app *gin.Engine) {
	auth := auth.New()
	routeAuth := app.Group("/auth")
	routeAuth.Use(middlewares.GetSession())
	{
		routeAuth.POST("/login", auth.Login)
		routeAuth.POST("/register", auth.Register)
		routeAuth.GET("/verify-login", auth.VerifyLogin)
	}

	route := app.Group("/auth")
	{
		route.GET("/verify-token", auth.VerifyToken)
		route.GET("/logout", middlewares.GetSession(), auth.Logout)
		route.GET("/logout/:globalId", auth.Logout)
	}

}
