package bot

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func connectDB() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DB_PULLER"))
	if err != nil {
		log.Fatal(err)
	}
	return conn
}

// healthCheckDB insère et supprime une ligne de test pour maintenir
// la connexion Supabase active au moins une fois par jour
func healthCheckDB(conn *pgx.Conn) {
	const testChatID = int64(999999)
	insertUser(conn, testChatID, "TestCity", "TestName")
	removeUser(conn, testChatID)
}

func fetchUsersFromDB() []Message {
	conn := connectDB()
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(),
		"SELECT id, chat_id, city, chat_name FROM chat_info_telegram_toobo",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var messages []Message
	for rows.Next() {
		var id, chatID int64
		var city, name string
		if err := rows.Scan(&id, &chatID, &city, &name); err != nil {
			log.Fatal(err)
		}
		messages = append(messages, Message{
			City: city,
			User: User{ChatID: chatID, Name: name},
		})
	}
	return messages
}

func upsertUser(conn *pgx.Conn, chatID int64, city string, name string) {
	var existingID int64
	err := conn.QueryRow(context.Background(),
		"SELECT chat_id FROM chat_info_telegram_toobo WHERE chat_id = $1",
		chatID,
	).Scan(&existingID)

	if err == pgx.ErrNoRows {
		insertUser(conn, chatID, city, name)
	} else if err != nil {
		log.Fatal(err)
	} else {
		_, err = conn.Exec(context.Background(),
			"UPDATE chat_info_telegram_toobo SET city = $1, chat_name = $2 WHERE chat_id = $3",
			city, name, chatID,
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func insertUser(conn *pgx.Conn, chatID int64, city string, name string) {
	_, err := conn.Exec(context.Background(),
		"INSERT INTO chat_info_telegram_toobo (chat_id, city, chat_name) VALUES ($1, $2, $3)",
		chatID, city, name,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func removeUser(conn *pgx.Conn, chatID int64) {
	_, err := conn.Exec(context.Background(),
		"DELETE FROM chat_info_telegram_toobo WHERE chat_id = $1",
		chatID,
	)
	if err != nil {
		log.Fatal(err)
	}
}
