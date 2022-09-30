/*
 * @Author: liangdong09
 * @Date: 2022-07-19 00:31:13
 * @LastEditTime: 2022-10-01 00:11:31
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/internal/middleware/logger.go
 */
package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/liangdong/my-gin/config"
	"github.com/liangdong/my-gin/data"
	sys_model "github.com/liangdong/my-gin/internal/model/system"
	log "github.com/liangdong/my-gin/internal/pkg/logger"
	"github.com/liangdong/my-gin/pkg/utils"
	"go.uber.org/zap"
)

type responseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w responseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w responseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// CustomLogger 接收gin框架默认的日志
func CustomLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		blw := &responseWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		cost := time.Since(c.GetTime("requestStartTime")).Microseconds()
		if config.Config.AppEnv != "production" {
			var msg sys_model.SysAccessHistories
			msg.IP = c.ClientIP()
			msg.Cost = float32(cost)
			msg.Status = c.Writer.Status()
			msg.Method = c.Request.Method
			msg.Path = path
			msg.Query = query
			msg.Ua = c.Request.UserAgent()
			msg.Errors = c.Errors.ByType(gin.ErrorTypePrivate).String()
			msg.Response = blw.body.Bytes()
			rsp := make(map[string]interface{})
			err := json.Unmarshal(msg.Response, &rsp)
			if err != nil {
				rsp["content"] = utils.ByteSliceToString(msg.Response)
				bt, _ := json.Marshal(rsp)
				msg.Response = bt
			}
			err = data.MysqlDB.Create(&msg).Error
			if err != nil {
				fmt.Println(err.Error())
				msg.Errors = err.Error()
			}

			// sendWcByte, _ := json.Marshal(msg)
			// sendWcStr := utils.ByteSliceToString(sendWcByte)
			// go func() {
			// 	service.SendWeChat(sendWcStr, "text")
			// }()

			log.Logger.Info(path,
				zap.Int("status", msg.Status),
				zap.String("method", msg.Method),
				zap.String("path", msg.Path),
				zap.String("query", msg.Query),
				zap.String("ip", msg.IP),
				zap.String("user-agent", msg.Ua),
				zap.String("errors", msg.Errors),
				zap.Float32("cost", msg.Cost),
				zap.ByteString("response", msg.Response),
			)
		}
	}
}
