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

func NewMessageStorageDB(iConnectDB IConnectDB) *MessageStorageDB {

	sdb := new(MessageStorageDB)
	sdb.connect = iConnectDB.Use()
	return sdb
}

func (messageStorageDB *MessageStorageDB) Create(userId uint, text string) (MessageModel, error) {
	var id uint
	query := `INSERT INTO "restchat".messages (userid,"text") VALUES($1, $2) RETURNING id;`
	row := messageStorageDB.connect.QueryRow(context.Background(), query, userId, text)
	err := row.Scan(&id)
	if err != nil {
		return MessageModel{}, fmt.Errorf(err.Error())
	}
	return MessageModel{id, userId, text, time.Now()}, nil
}

func (messageStorageDB *MessageStorageDB) GetLast(n uint) ([]MessageModel, error) {
	messageStorageDBNew := new(MessageStorageDB)
	query := `select * from "restchat".messages u order by id asc  limit $1`
	commandTag, err := messageStorageDB.connect.Query(context.Background(), query, n)
	if err != nil {
		return []MessageModel{}, fmt.Errorf(err.Error())
	}

	for commandTag.Next() {
		messageModel := new(MessageModel)
		err := commandTag.Scan(&messageModel.ID, &messageModel.UserId, &messageModel.Text, &messageModel.Time)
		messageStorageDBNew.Messages = append(messageStorageDBNew.Messages, MessageModel{messageModel.ID, messageModel.UserId, messageModel.Text, messageModel.Time})
		if err != nil {
			return []MessageModel{}, fmt.Errorf(err.Error())
		}
	}

	return messageStorageDBNew.Messages, nil
}
