package telegram

import (
	"net/http"
	"net/url"
	"strconv"
)

type Client struct {
	host     string // хост апи сервиса телеграма
	basePath string // базовый путь - префикс, с которого идут все запросы
	client   http.Client
}

//tg-bot.com/bot<token> -- хост с базовым путем

func New(host, token string) Client {
	return Client{
		host:     host,
		basePath: newBasePath(token),
		client:   http.Client{},
	}
}

/*мы делаем отдельную функцию для формирования пути для того, чтобы не бегать
по коду и не менять все это дело руками в случае, если телега решит поменять
вид основного пути*/

func newBasePath(token string) (path string) {

	return "bot" + token

}

func (c *Client) Updates(offset int, limit int) ([]Update, error) {
	q := url.Values{}
	q.Add("offset", strconv.Itoa(offset))
	q.Add("limit", strconv.Itoa(limit))

}

func (c *Client) SendMessage() {

}
