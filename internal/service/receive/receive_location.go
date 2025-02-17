/*
 * @Author: liangdong09
 * @Date: 2022-08-05 19:44:40
 * @LastEditTime: 2022-10-06 14:50:58
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
	"github.com/liangdong/my-gin/internal/pkg/logger"
	"github.com/liangdong/my-gin/pkg/utils"
)

var mux sync.RWMutex

type ReceiveLocation struct {
	Msg model.MsgContent

	loc map[string]interface{}
}

type WeatherInfo struct {
	Status   string `json:"status"`
	Count    string `json:"count"`
	Info     string `json:"info"`
	InfoCode string `json:"infocode"`
	Lives    []struct {
		Province      string `json:"province"`
		City          string `json:"city"`
		Adcode        string `json:"adcode"`
		Weather       string `json:"weather"`
		Temperature   string `json:"temperature"`
		Winddirection string `json:"winddirection"`
		Windpower     string `json:"windpower"`
		Humidity      string `json:"humidity"`
		Reporttime    string `json:"reporttime"`
	} `json:"lives"`
}

func (r *ReceiveLocation) ReplyMsg() (model.MsgContent, error) {
	// double check的写法
	mux.RLock()
	key := "location_" + r.Msg.FromUsername
	loc := data.GetRedis(key)
	mux.RUnlock()
	if loc != "" {
		return r.Msg, nil
	}
	mux.Lock()
	defer mux.Unlock()
	loc = data.GetRedis(key)
	if loc != "" {
		return r.Msg, nil
	}
	// 通过API获取当前位置信息
	err := r.getLocationByCode(&r.loc)
	if err != nil {
		return r.Msg, err
	}
	// 获取结果实际地址
	addr := r.getRealAddr()
	// 获取天气信息
	wInfo, err := r.getWeatherMsg()
	if err != nil {
		return r.Msg, err
	}

	r.Msg.MsgType = "text"
	r.Msg.Content = fmt.Sprintf("你好 %s\n \n", r.Msg.FromUsername)
	r.Msg.Content += fmt.Sprintf("当前所在地址为:\n %s\n", addr)
	r.Msg.Content += "=================\n"
	r.Msg.Content += wInfo

	addr_byte, err := json.Marshal(r.loc)
	if err != nil {
		return r.Msg, err
	}
	data.SetRedis(key, utils.ByteSliceToString(addr_byte), 86400)
	return r.Msg, nil
}

func (r *ReceiveLocation) getLocationByCode(loc *map[string]interface{}) error {
	hp := &utils.HttpRequest{}
	url := fmt.Sprintf("https://restapi.amap.com/v3/geocode/regeo?key=%s&location=%f,%f",
		config.Config.Amap.Key, r.Msg.Longitude, r.Msg.Latitude)
	hp.Request("GET", url, nil)
	err := hp.ParseJson(&loc)
	if err != nil {
		return err
	}
	return nil
}

func (r *ReceiveLocation) getRealAddr() string {
	regeocode := r.loc["regeocode"].(map[string]interface{})
	addr := regeocode["formatted_address"].(string)
	logger.Logger.Sugar().Infof("当前位置为: %s", addr)
	return addr
}

func (r *ReceiveLocation) getCityCode() string {
	cityCode := r.loc["regeocode"].(map[string]interface{})["addressComponent"].(map[string]interface{})["adcode"].(string)
	return cityCode
}

func (r *ReceiveLocation) getWeatherMsg() (string, error) {
	cityCode := r.getCityCode()
	hp := &utils.HttpRequest{}
	url := fmt.Sprintf("https://restapi.amap.com/v3/weather/weatherInfo?key=%s&city=%s",
		config.Config.Amap.Key, cityCode)
	hp.Request("GET", url, nil)
	var wMsg WeatherInfo
	err := hp.ParseJson(&wMsg)
	if err != nil {
		return "", err
	}
	msg := fmt.Sprintf("地点:%s%s\n", wMsg.Lives[0].Province, wMsg.Lives[0].City)
	msg += fmt.Sprintf("天气:%s\n", wMsg.Lives[0].Weather)
	msg += fmt.Sprintf("气温:%s\n", wMsg.Lives[0].Temperature)
	msg += fmt.Sprintf("湿度:%s", wMsg.Lives[0].Humidity)
	return msg, nil
}
