package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetUpdates(botURL string, offset int) ([]Update, error) {

	resp, getErr := http.Get(botURL + "getUpdates" + "?offset=" + strconv.Itoa(offset))
	if getErr != nil {
		return nil, fmt.Errorf("cant reach updates from server: %w", getErr)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
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

func Respond(botURL string, update Update) error {
	var botMessage BotMessage
	botMessage.ChatID = update.Message.Chat.ChatId

	if update.Message.Text == "/getme" {
		info, _ := GetMeResponse(botURL)
		info += "\n" + "kekeke"

		botMessage.Text = info
	} else {
		botMessage.Text = update.Message.Text
	}

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

func GetMeResponse(botURL string) (string, error) {

	resp, getErr := http.Get(botURL + "getME")
	if getErr != nil {
		return "", fmt.Errorf("cant reach updates from server: %w", getErr)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, bodyReadErr := io.ReadAll(resp.Body)
	if bodyReadErr != nil {
		return "", fmt.Errorf("cant read response body: %w", bodyReadErr)
	}
	var getMeresponse GetME
	unmarshErr := json.Unmarshal(body, &getMeresponse)
	if unmarshErr != nil {
		return "", fmt.Errorf("cant unmarshal json.body: %w", unmarshErr)
	}
	return getMeresponse.Result.FirstName, nil

}
