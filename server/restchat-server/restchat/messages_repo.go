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
	Create(user_id int, text string) (MessageModel, error)
	GetLast(n int) ([]MessageModel, error)
}

func NewMessagesMemRepo() *MessagesMemRepo {
	return &MessagesMemRepo{}
}

func (mmr *MessagesMemRepo) Create(user_id int, text string) (MessageModel, error) {

	id := 1 // добавить функцию получения id массива, выдавать последний и плюсовать.
	// Так же нужна функция, при выходе все участников чата отчищала массив сообщений
	lenlastmessages := len(mmr.Messages)
	mmr.Messages = append(mmr.Messages, MessageModel{ID: id, UserId: user_id, Text: text, TimeMessage: time.Now()})

	if lenlastmessages >= len(mmr.Messages) {
		return MessageModel{ID: 0, UserId: 0, Text: "", TimeMessage: time.Time{}}, fmt.Errorf("%s", "Не удалось добавить сообщение")
	}

	return mmr.Messages[id], fmt.Errorf("cообщение создалось: %v", mmr.Messages[id])
}

func (mmr *MessagesMemRepo) GetLast(n int) ([]MessageModel, error) {

	if mmr == nil {
		return mmr.Messages, fmt.Errorf("%s", "Нет сообщений")
	}

	sort.Slice(mmr.Messages, func(i, j int) (less bool) {
		return mmr.Messages[i].ID > mmr.Messages[j].ID
	})
	cplastmessages := mmr.Messages[0:n]

	return cplastmessages, fmt.Errorf("выгрузка удалась: %v", cplastmessages)
}
