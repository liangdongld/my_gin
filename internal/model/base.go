/*
 * @Author: liangdong09
 * @Date: 2022-07-19 00:31:13
 * @LastEditTime: 2022-07-31 15:10:48
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/internal/model/base.go
 */
/*
 * @Author: liangdong09
 * @Date: 2022-07-19 00:31:13
 * @LastEditTime: 2022-07-31 02:21:38
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/internal/model/base.go
 */
package model

import (
	"github.com/liangdong/my-gin/data"
	"github.com/liangdong/my-gin/internal/pkg/utils"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type BaseModel struct {
	ID        uint             `gorm:"column:id;type:int(11) unsigned AUTO_INCREMENT;not null;primarykey" json:"id"`
	CreatedAt utils.FormatDate `gorm:"column:created_at;type:timestamp;default:null" json:"created_at"`
	UpdatedAt utils.FormatDate `gorm:"column:updated_at;type:timestamp;default:null" json:"updated_at"`
}

func (model *BaseModel) DB() *gorm.DB {
	return DB()
}

type ContainsDeleteBaseModel struct {
	BaseModel
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:int(11) unsigned;not null;default:0;index" json:"-"`
}

func DB() *gorm.DB {
	return data.MysqlDB
}
