package bot

import (
	"log"
	"os"
	"toobo/meteo"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func newBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	return bot
}

func Run() {
	bot := newBot()
	messages := fetchMessages(bot)
	sendings := groupByCity(messages)
	for _, sending := range sendings {
		weather := meteo.FetchWeather(sending.City)
		sendWeatherToUsers(bot, sending.Users, weather)
	}
}
