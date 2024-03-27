package cronjob

import (
	"krm-backend/controllers/cronjob"

	"github.com/gin-gonic/gin"
)

func CronjobRouters(r *gin.RouterGroup) {
	r.GET("/cronjob", cronjob.List)
	r.GET("/cronjob/:name", cronjob.Get)
	r.POST("/cronjob/", cronjob.Create)
	r.PUT("/cronjob/:name", cronjob.Update)
	r.DELETE("/cronjob/:name", cronjob.Delete)
}
