/*
 * @Author: liangdong09
 * @Date: 2022-10-04 21:38:55
 * @LastEditTime: 2022-10-04 22:03:40
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/internal/task/corn.go
 */
package task

import (
	"fmt"

	"github.com/liangdong/my-gin/internal/service"
	"github.com/liangdong/my-gin/pkg/calendar"
	"github.com/robfig/cron"
)

func InitTask() {
	c := cron.New()
	c.AddFunc("0 0 12 * * *", SendDayilyHolidayMsg)
	c.Start()
	// select {}
}

func SendDayilyHolidayMsg() {
	str := GenNextHolidayMsg()
	str = addAutoSendFlag(str)
	service.SendWeChat(str, "text", "panghu")
}

func addAutoSendFlag(str string) string {
	str = fmt.Sprintf("------自动播报------\n\n%s", str)
	return str
}

// GenNextHolidayMsg, 获取下一个日期的消息
func GenNextHolidayMsg() string {
	nextHoliday := calendar.GetNextHoliday()
	gapDays := calendar.GetGapTime(nextHoliday)
	retStr := fmt.Sprintf("下一个节假日: %s\n", nextHoliday.Name)
	retStr = fmt.Sprintf("%s日期: %d-%02d-%02d\n", retStr, nextHoliday.Year, nextHoliday.Month, nextHoliday.Day)
	retStr = fmt.Sprintf("%s距今: %d 天", retStr, gapDays)
	return retStr
}
