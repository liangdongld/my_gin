package routers

import (
	"github.com/gin-gonic/gin"
	w "github.com/wannanbigpig/gin-layout/internal/controller/wechat"
)

func setWeChatRouter(r *gin.Engine) {
	// version 1
	v1 := r.Group("wechat")
	{
		v1.GET("/send", w.SendMsg)
	}
}
