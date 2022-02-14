package postgresDB

import (
	"restchat-server/restchat"

	"github.com/jackc/pgx/v4"
)

type MessageStorageDB struct {
	Messages []restchat.MessageModel
	connect  *pgx.Conn
}

type IMessageStorageDB interface {
	Create(userId uint, text string) (restchat.MessageModel, error)
	GetLast(n uint) ([]restchat.MessageModel, error)
}

func NewMessageStorageDB() *MessageStorageDB {
	msdb := new(MessageStorageDB)
	msdb.connect = connectDB()
	return msdb
}
