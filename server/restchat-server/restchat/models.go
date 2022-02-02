package restchat

import "time"

type UserModel struct {
	ID           uint   `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
}

type SessionModel struct {
	ID         uint   `json:"id"`
	UserId     uint   `json:"name"`
	Auth_token string `json:"auth_token"`
}

type MessageModel struct {
	ID          uint      `json:"id"`
	UserId      uint      `json:",omitempty"`
	Text        string    `json:"text"`
	TimeMessage time.Time `json:"time"` //utc
}
