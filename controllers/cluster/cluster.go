package cluster

import (
	"encoding/json"
	"fmt"
	"krm-backend/config"
	"krm-backend/utils/logs"
	"net/http"

	"github.com/gin-gonic/gin"
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

func GetClusterList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "get cluster list",
	})
}

func GetCluster(c *gin.Context) {
	id := c.Param("id")

	fmt.Printf("id: %s", id)
	c.JSON(http.StatusOK, gin.H{
		"msg": "get cluster detail",
		"id":  id,
	})
}

func PostCluster(c *gin.Context) {
	clusterConfig := ClusterConfig{}
	returnData := config.NewReturnData()
	if err := c.ShouldBindJSON(&clusterConfig); err != nil {
		fmt.Printf("struck: %v", &clusterConfig)
		returnData.Status = 400
		returnData.Message = "集群数据不全" + err.Error()
		c.JSON(http.StatusOK, returnData)
		return
	}
	logs.Info(map[string]interface{}{"data": clusterConfig}, "请求参数")
	jsonStr, err := json.Marshal(clusterConfig)
	if err != nil {
		logs.Error(map[string]interface{}{"data": clusterConfig}, "结构体转JSON失败")
	}
	fmt.Print("json: ", string(jsonStr))
	returnData.Status = 201
	returnData.Message = "添加成功"
	returnData.Data = jsonStr
	c.JSON(201, returnData)

}

func UpdateCluster(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "update cluster",
	})
}
func DeleteCluster(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "delete cluster",
	})
}
