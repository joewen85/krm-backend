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

func Update(c *gin.Context) {
	returnData := config.NewReturnData()
	var ns corev1.Namespace
	clientSet, _, err := controllers.BaseInit(c, &ns)
	if err != nil {
		logs.Error(map[string]interface{}{"error": err.Error()}, "clientSet错误")
		returnData.Status = 500
		returnData.Message = err.Error()
		c.JSON(http.StatusOK, returnData)
		return
	}

	_, err = clientSet.CoreV1().Namespaces().Update(context.TODO(), &ns, metav1.UpdateOptions{})
	if err != nil {
		logs.Error(nil, "修改命名空间失败")
		returnData.Status = 500
		returnData.Message = err.Error()
	} else {
		returnData.Status = 200
		returnData.Message = "修改成功"
	}

	c.JSON(http.StatusOK, returnData)
}
