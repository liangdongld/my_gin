/*
 * @Author: liangdong09
 * @Date: 2022-07-23 20:36:14
 * @LastEditTime: 2022-08-06 01:02:09
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/internal/model/send_msg.go
 */
package model

type WcSendMsg struct {
	ToUser  string `json:"touser"`
	MsgType string `json:"msgtype"`
	AgentId string `json:"agentid"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

type MsgContent struct {
	ToUsername   string  `xml:"ToUserName" json:"ToUserName"`
	FromUsername string  `xml:"FromUserName" json:"FromUserName"`
	CreateTime   uint32  `xml:"CreateTime" json:"CreateTime"`
	MsgType      string  `xml:"MsgType" json:"MsgType"`
	Content      string  `xml:"Content" json:"Content"`
	PicUrl       string  `xml:"PicUrl" json:"PicUrl"`
	MediaId      string  `xml:"MediaId" json:"MediaId"`
	Msgid        string  `xml:"MsgId" json:"Msgid"`
	Format       string  `xml:"Format" json:"Format"`
	ThumbMediaId string  `xml:"ThumbMediaId" json:"ThumbMediaId"`
	Location_X   string  `xml:"Location_X" json:"Location_X"`
	Location_Y   string  `xml:"Location_Y" json:"Location_Y"`
	Scale        string  `xml:"Scale" json:"Scale"`
	Label        string  `xml:"Label" json:"Label"`
	Agentid      uint32  `xml:"AgentId" json:"Agentid"`
	Latitude     float32 `xml:"Latitude" json:"Latitude"`
	Longitude    float32 `xml:"Longitude" json:"Longitude"`
	Precision    int     `xml:"Precision" json:"Precision"`
}

func (t *WcSendMsg) SetAgentId(agentId string) {
	t.AgentId = agentId
}

func (t *WcSendMsg) SetMessage(message string) {
	t.Text.Content = message
}
