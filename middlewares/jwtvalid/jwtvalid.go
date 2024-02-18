package jwtvalid

import (
	"krm-backend/config"
	"krm-backend/utils/jwtauth"
	"krm-backend/utils/logs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtValid(c *gin.Context) {
	returnData := config.NewReturnData()
	requestURL := c.FullPath()
	logs.Deubg(map[string]interface{}{"请求地址": requestURL}, "")
	if requestURL == "/api/user/logout" || requestURL == "/api/user/login" {
		c.Next()
		return
	}
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		returnData.Status = 200401
		returnData.Message = "无登陆"
		c.JSON(http.StatusOK, returnData)
		c.Abort()
		return
	}
	claims, err := jwtauth.ParseJwt(token)
	if err != nil {
		returnData.Status = 200401
		returnData.Message = "token不合法"
		c.JSON(http.StatusOK, returnData)
		c.Abort()
		return
	}
	c.Set("claims", claims)
	c.Next()
}
