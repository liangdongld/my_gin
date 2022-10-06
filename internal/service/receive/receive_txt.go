/*
 * @Author: liangdong09
 * @Date: 2022-08-05 19:41:59
 * @LastEditTime: 2022-10-06 14:37:03
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/internal/service/receive/receive_txt.go
 */
package receive

import (
	"fmt"
	"strings"

	"github.com/liangdong/my-gin/internal/model"
	"github.com/liangdong/my-gin/pkg/calendar"
)

type ReceiveTxt struct {
	Msg model.MsgContent
}

func (r *ReceiveTxt) ReplyMsg() (model.MsgContent, error) {
	if strings.Contains(r.Msg.Content, "lsp") ||
		strings.Contains(r.Msg.Content, "罗圣鹏") {
		r.Msg.Content = "罗圣鹏是你儿子"
	} else if strings.HasPrefix(r.Msg.Content, "假期") {
		r.Msg.Content = GenNextHolidayMsg()
	}
	return r.Msg, nil
}

// GenNextHolidayMsg, 获取下一个日期的消息
func GenNextHolidayMsg() string {
	nextHoliday := calendar.GetNextHoliday()
	gapDays := calendar.GetUntilTime(nextHoliday)
	retStr := fmt.Sprintf("下一个节假日: %s\n", nextHoliday.Name)
	retStr = fmt.Sprintf("%s日期: %d-%02d-%02d\n", retStr, nextHoliday.Year, nextHoliday.Month, nextHoliday.Day)
	retStr = fmt.Sprintf("%s距今: %d 天", retStr, gapDays)
	return retStr
}

func GenNextPeriodMsg() string {
	nextPeriod := calendar.PredictNextPeriod()
	gapDays := calendar.GetUntilTime(nextPeriod)
	if gapDays > 7 {
		return ""
	}
	retStr := fmt.Sprintf("下一次姨妈预计: %d-%02d-%02d\n", nextPeriod.Year, nextPeriod.Month, nextPeriod.Day)
	retStr = fmt.Sprintf("%s距今: %d 天", retStr, gapDays)
	return retStr
}
