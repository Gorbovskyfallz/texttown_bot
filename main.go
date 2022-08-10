package main

import (
	"fmt"
	"log"

	"kek/clients/telegram"
	"kek/non_specific"
)

func main() {

	//https://api.telegram.org/bot<token>/METHOD_NAME.

	telegramApiURL := "https://api.telegram.org/bot"
	botToken := non_specific.MustBotToken()
	botURL := telegramApiURL + botToken + "/"
	offset := 0

	for {
		updates, updateErr := telegram.GetUpdates(botURL, offset)
		if updateErr != nil {
			log.Println("error in superloop with getting updates", updateErr)
		}
		for _, update := range updates {
			err := telegram.Respond(botURL, update)
			if err != nil {
				log.Println("kek")
			}
			offset = update.UpdateId + 1
			fmt.Println(update.UpdateId)
		}
		//fmt.Println(updates)
	}

}
