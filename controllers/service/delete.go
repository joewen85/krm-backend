package service

import (
	"krm-backend/controllers"
	"krm-backend/utils/kubeutils"

	"github.com/gin-gonic/gin"
	corev1 "k8s.io/api/core/v1"
)

func Delete(c *gin.Context) {
	var service corev1.Service
	var info controllers.Info
	info.Item = &service

	kuberconfig := controllers.NewInfo(c, &info, "删除成功")

	var kubeUtilsInterface kubeutils.KubeUtilsInterface
	resource := kubeutils.NewService(kuberconfig, &service)
	kubeUtilsInterface = resource
	info.Delete(c, kubeUtilsInterface)
}
