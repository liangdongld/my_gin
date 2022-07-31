/*
 * @Author: liangdong09
 * @Date: 2022-07-31 01:39:24
 * @LastEditTime: 2022-07-31 02:47:11
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/internal/model/system/sys_access_histories.go
 */
package sys_model

import (
	"github.com/wannanbigpig/gin-layout/internal/model"
	"gorm.io/gorm"
)

type SysAccessHistories struct {
	gorm.Model
	IP       string     `json:"ip" gorm:"comment:ip地址"`
	Path     string     `json:"path" gorm:"comment:访问路由"`
	Method   string     `json:"method" gorm:"comment:访问地址"`
	Query    string     `json:"query" gorm:"comment:请求体"`
	Ua       string     `json:"ua" gorm:"comment:user-agent"`
	Errors   string     `json:"errors" gorm:"comment:错误信息"`
	Cost     float32    `json:"cost" gorm:"comment:请求耗时"`
	Response model.JSON `json:"response" gorm:"comment:返回数据"`
	Status   int        `json:"status" gorm:"comment:状态码"`
}

func (SysAccessHistories) TableName() string {
	return "sys_access_histories"
}
