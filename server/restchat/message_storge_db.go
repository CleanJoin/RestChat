package restchat

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

type MessageStorageDB struct {
	Messages []MessageModel
	connect  *pgxpool.Pool
}

func NewMessageStorageDB() *MessageStorageDB {
	msdb := new(MessageStorageDB)
	msdb.connect = connectDB()
	return msdb
}

func (messageStorageDB *MessageStorageDB) Create(userId uint, text string) (MessageModel, error) {
	query := `INSERT INTO "UserModel".messages (text,user_id) VALUES($1, $2)`

	commandTag, err := messageStorageDB.connect.Exec(context.Background(), query, text, userId)
	if err != nil {
		return MessageModel{}, fmt.Errorf("не удалось добавить сообщение %s", err.Error())
	}
	commandTag.Insert()
	fmt.Println(commandTag.String())
	return MessageModel{1, userId, text, time.Now()}, nil
}

func (messageStorageDB *MessageStorageDB) GetLast(n uint) ([]MessageModel, error) {
	query := `select * from "UserModel".messages u order by id desc  limit $1`
	commandTag := messageStorageDB.connect.QueryRow(context.Background(), query, n)
	commandTag.Scan(messageStorageDB.Messages)
	return messageStorageDB.Messages, nil
}
