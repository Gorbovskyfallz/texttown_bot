package telegram

// impl method getME

type Update struct {
	updateId int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}
type Chat struct {
	ChatId int `json:"id"`
}
