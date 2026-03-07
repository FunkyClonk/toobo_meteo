package bot

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"toobo/meteo"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5"
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
	message += meteo.GetAdvices(dataMeteo)
	message += "Bonne journée!"
	// message := fmt.Sprintf("Température: %.1f°C\nVent: %.1f km/h\nWeather code: %d", dataMeteo.Daily.TemperatureMax[0], dataMeteo.Daily.WindSpeed[0], dataMeteo.Daily.WeatherCode[0])
	return message
}

func sendMessage(bot *tgbotapi.BotAPI, chatIds []int64, message string) {
	// ChatID cible (utilisateur, groupe ou canal)
	chatIDsStr := os.Getenv("CHAT_IDS")
	chatIDParts := strings.Split(chatIDsStr, ",")

	for _, part := range chatIDParts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		chatID, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			log.Printf("Invalid CHAT_ID '%s': %v", part, err)
			continue
		}

		msg := tgbotapi.NewMessage(chatID, message)
		if _, err := bot.Send(msg); err != nil {
			log.Printf("Failed to send to %d: %v", chatID, err)
		}
	}
	// TODO : Update with chatIds only if not present in defaut chatIds
	// fmt.Println("Deuxieme envoie")
	// for _, chatId := range chatIds {
	// 	msg := tgbotapi.NewMessage(chatId, "message")
	// 	if _, err := bot.Send(msg); err != nil {
	// 		log.Printf("Failed to send to %d: %v", chatId, err)
	// 	}
	// }
}

func getChatIds(bot *tgbotapi.BotAPI) []int64 {
	updates, err := bot.GetUpdates(tgbotapi.NewUpdate(0))
	if err != nil {
		log.Fatal(err)
	}

	seen := make(map[int64]bool)
	var ids []int64
	for _, u := range updates {
		id := u.Message.Chat.ID
		if !seen[id] {
			seen[id] = true
			ids = append(ids, id)
		}
	}
	return ids
}

func getRows(conn *pgx.Conn) {
	rows, err := conn.Query(context.Background(), "SELECT id, chat_id, city, chat_name FROM chat_info_telegram_toobo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id, chat_id int
		var city, chat_name string
		if err := rows.Scan(&id, &chat_id, &city, &chat_name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("id=%d chat_id=%d city=%s user=%s\n", id, chat_id, city, chat_name)
	}
}

func addRows(conn *pgx.Conn) {
	_, err := conn.Exec(context.Background(),
		"INSERT INTO chat_info_telegram_toobo (chat_id, city, chat_name) VALUES ($1, $2, $3)",
		"654", "Paris", "Le parisien",
	)
	if err != nil {
		log.Fatal(err)
	}
}
func pingDatabase() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_PULLER"))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer conn.Close(context.Background())

	// Example query to test connection
	var version string
	if err := conn.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
		log.Fatalf("Query failed: %v", err)
	}
	log.Println("Connected to:", version)

	rows, err := conn.Query(context.Background(), `
		SELECT table_name 
		FROM information_schema.tables 
		WHERE table_schema = 'public'
	`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		rows.Scan(&name)
		fmt.Println(name)
	}

	// getRows(conn)
	// addRows(conn)

}

func CallBotTelegram(dataMeteo meteo.MeteoResponse) {
	bot := getBot()
	pingDatabase()
	message := craftMessage(dataMeteo)
	chatIds := getChatIds(bot)
	sendMessage(bot, chatIds, message)
}

//TODO
// func SaveChatID(id int64) {
//     f, _ := os.OpenFile("chat_ids.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//     defer f.Close()
//     fmt.Fprintln(f, id)
// }
