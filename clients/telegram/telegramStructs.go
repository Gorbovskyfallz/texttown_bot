package telegram

type Update struct {
	UpdateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}
type Chat struct {
	ChatId int `json:"id"`
}

type RestResponse struct {
	Result []Update `json:"result"`
}

type BotMessage struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}

type GetME struct {
	Ok     bool   `json:"ok"`
	Result Result `json:"result"`
}

type Result struct {
	FirstName string `json:"first_name"`
	UserName  string `json:"username"`
}
