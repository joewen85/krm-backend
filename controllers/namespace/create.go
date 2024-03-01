package namespace

import (
	"context"
	"krm-backend/config"
	"krm-backend/controllers"
	"krm-backend/utils/logs"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Create(c *gin.Context) {
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

	_, err = clientSet.CoreV1().Namespaces().Create(context.TODO(), &resource, metav1.CreateOptions{})
	if err != nil {
		logs.Error(nil, "创建命名空间失败")
		returnData.Status = 500
		returnData.Message = err.Error()
	} else {
		returnData.Status = 200
		returnData.Message = "创建成功"
	}

	c.JSON(http.StatusOK, returnData)
}
