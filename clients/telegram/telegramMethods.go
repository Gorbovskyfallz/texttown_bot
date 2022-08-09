package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetUpdates(botURL string) ([]Update, error) {

	resp, getErr := http.Get(botURL + "getUpdates")
	if getErr != nil {
		return nil, fmt.Errorf("cant reach updates from server: %w", getErr)
	}
	defer resp.Body.Close()
	body, bodyReadErr := io.ReadAll(resp.Body)
	if bodyReadErr != nil {
		return nil, fmt.Errorf("cant read response body: %w", bodyReadErr)
	}
	var restResponse RestResponse
	unmarshErr := json.Unmarshal(body, &restResponse)
	if unmarshErr != nil {
		return nil, fmt.Errorf("cant unmarshal json.body: %w", unmarshErr)
	}
	return restResponse.Result, nil
}

func respond(botURL string, update Update) error {
	var botMessage BotMessage
	botMessage.ChatID = update.Message.Chat.ChatId
	botMessage.Text = update.Message.Text

	buf, marshErr := json.Marshal(botMessage)
	if marshErr != nil {
		return fmt.Errorf("problems with marshaling botmessage: %w", marshErr)
	}
	_, errPOST := http.Post(botURL+"sendMessage", "application/json", bytes.NewBuffer(buf))
	//какойто хитрый ридер, надо почитать
	if errPOST != nil {
		return fmt.Errorf("something get worng with post messages to server: %w", errPOST)
	}
	return nil
}
