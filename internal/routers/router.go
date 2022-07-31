/*
 * @Author: liangdong09
 * @Date: 2022-07-19 00:31:13
 * @LastEditTime: 2022-07-31 15:11:38
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/internal/routers/router.go
 */
package routers

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liangdong/my-gin/internal/middleware"
	"github.com/liangdong/my-gin/internal/pkg/error_code"
	response2 "github.com/liangdong/my-gin/internal/pkg/response"
	"github.com/liangdong/my-gin/config"
)

func SetRouters() *gin.Engine {
	var r *gin.Engine

	if config.Config.Debug == false {
		// 生产模式
		r = ReleaseRouter()
		r.Use(
			middleware.RequestCostHandler(),
			middleware.CustomLogger(),
			middleware.CustomRecovery(),
			middleware.CorsHandler(),
		)
	} else {
		// 开发调试模式
		r = gin.New()
		r.Use(
			middleware.RequestCostHandler(),
			gin.Logger(),
			middleware.CustomRecovery(),
			middleware.CorsHandler(),
		)
	}

	// ping
	r.GET("/ping", func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"message": "pong!",
		})
	})

	// 设置 API 路由
	setApiRoute(r)
	// 设置微信路由
	setWeChatRouter(r)

	r.NoRoute(func(c *gin.Context) {
		response2.Resp().SetHttpCode(http.StatusNotFound).FailCode(c, error_code.NotFound)
	})

	return r
}

// ReleaseRouter 生产模式使用官方建议设置为 release 模式
func ReleaseRouter() *gin.Engine {
	// 切换到生产模式
	gin.SetMode(gin.ReleaseMode)
	// 禁用 gin 输出接口访问日志
	gin.DefaultWriter = ioutil.Discard

	engine := gin.New()

	return engine
}
