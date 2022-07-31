/*
 * @Author: liangdong09
 * @Date: 2022-07-23 20:26:24
 * @LastEditTime: 2022-07-31 13:34:43
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/internal/service/weChat.go
 */
package service

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"

	c "github.com/wannanbigpig/gin-layout/config"
	"github.com/wannanbigpig/gin-layout/data"
	"github.com/wannanbigpig/gin-layout/internal/model"
	log "github.com/wannanbigpig/gin-layout/internal/pkg/logger"
	internal_utils "github.com/wannanbigpig/gin-layout/internal/pkg/utils"
	"github.com/wannanbigpig/gin-layout/pkg/utils"
	"go.uber.org/zap"
)

/**
 * @description: 给企微发送消息
 * @param {string} message 消息内容
 * @param {string} msgType 消息类型
 * @return {*}
 */
func SendWeChat(message string, msgType string) error {
	redis_key := "access_token"
	// 尝试从redis中读取token
	accessToken := data.GetRedis(redis_key)
	http := &utils.HttpRequest{}
	// 若redis中的token已过期，则重新请求api获取token
	if accessToken == "" {
		log.Logger.Info("access token is null, will recall")
		getTokenUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s",
			c.Config.WeChat.CompanyId, c.Config.WeChat.Secret)
		log.Logger.Info("token_url", zap.String("url", getTokenUrl))
		http.Request("GET", getTokenUrl, nil)
		ret := make(map[string]interface{})
		if err := http.ParseJson(&ret); err != nil {
			return err
		}
		marshal, _ := json.Marshal(ret)
		log.Logger.Info(string(marshal))
		accessToken = fmt.Sprintf("%v", ret["access_token"])
		// 写入redis 有效期2小时
		data.SetRedis(redis_key, accessToken, 7200)
	}
	msg := &model.WcSendMsg{
		ToUser:  "@all",
		MsgType: msgType,
		AgentId: c.Config.WeChat.AgentId,
	}
	msg.SetMessage(message)
	sendMsgUrl := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%v", accessToken)
	log.Logger.Info("sendMsgUrl = " + string(sendMsgUrl))
	header := map[string]string{"Content-Type": "application/json"}
	bytesData, _ := json.Marshal(msg)
	http.Request("POST", sendMsgUrl, bytes.NewReader(bytesData), header)
	log.Logger.Info("bytes data = " + string(bytesData))
	ret := make(map[string]interface{})
	err := http.ParseJson(&ret)
	if err != nil {
		return err
	}
	if ret["errcode"].(float64) != 0 {
		errmsg := fmt.Sprintf("%v", ret["errmsg"])
		return errors.New(errmsg)
	}
	return nil
}

// 验证
func VerifyMsg(verifyMsgSign string, verifyTimestamp string, verifyNonce string, verifyEchoStr string) (string, error) {
	token := c.Config.WeChat.Token
	EncodingAESKey := c.Config.WeChat.EncodingAesKey
	wxcpt := internal_utils.NewWXBizMsgCrypt(token, EncodingAESKey, "", internal_utils.XmlType)
	echoStr, cryptErr := wxcpt.VerifyURL(verifyMsgSign, verifyTimestamp, verifyNonce, verifyEchoStr)
	if nil != cryptErr {
		fmt.Println("verifyUrl fail", cryptErr)
		return "", errors.New(cryptErr.ErrMsg)
	}
	fmt.Println("verifyUrl success echoStr", string(echoStr))
	return string(echoStr), nil
}

func ReceiveMsg(reqMsgSign, reqTimestamp, reqNonce string, reqData []byte) (string, error) {
	token := c.Config.WeChat.Token
	EncodingAESKey := c.Config.WeChat.EncodingAesKey
	wxcpt := internal_utils.NewWXBizMsgCrypt(token, EncodingAESKey, "", internal_utils.XmlType)
	msg, cryptErr := wxcpt.DecryptMsg(reqMsgSign, reqTimestamp, reqNonce, reqData)
	if nil != cryptErr {
		fmt.Println("DecryptMsg fail", cryptErr)
	}
	var msgContent model.MsgContent
	err := xml.Unmarshal(msg, &msgContent)
	if err != nil {
		return "", errors.New("unmarshal fail")
	}
	bt, _ := json.Marshal(msgContent)
	str := utils.ByteSliceToString(bt)
	log.Logger.Sugar().Infof("received message: [%s]", str)
	// content := msgContent.Content
	// msgContent.Content = content
	bt, _ = xml.Marshal(msgContent)
	str = utils.ByteSliceToString(bt)
	encryptMsg, _ := wxcpt.EncryptMsg(str, reqTimestamp, reqNonce)
	return string(encryptMsg), nil
}
