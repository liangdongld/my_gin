/*
 * @Author: liangdong09
 * @Date: 2022-08-05 19:31:20
 * @LastEditTime: 2022-08-05 19:53:46
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/internal/service/receive/recevie.go
 */
package receive

import "github.com/liangdong/my-gin/internal/model"

type Receive interface {
	ReplyMsg() (model.MsgContent, error)
}

func ReceiveMsg(msg model.MsgContent) (model.MsgContent, error) {
	switch msg.MsgType {
	case "event":
		m := &ReceiveLocation{Msg: msg}
		return m.ReplyMsg()
	case "txt":
		m := &ReceiveTxt{Msg: msg}
		return m.ReplyMsg()
	default:
		m := &ReceiveTxt{Msg: msg}
		return m.ReplyMsg()
	}
}
