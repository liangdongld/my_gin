package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/liangdong/my-gin/internal/pkg/error_code"
	"github.com/liangdong/my-gin/internal/pkg/logger"
	response2 "github.com/liangdong/my-gin/internal/pkg/response"
	"github.com/liangdong/my-gin/config"
)

// CustomRecovery 自定义错误 (panic) 拦截中间件、对可能发生的错误进行拦截、统一记录
func CustomRecovery() gin.HandlerFunc {
	DefaultErrorWriter := &PanicExceptionRecord{}
	return gin.RecoveryWithWriter(DefaultErrorWriter, func(c *gin.Context, err interface{}) {
		// 这里针对发生的panic等异常进行统一响应即
		// 这里针对发生的panic等异常进行统一响应即
		errStr := ""
		if config.Config.Debug == true {
			errStr = fmt.Sprintf("%v", err)
		}
		response2.Resp().SetHttpCode(http.StatusInternalServerError).FailCode(c, error_code.ServerError, errStr)
	})
}

// PanicExceptionRecord  panic等异常记录
type PanicExceptionRecord struct{}

func (p *PanicExceptionRecord) Write(b []byte) (n int, err error) {
	s1 := "An error occurred in the server's internal code："
	var build strings.Builder
	build.WriteString(s1)
	build.Write(b)
	errStr := build.String()
	logger.Logger.Error(errStr)
	return len(errStr), errors.New(errStr)
}
