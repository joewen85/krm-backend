package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "get user obj",
	})
}

func PostUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "create user",
	})
}
