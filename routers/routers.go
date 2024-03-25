package routers

import (
	"krm-backend/routers/cluster"
	"krm-backend/routers/configmap"
	"krm-backend/routers/deployment"
	"krm-backend/routers/namespace"
	"krm-backend/routers/pod"
	"krm-backend/routers/secret"
	"krm-backend/routers/service"
	"krm-backend/routers/statefulset"
	"krm-backend/routers/user"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(r *gin.Engine) {
	apiGroup := r.Group("/api")
	user.UserRouters(apiGroup)
	cluster.ClusterRouters(apiGroup)
	namespace.NamespaceRouters(apiGroup)
	pod.PodRouters(apiGroup)
	deployment.DeploymentRouters(apiGroup)
	service.ServiceRouters(apiGroup)
	configmap.ConfigMapRouters(apiGroup)
	secret.SecretRouters(apiGroup)
	statefulset.StateFulSetRouters(apiGroup)
}
