/*
 * @Author: liangdong09
 * @Date: 2022-10-06 14:16:54
 * @LastEditTime: 2022-10-06 14:17:47
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/pkg/calendar/period_test.go
 */
package calendar

import "testing"

func TestPredictNextPeriod(t *testing.T) {
	ret := PredictNextPeriod()
	t.Log(ret)
}
