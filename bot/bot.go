package bot

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"toobo/meteo"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jackc/pgx/v5"
)

type Sending struct {
	City      string
	UserInfos []UserInfo
}

type UserInfo struct {
	ChatID   int64
	ChatName string
	Delete   bool
}

type Message struct {
	City    string
	UserInf UserInfo
}

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

func craftCustomMessage(userInfo UserInfo, dataMeteo meteo.MeteoResponse) string {
	var message string
	message += fmt.Sprintf("Salut %s ☀️!\nAlors, quel est la température aujourd'hui?\n", userInfo.ChatName)
	message += meteo.GetAdvices(dataMeteo)
	message += "Bonne journée!"
	return message
}

func sendMessage(bot *tgbotapi.BotAPI, message string, chatId int64) {
	msg := tgbotapi.NewMessage(chatId, message)
	if _, err := bot.Send(msg); err != nil {
		log.Printf("Failed to send to %d: %v", chatId, err)
	}
}

// TODO : add database upgrade in this function instead of just append the 2 lists
func getChatIds(bot *tgbotapi.BotAPI) []Message {
	updateChatIdsToDBFromBot(bot)
	chatIdsDB := getChatIdsFromDB()
	return chatIdsDB
}

func testAddAndDelete(conn *pgx.Conn) {
	testChatId := int64(999999)
	testCity := "TestCity"
	testName := "TestName"

	addRows(conn, testChatId, testCity, testName)

	deleteRow(conn, testChatId)
}

func updateChatIdsToDBFromBot(bot *tgbotapi.BotAPI) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_PULLER"))
	if err != nil {
		log.Fatal(err)
	}
	// Just to trigger the Supabase table at least once a day
	testAddAndDelete(conn)
	updates, err := bot.GetUpdates(tgbotapi.NewUpdate(0))
	if err != nil {
		log.Fatal(err)
	}

	seen := make(map[int64]bool)
	for i := len(updates) - 1; i >= 0; i-- {
		u := updates[i]
		id := u.Message.Chat.ID

		if !seen[id] {
			seen[id] = true
			mots := strings.Fields(u.Message.Text)
			if len(mots) > 1 {
				if mots[1] == "Supprimer" {
					deleteRow(conn, u.Message.Chat.ID)
				} else {
					//Remove Leading Slash for groups that needs it to contact the bot
					checkToAddRows(conn, u.Message.Chat.ID, mots[1], removeLeadingSlash(mots[0]))
				}
			} else {
				// No user or city can be determined, use default
				checkToAddRows(conn, u.Message.Chat.ID, "Paris", "les amis")
			}
		}
	}
}

func removeLeadingSlash(s string) string {
	return strings.TrimPrefix(s, "/")
}

func getChatIdsFromDB() []Message {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_PULLER"))
	rows, err := conn.Query(context.Background(), "SELECT id, chat_id, city, chat_name FROM chat_info_telegram_toobo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var chatIdsDB []Message
	for rows.Next() {
		var id, chat_id int64
		var city, chat_name string
		if err := rows.Scan(&id, &chat_id, &city, &chat_name); err != nil {
			log.Fatal(err)
		}
		ui := UserInfo{
			ChatID:   chat_id,
			ChatName: chat_name,
			Delete:   false,
		}
		m := Message{
			City:    city,
			UserInf: ui,
		}
		chatIdsDB = append(chatIdsDB, m)
	}
	return chatIdsDB
}

func deleteRow(conn *pgx.Conn, chatId int64) {
	_, err := conn.Exec(context.Background(),
		"DELETE FROM chat_info_telegram_toobo WHERE chat_id = $1",
		chatId,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func checkToAddRows(conn *pgx.Conn, chatId int64, city string, name string) {
	var existingId int64
	err := conn.QueryRow(context.Background(),
		"SELECT chat_id FROM chat_info_telegram_toobo WHERE chat_id = $1",
		chatId,
	).Scan(&existingId)

	if err == pgx.ErrNoRows {
		addRows(conn, chatId, city, name)
	} else if err != nil {
		log.Fatal(err)
	} else {
		_, err = conn.Exec(context.Background(),
			"UPDATE chat_info_telegram_toobo SET city = $1, chat_name = $2 WHERE chat_id = $3",
			city, name, chatId,
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func addRows(conn *pgx.Conn, chatId int64, city string, name string) {
	_, err := conn.Exec(context.Background(),
		"INSERT INTO chat_info_telegram_toobo (chat_id, city, chat_name) VALUES ($1, $2, $3)",
		chatId, city, name,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func connectDatabase() *pgx.Conn {
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
	}
	return conn
}

func getSendings(chatIds []Message) []Sending {
	cityMap := make(map[string][]UserInfo)
	cityOrder := []string{}

	for _, chatId := range chatIds {
		if _, exists := cityMap[chatId.City]; !exists {
			cityOrder = append(cityOrder, chatId.City)
		}
		cityMap[chatId.City] = append(cityMap[chatId.City], chatId.UserInf)
	}

	sendings := make([]Sending, 0, len(cityMap))
	for _, city := range cityOrder {
		sendings = append(sendings, Sending{
			City:      city,
			UserInfos: cityMap[city],
		})
	}
	return sendings
}

func sendMessages(bot *tgbotapi.BotAPI, userInfos []UserInfo, dataMeteo meteo.MeteoResponse) {
	for _, userInfo := range userInfos {
		//For each user
		messageCustom := craftCustomMessage(userInfo, dataMeteo)
		sendMessage(bot, messageCustom, userInfo.ChatID)
	}
}

func CallBotTelegram() {
	bot := getBot()
	chatIds := getChatIds(bot)
	sendings := getSendings(chatIds)
	//For each city
	for _, sending := range sendings {
		dataMeteo := meteo.CallMeteo(sending.City)
		sendMessages(bot, sending.UserInfos, dataMeteo)
	}
}
