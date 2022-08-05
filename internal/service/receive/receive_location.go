package receive

import (
	"encoding/json"
	"fmt"

	"github.com/liangdong/my-gin/config"
	"github.com/liangdong/my-gin/data"
	"github.com/liangdong/my-gin/internal/model"
	"github.com/liangdong/my-gin/pkg/utils"
)

type ReceiveLocation struct {
	Msg model.MsgContent
}

func (r *ReceiveLocation) ReplyMsg() (model.MsgContent, error) {
	if loc := data.GetRedis("location"); loc == "" {
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
		addr_byte, _ := json.Marshal(ret)
		data.SetRedis("location", utils.ByteSliceToString(addr_byte), 600)
	}
	return r.Msg, nil
}
