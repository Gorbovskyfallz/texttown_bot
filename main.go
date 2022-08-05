package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {

	fmt.Println(mustToken())

	//fetcher = fetcher.New()
	//processor - processor.New()
	// tgClient = telegram.New(token) // нельзя создавать токен в виде константы в коде, так как через открытые
	// getToken = flags.Get(token)
	// это дело может уйти в другие руки
	// consumer.Start (fetcher, processor)

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
