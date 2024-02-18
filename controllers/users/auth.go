package users

import (
	"krm-backend/config"
	"krm-backend/utils/jwtauth"
	"krm-backend/utils/logs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	userinfo := UserInfo{}
	returnData := config.NewReturnData()
	if err := c.ShouldBindJSON(&userinfo); err != nil {
		returnData.Status = 200500
		returnData.Message = err.Error()
		c.JSON(http.StatusOK, returnData)
		return
	}
	//TODO 修改为读取数据库判断
	if userinfo.Username == config.Username && userinfo.Password == config.Password {
		token, err := jwtauth.GenerateJwt(userinfo.Username)
		if err != nil {
			returnData.Status = 200401
			returnData.Message = "生成token失败"
			c.JSON(http.StatusOK, returnData)
			return
		}
		returnData.Status = 200200
		returnData.Message = "登陆成功"
		returnData.Data = map[string]interface{}{"token": token}
		c.JSON(http.StatusOK, returnData)
	} else {
		returnData.Status = 200401
		returnData.Message = "账号或密码不对"
		c.JSON(http.StatusOK, returnData)
		return
	}

	logs.Deubg(map[string]interface{}{"username": userinfo.Username, "password": userinfo.Password}, "login ok")

	//returndata.Status = 200000
	//returndata.Message = ""
	//returndata.Data = map[string]interface{}{"username": userinfo.Username, "password": userinfo.Password}
	//c.JSON(http.StatusOK, returndata)
}

func Logout(c *gin.Context) {
	returnData := config.NewReturnData()
	returnData.Status = 200000
	returnData.Message = ""
	c.JSON(http.StatusOK, returnData)
}
