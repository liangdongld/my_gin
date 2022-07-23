package model

type wcSendcontent struct {
	Content string `json:"content"`
}

type WcSendMsg struct {
	ToUser  string        `json:"touser"`
	MsgType string        `json:"msgtype"`
	AgentId string        `json:"agentid"`
	Text    wcSendcontent `json:"text"`
}

func (t *WcSendMsg) SetAgentId(agentId string) {
	t.AgentId = agentId
}

func (t *WcSendMsg) SetMessage(message string) {
	t.Text.Content = message
}
