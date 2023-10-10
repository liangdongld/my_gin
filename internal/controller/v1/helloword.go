/*
 * @Author: liangdong09
 * @Date: 2022-07-23 21:03:55
 * @LastEditTime: 2022-11-19 16:48:12
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/internal/controller/v1/helloword.go
 */
package v1

import (
	"fmt"

	"github.com/liangdong/my-gin/data"
	sys_model "github.com/liangdong/my-gin/internal/model/system"

	"github.com/gin-gonic/gin"
	r "github.com/liangdong/my-gin/internal/pkg/response"
)

// HelloWorld hello world
func HelloWorld(c *gin.Context) {
	err := data.MysqlDB.AutoMigrate(&sys_model.SysUserDates{})
	if err != nil {
		r.Fail(c, 500, fmt.Sprintf("%s", err.Error()))
		return
	}
	str, ok := c.GetQuery("name")
	if !ok {
		str = "gin-layout"
	}
	r.Success(c, fmt.Sprintf("hello %s", str))
}

func SubmitDates(c *gin.Context) {
	var submitDate sys_model.SysUserDates
	err := c.ShouldBindJSON(&submitDate)
	if err != nil {
		r.Fail(c, 500, fmt.Sprintf("%s", err.Error()))
		return
	}
	data.MysqlDB.Create(&submitDate)
	r.Success(c, fmt.Sprintf("success"))
}
