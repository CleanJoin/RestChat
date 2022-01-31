package restchat

import "time"

type UserModel struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
}

type SessionModel struct {
	ID         int    `json:"id"`
	Username   string `json:"name"`
	Auth_token string `json:"auth_token"`
}

type MessageModel struct {
	ID          int       `json:"id"`
	UserId      int       `json:",omitempty"`
	Text        string    `json:"text"`
	TimeMessage time.Time `json:"time"` //utc
}
