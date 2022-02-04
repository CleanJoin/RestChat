package restchat

import (
	"fmt"
	"sort"
	"time"
)

type MessageStorageMemory struct {
	Messages []MessageModel
}

type IMessageStorage interface {
	Create(user_id uint, text string) (MessageModel, error)
	GetLast(n uint) ([]MessageModel, error)
}

func NewMessageStorageMemory() *MessageStorageMemory {
	return &MessageStorageMemory{}
}

func getLastMessageId(msm *MessageStorageMemory) uint {
	if msm == nil || len(msm.Messages) == 0 {
		return 0
	}
	sort.Slice(msm.Messages, func(i, j int) (less bool) {
		return msm.Messages[i].ID > msm.Messages[j].ID
	})
	return msm.Messages[0].ID
}

func (msm *MessageStorageMemory) Create(user_id uint, text string) (MessageModel, error) {

	id := getLastMessageId(msm)
	id++
	lenlastmessages := len(msm.Messages)
	msm.Messages = append(msm.Messages, MessageModel{ID: id, UserId: user_id, Text: text, TimeMessage: time.Now()})

	if lenlastmessages >= len(msm.Messages) {
		return MessageModel{ID: 0, UserId: 0, Text: "", TimeMessage: time.Time{}}, fmt.Errorf("%s", "Не удалось добавить сообщение")
	}

	return msm.Messages[len(msm.Messages)-1], fmt.Errorf("сообщение создалось: %v", msm.Messages[len(msm.Messages)-1])
}

func (msm *MessageStorageMemory) GetLast(n uint) ([]MessageModel, error) {

	if msm == nil || len(msm.Messages) == 0 {
		return msm.Messages, fmt.Errorf("%s", "В памяти нет сообщений")
	}

	sort.Slice(msm.Messages, func(i, j int) (less bool) {
		return msm.Messages[i].ID > msm.Messages[j].ID
	})
	copylastmessages := msm.Messages[0:n]

	return copylastmessages, fmt.Errorf("последние сообщения получены: %v", copylastmessages)
}
