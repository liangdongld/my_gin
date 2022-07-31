/*
 * @Author: liangdong09
 * @Date: 2022-07-24 00:58:34
 * @LastEditTime: 2022-07-31 15:12:03
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/internal/routers/weChatRouter.go
 */
package routers

import (
	"github.com/gin-gonic/gin"
	w "github.com/liangdong/my-gin/internal/controller/wechat"
)

func setWeChatRouter(r *gin.Engine) {
	// version 1
	v1 := r.Group("wechat")
	{
		// v1.GET("/send", w.SendMsg)
		v1.GET("/receive", w.VerifyMsg)
		v1.POST("/receive", w.ReceiveMsg)
	}
}
