/*
 * @Author: liangdong09
 * @Date: 2022-10-04 16:38:41
 * @LastEditTime: 2022-10-04 16:38:57
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/pkg/calendar/common_test.go
 */
package calendar

import "testing"

func TestGetGapDay(t *testing.T) {
	H := GetNextHoliday()
	i := GetUntilTime(H)
	t.Log(i)
}
