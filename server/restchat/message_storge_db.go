package restchat

import (
	"github.com/jackc/pgx/v4"
)

type MessageStorageDB struct {
	Messages []MessageModel
	connect  *pgx.Conn
}

type IMessageStorageDB interface {
	Create(userId uint, text string) (MessageModel, error)
	GetLast(n uint) ([]MessageModel, error)
}

func NewMessageStorageDB() *MessageStorageDB {
	msdb := new(MessageStorageDB)
	msdb.connect = connectDB()
	return msdb
}
