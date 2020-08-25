package middlewares

import (
	"winddies/sso-api/global"
	"winddies/sso-api/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetSession() gin.HandlerFunc {
	return sessions.Sessions(global.Conf.Session.Name, models.RedisSessionStore)
}
