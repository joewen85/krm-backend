package namespace

import (
	"context"
	"fmt"
	"krm-backend/config"
	"krm-backend/controllers"
	"krm-backend/utils/logs"
	"net/http"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func List(c *gin.Context) {
	returnData := config.NewReturnData()
	clientSet, baseInfo, err := controllers.BaseInit(c)
	if err != nil {
		logs.Error(map[string]interface{}{"error": err.Error()}, "clientSet错误")
		returnData.Status = 500
		returnData.Message = err.Error()
		c.JSON(http.StatusOK, returnData)
		return
	}

	var resource corev1.Namespace
	resource.Name = baseInfo.Name
	if len(baseInfo.Labels) > 1 {
		for k, v := range baseInfo.Labels {
			fmt.Print(k, ":", v)
		}
	}

	namespaceList, err := clientSet.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logs.Error(nil, "获取命名空间列表失败")
		returnData.Status = 500
		returnData.Message = err.Error()
	} else {
		returnData.Status = 200
		returnData.Message = "获取成功"
		returnData.Data["items"] = namespaceList.Items
	}

	c.JSON(http.StatusOK, returnData)
}
