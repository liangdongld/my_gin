/*
 * @Author: liangdong09
 * @Date: 2022-07-23 20:36:14
 * @LastEditTime: 2022-11-06 22:13:20
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/internal/model/send_msg.go
 */
package model

import "gorm.io/gorm"

type WcSendMsg struct {
	ToUser  string `json:"touser"`
	MsgType string `json:"msgtype"`
	AgentId string `json:"agentid"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

type MsgContent struct {
	gorm.Model
	ToUsername   string  `xml:"ToUserName" json:"ToUserName"`
	FromUsername string  `xml:"FromUserName" json:"FromUserName"`
	CreateTime   uint32  `xml:"CreateTime" json:"CreateTime"`
	MsgType      string  `xml:"MsgType" json:"MsgType"`
	Content      string  `xml:"Content" json:"Content"`
	PicUrl       string  `xml:"PicUrl" json:"PicUrl"`
	MediaId      string  `xml:"MediaId" json:"MediaId"`
	MsgId        string  `xml:"MsgId" json:"MsgId"`
	Format       string  `xml:"Format" json:"Format"` //多媒体数据格式
	ThumbMediaId string  `xml:"ThumbMediaId" json:"ThumbMediaId"`
	Location_X   string  `xml:"Location_X" json:"Location_X"`
	Location_Y   string  `xml:"Location_Y" json:"Location_Y"`
	Scale        string  `xml:"Scale" json:"Scale"`
	Label        string  `xml:"Label" json:"Label"`
	Agentid      uint32  `xml:"AgentId" json:"Agentid"`
	Latitude     float64 `xml:"Latitude" json:"Latitude"`
	Longitude    float64 `xml:"Longitude" json:"Longitude"`
	Precision    int     `xml:"Precision" json:"Precision"`
	MarkDown     struct {
		Content string `xml:"Content" json:"Content"`
	} `xml:"MarkDown" json:"MarkDown"`
}

func (t *WcSendMsg) SetAgentId(agentId string) {
	t.AgentId = agentId
}

func (t *WcSendMsg) SetMessage(message string) {
	t.Text.Content = message
}
