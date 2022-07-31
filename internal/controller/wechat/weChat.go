/*
 * @Author: liangdong09
 * @Date: 2022-07-23 20:20:23
 * @LastEditTime: 2022-07-31 12:10:38
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/internal/controller/wechat/weChat.go
 */
package wechat

import (
	"github.com/gin-gonic/gin"
	"github.com/wannanbigpig/gin-layout/internal/pkg/error_code"
	log "github.com/wannanbigpig/gin-layout/internal/pkg/logger"
	r "github.com/wannanbigpig/gin-layout/internal/pkg/response"
	"github.com/wannanbigpig/gin-layout/internal/service"
)

func SendMsg(c *gin.Context) {
	msg, ok := c.GetQuery("msg")
	if !ok {
		msg = "please input message"
	}
	log.Logger.Info("send wechat message: " + msg)
	err := service.SendWeChat(msg, "text")
	if err != nil {
		r.Resp().FailCode(c, error_code.FAILURE, err.Error())
		return
	}
	r.Success(c, "success")
}

func ReceiveMsg(c *gin.Context) {
	verifyMsgSign, _ := c.GetQuery("msg_signature")
	verifyTimestamp, _ := c.GetQuery("timestamp")
	verifyNonce, _ := c.GetQuery("nonce")
	b, _ := c.GetRawData()
	msg, err := service.ReceiveMsg(verifyMsgSign, verifyTimestamp, verifyNonce, b)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return
	}
	if msg == "" {
		msg = "NULL"
	}
	c.Writer.Write([]byte(msg))
}

func VerifyMsg(c *gin.Context) {
	verifyMsgSign, _ := c.GetQuery("msg_signature")
	verifyTimestamp, _ := c.GetQuery("timestamp")
	verifyNonce, _ := c.GetQuery("nonce")
	verifyEchoStr, _ := c.GetQuery("echostr")
	msg, err := service.VerifyMsg(verifyMsgSign, verifyTimestamp, verifyNonce, verifyEchoStr)
	if err != nil {
		log.Logger.Fatal(err.Error())
		return
	}
	if msg == "" {
		msg = "NULL"
		r.Fail(c, 200, "解析信息失败")
	}
	c.Writer.Write([]byte(msg))
}
