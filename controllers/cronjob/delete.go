package cronjob

import (
	"krm-backend/controllers"
	"krm-backend/utils/kubeutils"

	"github.com/gin-gonic/gin"
	batchv1 "k8s.io/api/batch/v1"
)

func Delete(c *gin.Context) {
	var cronjob batchv1.CronJob
	var info controllers.Info
	info.Item = &cronjob

	kuberconfig := controllers.NewInfo(c, &info, "删除成功")

	var kubeUtilsInterface kubeutils.KubeUtilsInterface
	resource := kubeutils.NewCronjob(kuberconfig, &cronjob)
	kubeUtilsInterface = resource
	info.Delete(c, kubeUtilsInterface)
}
