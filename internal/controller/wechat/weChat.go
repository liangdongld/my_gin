package wechat

import (
	"github.com/gin-gonic/gin"
	"github.com/wannanbigpig/gin-layout/internal/pkg/error_code"
	r "github.com/wannanbigpig/gin-layout/internal/pkg/response"
	"github.com/wannanbigpig/gin-layout/internal/service"
)

func SendMsg(c *gin.Context) {
	msg, ok := c.GetQuery("msg")
	if !ok {
		msg = "please input message"
	}
	err := service.SendWeChat(msg, "text")
	if err != nil {
		r.Resp().FailCode(c, error_code.FAILURE, err.Error())
		return
	}
	r.Success(c, "success")
}
