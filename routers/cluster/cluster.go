package cluster

import (
	"krm-backend/controllers/cluster"

	"github.com/gin-gonic/gin"
)

func ClusterRouters(r *gin.RouterGroup) {
	r.GET("/cluster", cluster.GetClusterList)
	r.GET("/cluster/:id", cluster.GetCluster)
	r.POST("/cluster", cluster.PostCluster)
	r.PUT("/cluster/:id", cluster.UpdateCluster)
	r.DELETE("/cluster/:id", cluster.DeleteCluster)

}
