package bot

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"toobo/meteo"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func getBot() *tgbotapi.BotAPI {
	token := os.Getenv("TELEGRAM_TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}
	return bot
}

func craftMessage(dataMeteo meteo.MeteoResponse) string {
	message := "Salut les copains ☀️!\nAlors, quel est la température aujourd'hui?\n"
	message += fmt.Sprintf("Température: %.1f\n", dataMeteo.Daily.TemperatureMax[0])
	message += "Avec cette température je te conseil je t'habiller comme ca!\n"
	message += meteo.GetAdviceTemperature(dataMeteo)
	message += "Bonne journée!"
	// message := fmt.Sprintf("Température: %.1f°C\nVent: %.1f km/h\nWeather code: %d", dataMeteo.Daily.TemperatureMax[0], dataMeteo.Daily.WindSpeed[0], dataMeteo.Daily.WeatherCode[0])
	return message
}

func sendMessage(bot *tgbotapi.BotAPI, message string) {
	// ChatID cible (utilisateur, groupe ou canal)
	chatIDStr := os.Getenv("CHAT_ID")
	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	msg := tgbotapi.NewMessage(chatID, message)
	_, err = bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}
}
func CallBotTelegram(dataMeteo meteo.MeteoResponse) {
	bot := getBot()
	message := craftMessage(dataMeteo)
	sendMessage(bot, message)
}
