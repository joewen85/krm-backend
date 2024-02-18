package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	url := c.FullPath()
	fmt.Print("请求: ", url)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
