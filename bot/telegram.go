package bot

import (
	"context"
	"fmt"
	"log"
	"strings"
	"toobo/meteo"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5"
)

func fetchMessages(bot *tgbotapi.BotAPI) []Message {
	syncUsersFromTelegram(bot)
	return fetchUsersFromDB()
}

func syncUsersFromTelegram(bot *tgbotapi.BotAPI) {
	conn := connectDB()
	defer conn.Close(context.Background())

	healthCheckDB(conn)

	updates, err := bot.GetUpdates(tgbotapi.NewUpdate(0))
	if err != nil {
		log.Fatal(err)
	}

	seen := make(map[int64]bool)
	for i := len(updates) - 1; i >= 0; i-- {
		u := updates[i]
		if u.Message == nil {
			continue
		}
		chatID := u.Message.Chat.ID
		if seen[chatID] {
			continue
		}
		seen[chatID] = true
		handleUpdate(conn, chatID, u.Message.Text)
	}
}

func handleUpdate(conn *pgx.Conn, chatID int64, text string) {
	words := strings.Fields(text)
	if len(words) == 0 {
		upsertUser(conn, chatID, "Paris", "les amis")
		return
	}

	cmd := strings.TrimPrefix(words[0], "/")
	switch {
	case strings.HasPrefix(cmd, "delete"):
		removeUser(conn, chatID)
	case cmd == "add" && len(words) >= 3:
		// /add <name> <city>
		upsertUser(conn, chatID, words[2], words[1])
	default:
		upsertUser(conn, chatID, "Paris", "les amis")
	}
}

func groupByCity(messages []Message) []Sending {
	cityMap := make(map[string][]User)
	cityOrder := []string{}

	for _, msg := range messages {
		if _, exists := cityMap[msg.City]; !exists {
			cityOrder = append(cityOrder, msg.City)
		}
		cityMap[msg.City] = append(cityMap[msg.City], msg.User)
	}

	sendings := make([]Sending, 0, len(cityMap))
	for _, city := range cityOrder {
		sendings = append(sendings, Sending{City: city, Users: cityMap[city]})
	}
	return sendings
}

func sendWeatherToUsers(bot *tgbotapi.BotAPI, users []User, weather meteo.MeteoResponse) {
	for _, user := range users {
		text, imageRef := buildWeatherMessage(user, weather)
		sendMessageWithPhoto(bot, user.ChatID, text, imageRef)
	}
}

func buildWeatherMessage(user User, weather meteo.MeteoResponse) (string, string) {
	msg := fmt.Sprintf("Salut %s ☀️!\nAlors, quel est la température aujourd'hui?\n", user.Name)
	advice, imageRef := meteo.BuildAdvice(weather)
	msg += advice + "Bonne journée!"
	return msg, imageRef
}

func sendMessageWithPhoto(bot *tgbotapi.BotAPI, chatID int64, text string, imageRef string) {
	msg := tgbotapi.NewMessage(chatID, text)
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Failed to send message to %d: %v", chatID, err)
		return
	}
	photo := tgbotapi.NewPhoto(chatID, tgbotapi.FilePath(imageRef))
	if _, err := bot.Send(photo); err != nil {
		log.Printf("Failed to send photo to %d: %v", chatID, err)
	}
}
