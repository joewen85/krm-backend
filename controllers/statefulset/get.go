package statefulset

import (
	"krm-backend/controllers"
	"krm-backend/utils/kubeutils"

	"github.com/gin-gonic/gin"
	appsv1 "k8s.io/api/apps/v1"
)

func Get(c *gin.Context) {
	var statefulSet appsv1.StatefulSet
	var info controllers.Info
	info.Item = &statefulSet

	kuberconfig := controllers.NewInfo(c, &info, "获取成功")

	var kubeUtilsInterface kubeutils.KubeUtilsInterface
	resource := kubeutils.NewStatefulSet(kuberconfig, &statefulSet)
	kubeUtilsInterface = resource
	info.Get(c, kubeUtilsInterface)
}
