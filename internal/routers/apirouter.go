package routers

import (
	"github.com/gin-gonic/gin"
	c "github.com/liangdong/my-gin/internal/controller/v1"
)

func setApiRoute(r *gin.Engine) {
	// version 1
	v1 := r.Group("/api/v1")
	{
		//v1.POST("/login", c.Login)
		// v1.GET("/hello-world", c.HelloWorld)
		//v1.GET("/", c.HelloWorld)
		v1.POST("/upload", c.SubmitDates)
	}
}
