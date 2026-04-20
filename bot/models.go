package bot

type Sending struct {
	City  string
	Users []User
}

type User struct {
	ChatID int64
	Name   string
}

type Message struct {
	City string
	User User
}
