package restchat

import "time"

type UserModel struct {
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
}

type SessionModel struct {
	ID       uint   `json:"id"`
	UserId   uint   `json:"name"`
	ApiToken string `json:"api_token"`
}

type MessageModel struct {
	ID     uint      `json:"id"`
	UserId uint      `json:",omitempty"`
	Text   string    `json:"text"`
	Time   time.Time `json:"time"` //utc
}
