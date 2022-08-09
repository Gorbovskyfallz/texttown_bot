package main

import (
	"fmt"

	"kek/clients/telegram"
	"kek/non_specific"
)

func main() {

	//https://api.telegram.org/bot<token>/METHOD_NAME.

	telegramApiURL := "https://api.telegram.org/bot"
	botToken := non_specific.MustToken()
	botURL := telegramApiURL + botToken + "/"
	//offset := 0

	//for {
	//	updates, updateErr := getUpdates(botURL, offset)
	//	if updateErr != nil {
	//		log.Println("error in superloop with getting updates", updateErr)
	//	}
	//	for _, update := range updates {
	//		err := respond(botURL, update)
	//		if err != nil {
	//			log.Println("kek")
	//		}
	//		offset = update.UpdateId + 1
	//		fmt.Println(update.UpdateId)
	//	}
	//	//fmt.Println(updates)
	//}

	res, _ := telegram.GetUpdates(botURL)
	fmt.Println(res)

}

// getting updates from telegram api https://core.telegram.org/bots/api#getting-updates
//func getUpdates(botURL string) ([]telegram.Update, error) {
//
//	resp, getErr := http.Get(botURL + "getUpdates")
//	if getErr != nil {
//		return nil, fmt.Errorf("cant reach updates from server: %w", getErr)
//	}
//	defer resp.Body.Close()
//	body, bodyReadErr := io.ReadAll(resp.Body)
//	if bodyReadErr != nil {
//		return nil, fmt.Errorf("cant read response body: %w", bodyReadErr)
//	}
//	var restResponse telegram.RestResponse
//	unmarshErr := json.Unmarshal(body, &restResponse)
//	if unmarshErr != nil {
//		return nil, fmt.Errorf("cant unmarshal json.body: %w", unmarshErr)
//	}
//	return restResponse.Result, nil
//}

//func respond(botURL string, update telegram.Update) error {
//	var botMessage telegram.BotMessage
//	botMessage.ChatID = update.Message.Chat.ChatId
//	botMessage.Text = update.Message.Text
//
//	buf, marshErr := json.Marshal(botMessage)
//	if marshErr != nil {
//		return fmt.Errorf("problems with marshaling botmessage: %w", marshErr)
//	}
//	_, errPOST := http.Post(botURL+"sendMessage", "application/json", bytes.NewBuffer(buf))
//	//какойто хитрый ридер, надо почитать
//	if errPOST != nil {
//		return fmt.Errorf("something get worng with post messages to server: %w", errPOST)
//	}
//	return nil
//}
