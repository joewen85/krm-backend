package daemonset

import (
	"krm-backend/controllers/daemonset"

	"github.com/gin-gonic/gin"
)

func DaemonsetRouters(r *gin.RouterGroup) {
	r.GET("/daemonset", daemonset.List)
	r.GET("/daemonset/:name", daemonset.Get)
	r.POST("/daemonset/", daemonset.Create)
	r.PUT("/daemonset/:name", daemonset.Update)
	r.DELETE("/daemonset/:name", daemonset.Delete)
}
