package main

import (
	"krm-backend/config"
	"krm-backend/middlewares/jwtvalid"
	"krm-backend/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(jwtvalid.JwtValid)
	routers.RegisterRouters(r)
	r.Run(config.Port)
}
