package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	c "github.com/wannanbigpig/gin-layout/config"
	"github.com/wannanbigpig/gin-layout/data"
	"github.com/wannanbigpig/gin-layout/internal/model"
	log "github.com/wannanbigpig/gin-layout/internal/pkg/logger"
	"github.com/wannanbigpig/gin-layout/pkg/utils"
	"go.uber.org/zap"
)

func SendWeChat(message string, msgType string) error {
	redis_key := "access_token"
	accessToken := data.GetRedis(redis_key)
	http := &utils.HttpRequest{}
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
