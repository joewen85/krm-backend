package namespace

import (
	"context"
	"krm-backend/config"
	"krm-backend/controllers"
	"krm-backend/utils/logs"
	"net/http"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Delete(c *gin.Context) {
	clientSet, baseInfo, err := controllers.BaseInit(c, nil)
	returnData := config.NewReturnData()
	if err != nil {
		returnData.Status = 500
		returnData.Message = err.Error()
		c.JSON(http.StatusOK, returnData)
		return
	}
	resourceName := baseInfo.Name
	if resourceName == "kube-system" {
		logs.Deubg(nil, "不能删除kubernetes核心命名空间")
		returnData.Status = 403
		returnData.Message = "禁止删除kube-system"
		c.JSON(http.StatusOK, returnData)
		return
	}
	err = clientSet.CoreV1().Namespaces().Delete(context.TODO(), resourceName, metav1.DeleteOptions{})
	if err != nil {
		logs.Error(map[string]interface{}{"error": err.Error()}, "删除资源失败")
		returnData.Status = 500
		returnData.Message = err.Error()
	} else {
		returnData.Status = 200
		returnData.Message = "删除资源成功"
		logs.Deubg(map[string]interface{}{"data": baseInfo}, "删除资源成功")
	}
	c.JSON(http.StatusOK, returnData)
}
