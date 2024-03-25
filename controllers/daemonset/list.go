package daemonset

import (
	"krm-backend/controllers"
	"krm-backend/utils/kubeutils"

	"github.com/gin-gonic/gin"

	appsv1 "k8s.io/api/apps/v1"
)

func List(c *gin.Context) {
	var daemonset appsv1.DaemonSet
	var info controllers.Info
	info.Item = &daemonset

	kuberconfig := controllers.NewInfo(c, &info, "获取成功")

	var kubeUtilsInterface kubeutils.KubeUtilsInterface
	resource := kubeutils.NewDaemonSet(kuberconfig, &daemonset)
	kubeUtilsInterface = resource
	info.List(c, kubeUtilsInterface)
}
