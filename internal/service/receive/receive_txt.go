/*
 * @Author: liangdong09
 * @Date: 2022-08-05 19:41:59
 * @LastEditTime: 2022-08-05 19:44:47
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/internal/service/receive/receive_txt.go
 */
package receive

import "github.com/liangdong/my-gin/internal/model"

type ReceiveTxt struct {
	Msg model.MsgContent
}

func (r *ReceiveTxt) ReplyMsg() (model.MsgContent, error) {
	return r.Msg, nil
}
