package autoload

type WeChatConfig struct {
	AgentId   string `ini:"wechat" yaml:"agent_id"`
	Secret    string `ini:"wechat" yaml:"secret"`
	CompanyId string `ini:"wechat" yaml:"companyId"`
}

var WeChat = WeChatConfig{}
