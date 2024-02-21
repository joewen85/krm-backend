package user

import (
	"krm-backend/controllers/users"

	"github.com/gin-gonic/gin"
)

func UserRouters(r *gin.RouterGroup) {
	r.POST("/user/login", users.Login)
	r.GET("/user/logout", users.Logout)
	r.GET("/user", users.GetUserList)
	r.GET("/user/:id/details", users.GetUser)
	r.POST("/user", users.PostUser)
}
