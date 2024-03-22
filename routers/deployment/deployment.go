package deployment

import (
	"krm-backend/controllers/deployment"

	"github.com/gin-gonic/gin"
)

func DeploymentRouters(r *gin.RouterGroup) {
	r.GET("/deployment", deployment.List)
	r.GET("/deployment/:name", deployment.Get)
	r.POST("/deployment/", deployment.Create)
	r.PUT("/deployment/:name", deployment.Update)
	r.DELETE("/deployment/:name", deployment.Delete)
}
