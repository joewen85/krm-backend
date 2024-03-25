package secret

import (
	"krm-backend/controllers"
	"krm-backend/utils/kubeutils"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
)

func List(c *gin.Context) {
	var secret corev1.Secret
	var info controllers.Info
	info.Item = &secret

	kuberconfig := controllers.NewInfo(c, &info, "获取成功")

	var kubeUtilsInterface kubeutils.KubeUtilsInterface
	resource := kubeutils.NewSecret(kuberconfig, &secret)
	kubeUtilsInterface = resource
	info.List(c, kubeUtilsInterface)
}