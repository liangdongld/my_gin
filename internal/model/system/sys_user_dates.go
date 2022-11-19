package sys_model

import (
	"gorm.io/gorm"
	"time"
)

type DateType int

const (
	Memorial DateType = iota
	Period
	Reckon
)

type SysUserDates struct {
	gorm.Model
	Name string    `json:"name"`
	Date time.Time `json:"date"`
	Type DateType  `json:"type"`
	Note string    `json:"note"`
}

func (SysUserDates) TableName() string {
	return "sys_user_dates"
}
