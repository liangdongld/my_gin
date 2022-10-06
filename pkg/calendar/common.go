/*
 * @Author: liangdong09
 * @Date: 2022-10-04 13:31:46
 * @LastEditTime: 2022-10-06 14:23:18
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/pkg/calendar/common.go
 */
package calendar

import (
	"fmt"
	"time"
)

type Date struct {
	Name  string
	Year  int
	Month time.Month
	Day   int
}

func GetUntilTime(date Date) int {
	h, err := DayToTime(date)
	if err != nil {
		return -1
	}
	t := time.Until(h).Hours()/24 + 1
	return int(t)
}

func GetSinceTime(date Date) int {
	h, err := DayToTime(date)
	if err != nil {
		return -1
	}
	t := time.Since(h).Hours()/24 + 1
	return int(t)
}

func DayToTime(date Date) (time.Time, error) {
	dayStr := fmt.Sprintf("%d%02d%02d", date.Year, int(date.Month), date.Day)
	h, err := time.ParseInLocation("20060102", dayStr, time.Local)
	if err != nil {
		return time.Now(), err
	}
	return h, nil
}
