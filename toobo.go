package main

import (
	"toobo/bot"
	"toobo/meteo"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	dataMeteo := meteo.CallMeteo()
	bot.CallBotTelegram(dataMeteo)
}
