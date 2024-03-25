package service

import (
	"krm-backend/controllers/service"

	"github.com/gin-gonic/gin"
)

func ServiceRouters(r *gin.RouterGroup) {
	r.GET("/service", service.List)
	r.GET("/service/:name", service.Get)
	r.POST("/service", service.Create)
	r.PUT("/service/:name", service.Update)
	r.DELETE("/service/:name", service.Delete)
}
