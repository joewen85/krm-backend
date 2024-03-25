package controllers

import (
	"errors"
	"krm-backend/config"
	"krm-backend/utils/kubeutils"
	"krm-backend/utils/logs"
	"net/http"

	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// 基础结构
type BaseInfo struct {
	ClusterId  string      `json:"clusterid" form:"clusterid"`
	Namespace  string      `json:"namespace" form:"namespace"`
	Name       string      `json:"name" form:"name"`
	Item       interface{} `json:"item"`
	DeleteList []string    `json:"deletelist"`
}

type Info struct {
	BaseInfo
	ReturnData    config.ReturnData
	LabelSelector string `json:"labelselector"`
	FieldSelector string `json:"fieldselector"`
	Force         bool   `json:"force" form:"force"`
}

func NewInfo(c *gin.Context, info *Info, returnDataMsg string) (kuberconfig string) {
	var err error

	if c.Request.Method == "GET" {
		// shouldBindQuery需要在结构体内添加form注解,否则无法绑定
		err = c.ShouldBindQuery(&info)
	} else {
		err = c.ShouldBindJSON(&info)
	}
	if err != nil {
		logs.Error(map[string]interface{}{"data": info}, "解析参数绑定失败")
		msg := "提交参数绑定失败" + err.Error()
		info.ReturnData.Status = 500
		info.ReturnData.Message = msg
		c.JSON(http.StatusOK, info.ReturnData)
		return
	}
	info.ReturnData.Status = 200
	info.ReturnData.Message = returnDataMsg
	info.ReturnData.Data = make(map[string]interface{})
	kuberconfig = config.ClusterList[info.ClusterId]
	return kuberconfig
}

func (i *Info) Create(c *gin.Context, kubeUtilInterface kubeutils.KubeUtilsInterface) {
	err := kubeUtilInterface.Create(i.Namespace)
	if err != nil {
		msg := "创建失败" + err.Error()
		i.ReturnData.Message = msg
		i.ReturnData.Status = 500
		logs.Error(nil, msg)
	}
	c.JSON(http.StatusOK, i.ReturnData)
}

func (i *Info) Update(c *gin.Context, kubeUtilInterface kubeutils.KubeUtilsInterface) {
	err := kubeUtilInterface.Update(i.Namespace)
	if err != nil {
		msg := "更新失败" + err.Error()
		i.ReturnData.Message = msg
		i.ReturnData.Status = 500
		logs.Error(nil, msg)
	}
	c.JSON(http.StatusOK, i.ReturnData)
}

func (i *Info) Delete(c *gin.Context, kubeUtilInterface kubeutils.KubeUtilsInterface) {
	var err error
	var gracePeriodSeconds int64
	if i.Force {
		gracePeriodSeconds = 0
	}
	if i.DeleteList == nil {
		err = kubeUtilInterface.Delete(i.Namespace, i.Name, &gracePeriodSeconds)
	} else {
		err = kubeUtilInterface.DeleteList(i.Namespace, i.DeleteList, &gracePeriodSeconds)
	}

	if err != nil {
		msg := "删除失败" + err.Error()
		i.ReturnData.Message = msg
		i.ReturnData.Status = 500
		logs.Error(nil, msg)
	}
	c.JSON(http.StatusOK, i.ReturnData)
}

func (i *Info) List(c *gin.Context, kubeUtilInterface kubeutils.KubeUtilsInterface) {
	items, err := kubeUtilInterface.List(i.Namespace, i.LabelSelector, i.FieldSelector)
	if err != nil {
		msg := "查询失败" + err.Error()
		i.ReturnData.Message = msg
		i.ReturnData.Status = 500
		logs.Error(nil, msg)
	} else {
		i.ReturnData.Data["items"] = items
	}
	c.JSON(http.StatusOK, i.ReturnData)
}

func (i *Info) Get(c *gin.Context, kubeUtilInterface kubeutils.KubeUtilsInterface) {
	item, err := kubeUtilInterface.Get(i.Namespace, i.Name)
	if err != nil {
		msg := "查询失败" + err.Error()
		i.ReturnData.Message = msg
		i.ReturnData.Status = 500
		logs.Error(nil, msg)
	} else {
		i.ReturnData.Data["item"] = item
	}
	c.JSON(http.StatusOK, i.ReturnData)
}

func BaseInit(c *gin.Context, item interface{}) (clientSet *kubernetes.Clientset, baseinfo BaseInfo, err error) {
	baseInfo := BaseInfo{}
	baseInfo.Item = item
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
