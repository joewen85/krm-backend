package statefulset

import (
	"krm-backend/controllers"
	"krm-backend/utils/kubeutils"

	"github.com/gin-gonic/gin"
	appsv1 "k8s.io/api/apps/v1"
)

func Create(c *gin.Context) {
	// TODO: 创建sts时,定义创建service, 还是选择已存在的service绑定. handle service. serviceName参数
	var statefulSet appsv1.StatefulSet
	var info controllers.Info
	info.Item = &statefulSet

	kuberconfig := controllers.NewInfo(c, &info, "创建成功")

	var kubeUtilsInterface kubeutils.KubeUtilsInterface
	resource := kubeutils.NewStatefulSet(kuberconfig, &statefulSet)
	kubeUtilsInterface = resource
	info.Create(c, kubeUtilsInterface)
}
