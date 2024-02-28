package namespace

import (
	"context"
	"fmt"
	"krm-backend/config"
	"krm-backend/utils/logs"
	"net/http"

	"github.com/gin-gonic/gin"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type ClusterInfo struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	City        string `json:"city"`
	District    string `json:"district"`
}

type ClusterConfig struct {
	ClusterInfo
	KuberConfig string `json:"kuberConfig"`
}

type ClusterStatus struct {
	ClusterInfo
	Version string `json:"version"`
	Status  string `json:"status"`
}

// 结构体的方法, 判断集群是否可用
func (c *ClusterConfig) GetClusterStatus() (ClusterStatus, error) {
	clusterStatus := ClusterStatus{}
	clusterStatus.ClusterInfo = c.ClusterInfo
	restconfig, err := clientcmd.RESTConfigFromKubeConfig([]byte(c.KuberConfig))
	if err != nil {
		return clusterStatus, err
	}
	clientset, err := kubernetes.NewForConfig(restconfig)
	if err != nil {
		return clusterStatus, err
	}
	serverVersion, err := clientset.Discovery().ServerVersion()
	if err != nil {
		return clusterStatus, err
	}
	version := serverVersion.String()
	clusterStatus.Version = version
	clusterStatus.Status = "Active"
	return clusterStatus, nil
}

func GetNamespaceList(c *gin.Context) {
	// clusterInfo := ClusterInfo{}
	returnData := config.NewReturnData()
	// if err := c.ShouldBindJSON(&clusterInfo); err != nil {
	// 	logs.Error(map[string]interface{}{"params": clusterInfo}, "获取集群列表失败")
	// 	returnData.Status = 400
	// 	returnData.Message = err.Error()
	// 	c.JSON(http.StatusOK, returnData)
	// 	return
	// }
	listOptions := metav1.ListOptions{}
	listOptions.LabelSelector = "apps=krm-backend"
	secretList, err := config.InClusterClient.CoreV1().Secrets(config.MetaDataNamespace).List(context.TODO(), listOptions)
	if err != nil {
		logs.Error(nil, "获取集群信息列表错误")
		returnData.Status = 400
		returnData.Message = err.Error()
		c.JSON(http.StatusOK, returnData)
		return
	}
	var clusterList []map[string]string

	for _, v := range secretList.Items {
		clusterList = append(clusterList, v.Annotations)
		fmt.Print(clusterList)
	}
	returnData.Data["items"] = clusterList
	c.JSON(http.StatusOK, returnData)
}

func GetNamespace(c *gin.Context) {
	id := c.Param("id")
	returnData := config.NewReturnData()
	clusterObj, err := config.InClusterClient.CoreV1().Secrets(config.MetaDataNamespace).Get(context.TODO(), id, metav1.GetOptions{})
	if err != nil {
		logs.Error(map[string]interface{}{"error": err.Error()}, "没有集群数据")
		returnData.Status = 400
		returnData.Message = "没有" + id + "资源"
		c.JSON(http.StatusOK, returnData)
		return
	} else {
		fmt.Print(clusterObj)
		returnData.Status = 200
		data := clusterObj.Annotations
		data["kubeconfig"] = string(clusterObj.Data["KuberConfig"])
		fmt.Print(data)
		returnData.Data["item"] = clusterObj.Annotations
	}

	c.JSON(http.StatusOK, returnData)
}

func PostNamespace(c *gin.Context) {
	CreateOrUpdate(c, "Create")
}

func UpdateNamespace(c *gin.Context) {
	CreateOrUpdate(c, "Update")
}
func DeleteNamespace(c *gin.Context) {
	clusterId := c.Param("id")
	returnData := config.NewReturnData()
	err := config.InClusterClient.CoreV1().Secrets(config.MetaDataNamespace).Delete(context.TODO(), clusterId, metav1.DeleteOptions{})
	if err != nil {
		logs.Error(map[string]interface{}{"clusterID": clusterId, "error": err.Error()}, "删除集群信息失败")
		returnData.Status = 204
		returnData.Message = "删除失败"
	} else {
		logs.Info(map[string]interface{}{"clusterID": clusterId}, "删除集群信息成功")
		returnData.Status = 200
		returnData.Message = "删除成功"
	}
	c.JSON(http.StatusOK, returnData)
}
