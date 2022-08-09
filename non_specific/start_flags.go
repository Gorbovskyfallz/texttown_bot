package non_specific

import (
	"flag"
	"log"
)

// getting bot token from programm start flag not to defer bot token in code
func MustToken() string {
	// приставка must функции говорит о том, что функция при неверной работе не возвращает
	//ошибку, а натурально падает, сообщая нам о проблеме
	token := flag.String("t", "", "token to access telegram bot")
	flag.Parse()
	if *token == "" {
		log.Fatal("invalid or not defined token for telegram bot")
	}
	return *token

}
