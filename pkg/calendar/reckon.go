/*
 * @Author: liangdong09
 * @Date: 2022-10-04 13:31:35
 * @LastEditTime: 2022-10-04 16:21:01
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/pkg/calendar/reckon.go
 */
package calendar

import (
	"sort"
	"time"
)

type Holiday struct {
	Name  string
	Year  int
	Month time.Month
	Day   int
}

var (
	holidayList = []Holiday{
		{Name: "元旦", Year: 2023, Month: time.January, Day: 1},
		{Name: "春节", Year: 2023, Month: time.January, Day: 22},
		{Name: "清明节", Year: 2023, Month: time.April, Day: 5},
		{Name: "劳动节", Year: 2023, Month: time.May, Day: 1},
		{Name: "端午节", Year: 2023, Month: time.June, Day: 22},
		{Name: "中秋节", Year: 2023, Month: time.September, Day: 29},
		{Name: "国庆节", Year: 2023, Month: time.October, Day: 1},
	}
)

func sortHoliday() {
	sort.SliceStable(holidayList, func(i, j int) bool {
		return holidayList[i].Day < holidayList[j].Day
	})
	sort.SliceStable(holidayList, func(i, j int) bool {
		return int(holidayList[i].Month) < int(holidayList[j].Month)
	})
	sort.SliceStable(holidayList, func(i, j int) bool {
		return holidayList[i].Year < holidayList[j].Year
	})
}

func GetNextHoliday() Holiday {
	sortHoliday()
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	i := 0
	for ; i < len(holidayList); i++ {
		if holidayList[i].Year < year {
			continue
		}
		if holidayList[i].Year == year &&
			holidayList[i].Month < month {
			continue
		}
		if holidayList[i].Year == year &&
			holidayList[i].Month == month &&
			holidayList[i].Day < day {
			continue
		}
		break
	}
	return holidayList[i]
}
