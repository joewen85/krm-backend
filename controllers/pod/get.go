package pod

import (
	"krm-backend/controllers"
	"krm-backend/utils/kubeutils"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
)

func Get(c *gin.Context) {
	var pod corev1.Pod
	var info controllers.Info
	info.Item = &pod

	kuberconfig := controllers.NewInfo(c, &info, "获取成功")

	var kubeUtilsInterface kubeutils.KubeUtilsInterface
	resource := kubeutils.NewPod(kuberconfig, &pod)
	kubeUtilsInterface = resource
	info.Get(c, kubeUtilsInterface)
}
