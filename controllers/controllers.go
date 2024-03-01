package controllers

import (
	"errors"
	"krm-backend/config"
	"krm-backend/utils/logs"

	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 基础结构
type BaseInfo struct {
	ClusterId string            `json:"clusterid" form:"clusterid"`
	Namespace string            `json:"namespace" form:"namespace"`
	Name      string            `json:"name" form:"name"`
	Labels    map[string]string `json:"labels" form:"labels"`
}

func BaseInit(c *gin.Context) (clientSet *kubernetes.Clientset, baseinfo BaseInfo, err error) {
	var baseInfo = BaseInfo{}
	if c.Request.Method == "GET" {
		// shouldBindQuery需要在结构体内添加form注解,否则无法绑定
		err = c.ShouldBindQuery(&baseInfo)
	} else {
		err = c.ShouldBindJSON(&baseInfo)
	}
	if err != nil {
		logs.Error(map[string]interface{}{"data": baseInfo}, "解析参数绑定失败")
		msg := "提交参数绑定失败" + err.Error()
		return nil, baseInfo, errors.New(msg)
	}

	clusterId := baseInfo.ClusterId
	clusterKuberConfig := config.ClusterList[clusterId]

	restConfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(clusterKuberConfig))
	if err != nil {
		logs.Error(map[string]interface{}{"error": err.Error()}, "kuberconfig错误")
		msg := "kuberconfig错误" + err.Error()
		return nil, baseInfo, errors.New(msg)
	}
	clientSet, err = kubernetes.NewForConfig(restConfig)
	if err != nil {
		logs.Error(map[string]interface{}{"error": err.Error()}, "创建ClientSet失败")
		msg := "创建ClientSet失" + err.Error()
		return nil, baseInfo, errors.New(msg)
	}
	return clientSet, baseInfo, nil
}
