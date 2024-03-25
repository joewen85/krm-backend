package daemonset

import (
	"krm-backend/controllers"
	"krm-backend/utils/kubeutils"

	"github.com/gin-gonic/gin"
	appsv1 "k8s.io/api/apps/v1"
)

func Create(c *gin.Context) {
	var daemonset appsv1.DaemonSet
	var info controllers.Info
	info.Item = &daemonset

	kuberconfig := controllers.NewInfo(c, &info, "创建成功")

	var kubeUtilsInterface kubeutils.KubeUtilsInterface
	resource := kubeutils.NewDaemonSet(kuberconfig, &daemonset)
	kubeUtilsInterface = resource
	info.Create(c, kubeUtilsInterface)
}
