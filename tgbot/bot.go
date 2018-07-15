package tgbot

const (
	baseURL = "https://api.telegram.org/bot"
)

type bot struct {
	url string
}

func newBot(c *config) *bot {
	b := &bot{
		url: baseURL + c.Token,
	}
	return b
}

func (b *bot) getChatID() {

}
