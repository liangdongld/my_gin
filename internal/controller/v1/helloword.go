/*
 * @Author: liangdong09
 * @Date: 2022-07-23 21:03:55
 * @LastEditTime: 2022-09-18 19:55:59
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/internal/controller/v1/helloword.go
 */
package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	r "github.com/liangdong/my-gin/internal/pkg/response"
)

// HelloWorld hello world
func HelloWorld(c *gin.Context) {
	str, ok := c.GetQuery("name")
	if !ok {
		str = "gin-layout"
	}

	r.Success(c, fmt.Sprintf("hello %s", str))
}
