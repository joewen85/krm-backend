package secret

import (
	"krm-backend/controllers/secret"

	"github.com/gin-gonic/gin"
)

func SecretRouters(r *gin.RouterGroup) {
	r.GET("/secret", secret.List)
	r.GET("/secret/:name", secret.Get)
	r.POST("/secret", secret.Create)
	r.PUT("/secret/:name", secret.Update)
	r.DELETE("/secret/:name", secret.Delete)
}
