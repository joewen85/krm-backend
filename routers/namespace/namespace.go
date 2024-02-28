package namespace

import (
	"krm-backend/controllers/namespace"

	"github.com/gin-gonic/gin"
)

func NamespaceRouters(r *gin.RouterGroup) {
	r.GET("/namespace", namespace.GetNamespaceList)
	r.GET("/namespace/:id", namespace.GetNamespace)
	r.POST("/namespace", namespace.PostNamespace)
	r.PUT("/namespace/:id", namespace.UpdateNamespace)
	r.DELETE("/namespace/:id", namespace.DeleteNamespace)

}
