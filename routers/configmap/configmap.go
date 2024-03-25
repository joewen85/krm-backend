package configmap

import (
	"krm-backend/controllers/configmap"

	"github.com/gin-gonic/gin"
)

func ConfigMapRouters(r *gin.RouterGroup) {
	r.GET("/configmap", configmap.List)
	r.GET("/configmap/:name", configmap.Get)
	r.POST("/configmap", configmap.Create)
	r.PUT("/configmap/:name", configmap.Update)
	r.DELETE("/configmap/:name", configmap.Delete)
}
