/*
 * @Author: liangdong09
 * @Date: 2022-10-04 22:15:52
 * @LastEditTime: 2022-11-06 21:39:57
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/pkg/calendar/period.go
 */
package calendar

import (
	"fmt"
	"time"
)

var (
	PeriodDays = []Date{
		{Name: "姨妈", Year: 2022, Month: time.September, Day: 10},
		{Name: "姨妈", Year: 2022, Month: time.October, Day: 5},
	}
	// 间隔30天
	gapDay = 30
)

func PredictNextPeriod() Date {
	recentPeriod := PeriodDays[len(PeriodDays)-1]
	nextHours := fmt.Sprintf("%dh", gapDay*24)
	m, _ := time.ParseDuration(nextHours)
	h, err := DayToTime(recentPeriod)
	if err != nil {
		return Date{}
	}
	ret := h.Add(m)
	return Date{Year: ret.Year(), Month: ret.Month(), Day: ret.Day()}
}
