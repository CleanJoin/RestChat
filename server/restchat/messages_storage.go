package restchat

import (
	"fmt"
	"sort"
	"time"
)

type MessageStorageMemory struct {
	Messages []MessageModel
	nextId   uint
}

type IMessageStorage interface {
	Create(userId uint, text string) (MessageModel, error)
	GetLast(n uint) ([]MessageModel, error)
}

func NewMessageStorageMemory() *MessageStorageMemory {
	msm := new(MessageStorageMemory)
	msm.nextId = 0
	return msm
}

func (messageStorage *MessageStorageMemory) Create(userId uint, text string) (MessageModel, error) {

	lenlastmessages := len(messageStorage.Messages)
	messageStorage.Messages = append(messageStorage.Messages, MessageModel{ID: messageStorage.nextId, UserId: userId, Text: text, Time: time.Now()})

	if lenlastmessages >= len(messageStorage.Messages) {
		return MessageModel{}, fmt.Errorf("%s", "Не удалось добавить сообщение")
	}
	messageStorage.nextId++
	return messageStorage.Messages[len(messageStorage.Messages)-1], nil
}

func (messageStorage *MessageStorageMemory) GetLast(n uint) ([]MessageModel, error) {

	if messageStorage == nil || len(messageStorage.Messages) == 0 {
		return messageStorage.Messages, fmt.Errorf("%s", "В памяти нет сообщений")
	}
	sort.Slice(messageStorage.Messages, func(i, j int) bool {
		return messageStorage.Messages[i].ID > messageStorage.Messages[j].ID
	})
	numMessages := minUint(len(messageStorage.Messages), n)
	copylastmessages := messageStorage.Messages[0:uint(numMessages)]
	return copylastmessages, nil
}

func minUint(len int, maxMessages uint) (min uint) {
	if uint(len) >= maxMessages {
		return maxMessages
	}
	return uint(len)
}
