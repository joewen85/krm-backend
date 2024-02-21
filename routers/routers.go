package routers

import (
	"krm-backend/routers/cluster"
	"krm-backend/routers/user"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(r *gin.Engine) {
	apiGroup := r.Group("/api")
	user.UserRouters(apiGroup)
	cluster.ClusterRouters(apiGroup)
}