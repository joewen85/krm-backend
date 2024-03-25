package configmap

import (
	"krm-backend/controllers"
	"krm-backend/utils/kubeutils"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
)

func Get(c *gin.Context) {
	var configmap corev1.ConfigMap
	var info controllers.Info
	info.Item = &configmap

	kuberconfig := controllers.NewInfo(c, &info, "获取成功")

	var kubeUtilsInterface kubeutils.KubeUtilsInterface
	resource := kubeutils.NewConfigMap(kuberconfig, &configmap)
	kubeUtilsInterface = resource
	info.Get(c, kubeUtilsInterface)
}
