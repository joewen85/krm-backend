package namespace

import (
	"krm-backend/controllers/namespace"

	"github.com/gin-gonic/gin"
)

func NamespaceRouters(r *gin.RouterGroup) {
	r.GET("/namespace", namespace.List)
	r.GET("/namespace/:name", namespace.Get)
	r.POST("/namespace/", namespace.Create)
	r.PUT("/namespace/:name", namespace.Update)
	r.DELETE("/namespace/:id", namespace.Delete)
}
