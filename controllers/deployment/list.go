package deployment

import (
	"krm-backend/controllers"
	"krm-backend/utils/kubeutils"

	"github.com/gin-gonic/gin"

	appsv1 "k8s.io/api/apps/v1"
)

func List(c *gin.Context) {
	var deployment appsv1.Deployment
	var info controllers.Info
	info.Item = &deployment

	kuberconfig := controllers.NewInfo(c, &info, "获取成功")

	var kubeUtilsInterface kubeutils.KubeUtilsInterface
	resource := kubeutils.NewDeployment(kuberconfig, &deployment)
	kubeUtilsInterface = resource
	info.List(c, kubeUtilsInterface)
}
