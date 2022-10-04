/*
 * @Author: liangdong09
 * @Date: 2022-07-23 20:41:00
 * @LastEditTime: 2022-10-04 13:05:17
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/config/autoload/weChat.go
 */
package autoload

type WeChatConfig struct {
	Rotbot         []Robot `ini:"wechat" yaml:"robot"`
	CompanyId      string  `ini:"wechat" yaml:"company_id"`
	Token          string  `ini:"wechat" yaml:"token"`
	EncodingAesKey string  `ini:"wechat" yaml:"encoding_aes_key"`
}

type Robot struct {
	Name    string `ini:"wechat" yaml:"name"`
	AgentId string `ini:"wechat" yaml:"agent_id"`
	Secret  string `ini:"wechat" yaml:"secret"`
}

var WeChat = WeChatConfig{}
