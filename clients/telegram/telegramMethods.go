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

	botMessage.Text = commandCases(botURL, update.Message.Text)

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

func GetMeResponse(botURL string) (GetME, error) {

	var getMeresponse GetME

	resp, getErr := http.Get(botURL + "getME")
	if getErr != nil {
		return getMeresponse, fmt.Errorf("cant reach updates from server: %w", getErr)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, bodyReadErr := io.ReadAll(resp.Body)
	if bodyReadErr != nil {
		return getMeresponse, fmt.Errorf("cant read response body: %w", bodyReadErr)
	}

	unmarshErr := json.Unmarshal(body, &getMeresponse)
	if unmarshErr != nil {
		return getMeresponse, fmt.Errorf("cant unmarshal json.body: %w", unmarshErr)
	}
	return getMeresponse, nil

}

func commandCases(botURL string, updateMessage string) (postString string) {

	switch updateMessage {

	case "/getme":
		{
			getMeResponse, _ := GetMeResponse(botURL)
			info := getMeResponse.Result.FirstName + "\n" + getMeResponse.Result.UserName

			postString = info

		}
	default:
		postString = "no such command, use /help to see bot commands"

	}

	return postString

}
