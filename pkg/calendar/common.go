/*
 * @Author: liangdong09
 * @Date: 2022-10-04 13:31:46
 * @LastEditTime: 2022-10-04 21:32:32
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

func GetGapTime(date Date) int {
	dayStr := fmt.Sprintf("%d%02d%02d", date.Year, int(date.Month), date.Day)
	fmt.Println(dayStr)
	h, err := time.ParseInLocation("20060102", dayStr, time.Local)
	if err != nil {
		return -1
	}
	t := time.Until(h).Hours()/24 + 1
	return int(t)
}
