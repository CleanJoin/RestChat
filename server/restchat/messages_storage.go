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
	Create(userId uint, text string) (MessageModel, error)
	GetLast(n uint) ([]MessageModel, error)
}

func NewMessageStorageMemory() *MessageStorageMemory {
	return new(MessageStorageMemory)
}

func getLastMessageId(messageStorage *MessageStorageMemory) uint {
	if messageStorage == nil || len(messageStorage.Messages) == 0 {
		return 0
	}
	sort.Slice(messageStorage.Messages, func(i, j int) bool {
		return messageStorage.Messages[i].ID > messageStorage.Messages[j].ID
	})
	return messageStorage.Messages[0].ID
}

func (messageStorage *MessageStorageMemory) Create(userId uint, text string) (MessageModel, error) {

	id := getLastMessageId(messageStorage)
	id++
	lenlastmessages := len(messageStorage.Messages)
	messageStorage.Messages = append(messageStorage.Messages, MessageModel{ID: id, UserId: userId, Text: text, Time: time.Now()})

	if lenlastmessages >= len(messageStorage.Messages) {
		return MessageModel{ID: 0, UserId: 0, Text: "", Time: time.Time{}}, fmt.Errorf("%s", "Не удалось добавить сообщение")
	}

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