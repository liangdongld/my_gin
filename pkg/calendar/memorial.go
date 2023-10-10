/*
 * @Author: liangdong09
 * @Date: 2022-10-04 13:30:34
 * @LastEditTime: 2023-10-10 17:29:18
 * @LastEditors: LiangDong
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
		{Name: "在一起 & 牵手", Year: 2022, Month: time.September, Day: 10},
		{Name: "游青岛", Year: 2023, Month: time.March, Day: 3},
		{Name: "游深圳", Year: 2023, Month: time.May, Day: 1},
		{Name: "见家长", Year: 2023, Month: time.September, Day: 29},
		{Name: "游重庆", Year: 2023, Month: time.October, Day: 1},
	}
)

func GetMemorialDays() []Date {
	// 暂无需处理
	return MemorialDays
}
