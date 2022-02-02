package restchat

import (
	"fmt"
	"sort"
	"time"
)

type MessagesMemRepo struct {
	Messages []MessageModel
}

type IMessagesRepo interface {
	Create(user_id uint, text string) (MessageModel, error)
	GetLast(n int) ([]MessageModel, error)
}

func NewMessagesMemRepo() *MessagesMemRepo {
	return &MessagesMemRepo{}
}

func getLastMessageId(mmr *MessagesMemRepo) uint {
	if mmr == nil || len(mmr.Messages) == 0 {
		return 0
	}
	sort.Slice(mmr.Messages, func(i, j int) (less bool) {
		return mmr.Messages[i].ID > mmr.Messages[j].ID
	})
	return mmr.Messages[0].ID
}

// Нужно реализоавать функцию которая будет отчищать массим сообщений, при выходе всех пользователей
// func DeleteAllMessagesMemRepo(mmr *MessagesMemRepo) *MessagesMemRepo {
// }

func (mmr *MessagesMemRepo) Create(user_id uint, text string) (MessageModel, error) {

	id := getLastMessageId(mmr)
	id++
	lenlastmessages := len(mmr.Messages)
	mmr.Messages = append(mmr.Messages, MessageModel{ID: id, UserId: user_id, Text: text, TimeMessage: time.Now()})

	if lenlastmessages >= len(mmr.Messages) {
		return MessageModel{ID: 0, UserId: 0, Text: "", TimeMessage: time.Time{}}, fmt.Errorf("%s", "Не удалось добавить сообщение")
	}

	return mmr.Messages[len(mmr.Messages)-1], fmt.Errorf("cообщение создалось: %v", mmr.Messages[len(mmr.Messages)-1])
}

func (mmr *MessagesMemRepo) GetLastMessages(n uint) ([]MessageModel, error) {

	if mmr == nil || len(mmr.Messages) == 0 {
		return mmr.Messages, fmt.Errorf("%s", "В памяти нет сообщений")
	}

	sort.Slice(mmr.Messages, func(i, j int) (less bool) {
		return mmr.Messages[i].ID > mmr.Messages[j].ID
	})
	copylastmessages := mmr.Messages[0:n]

	return copylastmessages, fmt.Errorf("выгрузка удалась: %v", copylastmessages)
}
