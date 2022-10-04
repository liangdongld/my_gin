/*
 * @Author: liangdong09
 * @Date: 2022-10-04 13:31:46
 * @LastEditTime: 2022-10-04 16:43:41
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/pkg/calendar/common.go
 */
package calendar

import (
	"fmt"
	"time"
)

type Holiday struct {
	Name  string
	Year  int
	Month time.Month
	Day   int
}

func GetGapTime(holiday Holiday) int {
	dayStr := fmt.Sprintf("%d%02d%02d", holiday.Year, int(holiday.Month), holiday.Day)
	fmt.Println(dayStr)
	h, err := time.ParseInLocation("20060102", dayStr, time.Local)
	if err != nil {
		return -1
	}
	t := time.Until(h).Hours()/24 + 1
	return int(t)
}
