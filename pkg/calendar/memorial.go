/*
 * @Author: liangdong09
 * @Date: 2022-10-04 13:30:34
 * @LastEditTime: 2022-11-14 12:31:46
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/pkg/calendar/memorial.go
 */
package calendar

import "time"

var (
	MemorialDays = []Date{
		{Name: "宝宝出生", Year: 1997, Month: time.January, Day: 4},
		{Name: "乖乖出生", Year: 1997, Month: time.January, Day: 16},
		{Name: "第一次见面", Year: 2022, Month: time.August, Day: 7},
		{Name: "在一起", Year: 2022, Month: time.September, Day: 10},
		{Name: "第一次牵手", Year: 2022, Month: time.September, Day: 10},
	}
)

func GetMemorialDays() []Date {
	// 暂无需处理
	return MemorialDays
}
