/*
 * @Author: liangdong09
 * @Date: 2022-08-05 19:31:20
 * @LastEditTime: 2022-08-06 00:57:30
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/internal/service/receive/recevie.go
 */
package receive

import "github.com/liangdong/my-gin/internal/model"

type Receive interface {
	ReplyMsg() (model.MsgContent, error)
}

func ReplyMsg(receive Receive) (model.MsgContent, error) {
	msg, err := receive.ReplyMsg()
	if err != nil {
		return msg, err
	}
	return msg, err
}

/**
 * @description: 针对不同类型消息进行处理
 * @param {model.MsgContent} msg
 * @return {*}
 */
func ReceiveMsg(msg model.MsgContent) (model.MsgContent, error) {
	switch msg.MsgType {
	case "event":
		return ReplyMsg(&ReceiveLocation{Msg: msg})
	case "txt":
		return ReplyMsg(&ReceiveTxt{Msg: msg})
	default:
		return ReplyMsg(&ReceiveTxt{Msg: msg})
	}
}
