package restchat

import "time"

type UserModel struct {
	ID           uint   `json:"id" db:"id"`
	Username     string `json:"username" db:"username"`
	PasswordHash string `json:"password" db:"password"`
}

type SessionModel struct {
	ID       uint   `json:"id"`
	UserId   uint   `json:"name"`
	ApiToken string `json:"api_token"`
}

type MessageModel struct {
	ID     uint      `json:"id" db:"id"`
	UserId uint      `json:" ,omitempty" db:"userid"`
	Text   string    `json:"text" db:"text"`
	Time   time.Time `json:"time" db:"message_time"` //utc
}
