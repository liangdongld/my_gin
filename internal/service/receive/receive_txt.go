/*
 * @Author: liangdong09
 * @Date: 2022-08-05 19:41:59
 * @LastEditTime: 2022-09-14 15:32:27
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/internal/service/receive/receive_txt.go
 */
package receive

import (
	"strings"

	"github.com/liangdong/my-gin/internal/model"
)

type ReceiveTxt struct {
	Msg model.MsgContent
}

func (r *ReceiveTxt) ReplyMsg() (model.MsgContent, error) {
	if strings.Contains(r.Msg.Content, "lsp") ||
		strings.Contains(r.Msg.Content, "罗圣鹏") {
		r.Msg.Content = "罗圣鹏是你儿子"
	}
	return r.Msg, nil
}
