package main

import (
	"toobo/bot"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	bot.Run()
}
