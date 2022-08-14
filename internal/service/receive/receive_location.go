/*
 * @Author: liangdong09
 * @Date: 2022-08-05 19:44:40
 * @LastEditTime: 2022-08-14 23:36:38
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/internal/service/receive/receive_location.go
 */
package receive

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/liangdong/my-gin/config"
	"github.com/liangdong/my-gin/data"
	"github.com/liangdong/my-gin/internal/model"
	"github.com/liangdong/my-gin/pkg/utils"
)

type ReceiveLocation struct {
	Msg model.MsgContent
	mux sync.RWMutex
}

func (r *ReceiveLocation) ReplyMsg() (model.MsgContent, error) {
	// double check的写法
	r.mux.RLock()
	loc := data.GetRedis("location")
	r.mux.RUnlock()
	if loc != "" {
		return r.Msg, nil
	}
	r.mux.Lock()
	defer r.mux.Unlock()
	loc = data.GetRedis("location")
	if loc != "" {
		return r.Msg, nil
	}
	r.Msg.MsgType = "text"
	hp := &utils.HttpRequest{}
	url := fmt.Sprintf("https://restapi.amap.com/v3/geocode/regeo?key=%s&location=%f,%f",
		config.Config.Amap.Key, r.Msg.Longitude, r.Msg.Latitude)
	hp.Request("GET", url, nil)
	ret := make(map[string]interface{})
	hp.ParseJson(&ret)
	regeocode := ret["regeocode"].(map[string]interface{})
	addr := regeocode["formatted_address"]
	r.Msg.Content = fmt.Sprintf("当前所在地址为:%s", addr)
	addr_byte, err := json.Marshal(ret)
	if err != nil {
		return r.Msg, err
	}
	data.SetRedis("location", utils.ByteSliceToString(addr_byte), 600)
	return r.Msg, nil
}
