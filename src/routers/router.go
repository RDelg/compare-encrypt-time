package routers

import (
	v1 "github.com/RDelg/compare-encrypt-time/src/routers/api/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/")
	{
		apiv1.POST("/", v1.RemoteFunctionAdapter)
		apiv1.POST("/encrypt", v1.Encrypt)
		apiv1.POST("/decrypt", v1.Encrypt)
	}

	return r
}
