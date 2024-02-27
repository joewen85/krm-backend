package cluster

import (
	"context"
	"krm-backend/config"
	"krm-backend/utils"
	"krm-backend/utils/logs"
	"net/http"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func CreateOrUpdate(c *gin.Context, method string) {
	var arg string
	if method == "Create" {
		arg = "创建"
	} else {
		arg = "修改"
	}

	clusterConfig := ClusterConfig{}
	returnData := config.NewReturnData()
	if err := c.ShouldBindJSON(&clusterConfig); err != nil {
		logs.Error(map[string]interface{}{"error": err.Error()}, "参数与结构体不一致")
		returnData.Status = 400
		returnData.Message = arg + "集群数据不全" + err.Error()
		c.JSON(http.StatusOK, returnData)
		return
	}
	clusterStatus, err := clusterConfig.GetClusterStatus()
	if err != nil {
		returnData.Status = 500
		msg := "kuberconfig参数错误" + err.Error()
		returnData.Message = msg
		logs.Error(map[string]interface{}{"error": err.Error()}, "提交kuberconfig有问题")
		c.JSON(http.StatusOK, returnData)
		return
	}

	logs.Info(map[string]interface{}{"data": clusterConfig}, "请求参数")
	// jsonStr, err := json.Marshal(clusterConfig)
	// if err != nil {
	// 	logs.Error(map[string]interface{}{"data": clusterConfig}, "结构体转JSON失败")
	// }
	// fmt.Print("json: ", string(jsonStr))

	// 创建secret
	// secret *v1.Secret
	var secretData corev1.Secret
	secretData.Name = clusterConfig.Id
	secretData.Labels = make(map[string]string)
	secretData.Labels["apps"] = "krm-backend"
	secretData.Annotations = make(map[string]string)
	clusterStatusMap := utils.StructToMap(clusterStatus)
	secretData.Annotations = clusterStatusMap

	secretData.StringData = make(map[string]string)
	secretData.StringData["KuberConfig"] = clusterConfig.KuberConfig
	if method == "Create" {
		_, err = config.InClusterClient.CoreV1().Secrets(config.MetaDataNamespace).Create(context.TODO(), &secretData, metav1.CreateOptions{})
	} else {
		_, err = config.InClusterClient.CoreV1().Secrets(config.MetaDataNamespace).Update(context.TODO(), &secretData, metav1.UpdateOptions{})
	}
	if err != nil {
		logs.Error(map[string]interface{}{"cluster": clusterConfig.Id, "clusterName": clusterConfig.DisplayName, "msg": err.Error()}, arg+"secret失败")
		returnData.Status = 500
		returnData.Message = err.Error()
		c.JSON(http.StatusOK, returnData)
		return
	}
	// m := utils.StructToReturnData(clusterConfig)
	returnData.Status = 201
	returnData.Message = arg + "成功"
	// TODO 返回请求数据, 后面改为返回创建资源的json
	// returnData.Data = m
	c.JSON(http.StatusOK, returnData)
}
