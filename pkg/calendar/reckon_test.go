/*
 * @Author: liangdong09
 * @Date: 2022-10-04 16:02:01
 * @LastEditTime: 2022-10-04 16:07:02
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/pkg/calendar/reckon_test.go
 */
package calendar

import (
	"testing"
)

func TestGetLasHoliday(t *testing.T) {
	h := GetNextHoliday()
	t.Log(h.Name)
}
