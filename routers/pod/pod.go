package pod

import (
	"krm-backend/controllers/pod"

	"github.com/gin-gonic/gin"
)

func PodRouters(r *gin.RouterGroup) {
	r.GET("/pod", pod.List)
	r.GET("/pod/:name", pod.Get)
	r.POST("/pod", pod.Create)
	r.PUT("/pod/:name", pod.Update)
	r.DELETE("/pod/:name", pod.Delete)
}
