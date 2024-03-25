package statefulset

import (
	"krm-backend/controllers/statefulset"

	"github.com/gin-gonic/gin"
)

func StateFulSetRouters(r *gin.RouterGroup) {
	r.GET("/statefulset", statefulset.List)
	r.GET("/statefulset/:name", statefulset.Get)
	r.POST("/statefulset/", statefulset.Create)
	r.PUT("/statefulset/:name", statefulset.Update)
	r.DELETE("/statefulset/:name", statefulset.Delete)
}
