package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"kek/clients/telegram"
	"log"
	"net/http"
)

func main() {

	//https://api.telegram.org/bot<token>/METHOD_NAME.

	api := "https://api.telegram.org/bot"
	token := mustToken()
	finalApi := api + token + "/"

	for {
		updates, updateErr := getUpdates(finalApi)
		if updateErr != nil {
			log.Println("error in superloop with getting updates", updateErr)
		}
		fmt.Println(updates)
	}

}

// функция для считывания токена с помощью флагов при запуске программы
func mustToken() string {
	// приставка must функции говорит о том, что функция при неверной работе не возвращает
	//ошибку, а натурально падает, сообщая нам о проблеме
	token := flag.String("t", "", "token to access telegram bot")
	flag.Parse()
	if *token == "" {
		log.Fatal("invalid or not defined token for telegram bot")
	}
	return *token

}

func getUpdates(botURL string) ([]telegram.Update, error) {

	resp, getErr := http.Get(botURL + "getUpdates")
	if getErr != nil {
		return nil, fmt.Errorf("cant reach updates from server: %w", getErr)
	}
	defer resp.Body.Close()
	body, bodyReadErr := io.ReadAll(resp.Body)
	if bodyReadErr != nil {
		return nil, fmt.Errorf("cant read response body: %w", bodyReadErr)
	}
	var restResponse telegram.RestResponse
	unmarshErr := json.Unmarshal(body, &restResponse)
	if unmarshErr != nil {
		return nil, fmt.Errorf("cant unmarshal json.body: %w", unmarshErr)
	}
	return restResponse.Result, nil
}

func respond() {

}
