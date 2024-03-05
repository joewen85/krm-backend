package namespace

import (
	"context"
	"krm-backend/config"
	"krm-backend/controllers"
	"krm-backend/utils/logs"
	"net/http"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Get(c *gin.Context) {
	returnData := config.NewReturnData()
	clientSet, baseInfo, err := controllers.BaseInit(c, nil)
	if err != nil {
		logs.Error(map[string]interface{}{"error": err.Error()}, "clientSet错误")
		returnData.Status = 500
		returnData.Message = err.Error()
		c.JSON(http.StatusOK, returnData)
		return
	}

	var resource corev1.Namespace
	resource.Name = baseInfo.Name

	namespace, err := clientSet.CoreV1().Namespaces().Get(context.TODO(), baseInfo.Name, metav1.GetOptions{})
	if err != nil {
		logs.Error(nil, "获取命名空间失败")
		returnData.Status = 500
		returnData.Message = err.Error()
	} else {
		returnData.Status = 200
		returnData.Message = "获取成功"
		returnData.Data["item"] = namespace
	}

	c.JSON(http.StatusOK, returnData)
}
