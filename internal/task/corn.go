/*
 * @Author: liangdong09
 * @Date: 2022-10-04 21:38:55
 * @LastEditTime: 2022-10-06 14:38:32
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/internal/task/corn.go
 */
package task

import (
	"fmt"

	"github.com/liangdong/my-gin/internal/service"
	"github.com/liangdong/my-gin/internal/service/receive"
	"github.com/robfig/cron"
)

func InitTask() {
	c := cron.New()
	c.AddFunc("0 0 12 * * *", SendDayilyHolidayMsg)
	c.AddFunc("0 0 8 * * *", SendDayilyPeriodMsg)
	c.Start()
}

func SendDayilyHolidayMsg() {
	str := receive.GenNextHolidayMsg()
	str = addAutoSendFlag(str)
	service.SendWeChat(str, "text", "panghu")
}

func SendDayilyPeriodMsg() {
	period := receive.GenNextPeriodMsg()
	period = addAutoSendFlag(period)
	service.SendWeChat(period, "text", "panghu")
}

func addAutoSendFlag(str string) string {
	str = fmt.Sprintf("------自动播报------\n\n%s", str)
	return str
}
