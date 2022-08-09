package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {

	//https://api.telegram.org/bot<token>/METHOD_NAME.

	api := "https://api.telegram.org/bot"
	token := mustToken()
	finalapi := api + token + "/" + "<method name>"
	fmt.Println(finalapi)

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

func getUpdates() {

}

func respond() {

}
