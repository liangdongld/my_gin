/*
 * @Author: liangdong09
 * @Date: 2022-07-23 20:41:00
 * @LastEditTime: 2022-08-05 20:56:43
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/config/autoload/weChat.go
 */
/*
 * @Author: liangdong09
 * @Date: 2022-07-23 20:41:00
 * @LastEditTime: 2022-08-05 20:55:57
 * @LastEditors: liangdong09
 * @Description:
 * @FilePath: /my_gin/config/autoload/weChat.go
 */
package autoload

type WeChatConfig struct {
	AgentId        string `ini:"wechat" yaml:"agent_id"`
	Secret         string `ini:"wechat" yaml:"secret"`
	CompanyId      string `ini:"wechat" yaml:"company_id"`
	Token          string `ini:"wechat" yaml:"token"`
	EncodingAesKey string `ini:"wechat" yaml:"encoding_aes_key"`
}

var WeChat = WeChatConfig{}
