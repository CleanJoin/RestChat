package restchat

import (
	"reflect"
	"testing"
	"time"
)

func TestMinUint(t *testing.T) {

	len := []int{3, 4, 5, 7, 8, 9, 10}
	maxMessages := []uint{5, 5, 5, 5, 5, 5, 5}
	min := []uint{}
	out := []uint{3, 4, 5, 5, 5, 5, 5}

	for i := range len {
		min = append(min, minUint(len[i], maxMessages[i]))
	}
	if !reflect.DeepEqual(min, out) {
		t.Errorf("Не правильно опредлись минимальные элементы массива:\n %v\n %v\n", min, out)
	}
}

func TestGetLastMessageIdEmpty(t *testing.T) {
	messageStorage := new(MessageStorageMemory)
	if getLastMessageId(messageStorage) != 0 {
		t.Errorf("Есть идентификатор сообщения: %v\n", getLastMessageId(messageStorage))
	}
}

func TestGetLastMessageId(t *testing.T) {
	messageStorage := new(MessageStorageMemory)
	messageStorage.Messages = append(messageStorage.Messages, MessageModel{ID: 4, UserId: 1, Text: "Первое сообщение чата", Time: time.Now()})
	messageStorage.Messages = append(messageStorage.Messages, MessageModel{ID: 2, UserId: 3, Text: "Второе сообщение чата", Time: time.Now()})
	if getLastMessageId(messageStorage) != 4 {
		t.Errorf("Не верное выводиться идентификатор последнего сообщения%v\n", getLastMessageId(messageStorage))
	}
}

func TestCreateMessage(t *testing.T) {
	messageStorage := new(MessageStorageMemory)
	newTime := time.Now()

	messageStorage.Messages = append(messageStorage.Messages, MessageModel{ID: 1, UserId: 1, Text: "Первое сообщение чата", Time: newTime})
	messageStorage.Messages = append(messageStorage.Messages, MessageModel{ID: 4, UserId: 3, Text: "Второе сообщение чата", Time: newTime})
	user_id := uint(1)
	text := "Первое сообщение"
	outmessage := new(MessageStorageMemory)
	outmessage.Messages = append(outmessage.Messages, MessageModel{ID: 5, UserId: 1, Text: "Первое сообщение", Time: newTime})
	request, err := messageStorage.Create(user_id, text)

	if request.ID != outmessage.Messages[0].ID {
		t.Errorf("Ошибка не верный вывод из функции %v,%v", request, outmessage.Messages[0])
	}
	if err != nil {
		t.Errorf("Сообщение не создалось %v", err)
	}
}

func TestGetLastEmpty(t *testing.T) {
	messageStorage := new(MessageStorageMemory)
	lastSequence := uint(3)
	request, err := messageStorage.GetLast(lastSequence)
	if request != nil || len(request) != 0 {
		t.Errorf("Массив сообщений не пустой")
	}
	if err == nil {
		t.Errorf("Не получили список последних сообщений %v", err)
	}
}

func TestGetLast(t *testing.T) {
	messageStorage := new(MessageStorageMemory)
	newTime := time.Now()

	messageStorage.Messages = append(messageStorage.Messages, MessageModel{ID: 1, UserId: 1, Text: "Первое сообщение чата", Time: newTime})
	messageStorage.Messages = append(messageStorage.Messages, MessageModel{ID: 4, UserId: 3, Text: "Второе сообщение чата", Time: newTime})
	messageStorage.Messages = append(messageStorage.Messages, MessageModel{ID: 5, UserId: 3, Text: "Второе сообщение чата", Time: newTime})
	messageStorage.Messages = append(messageStorage.Messages, MessageModel{ID: 7, UserId: 3, Text: "Второе сообщение чата", Time: newTime})

	lastSequence := uint(3)
	outmessage := new(MessageStorageMemory)
	outmessage.Messages = append(outmessage.Messages, MessageModel{ID: 7, UserId: 3, Text: "Второе сообщение чата", Time: newTime})
	outmessage.Messages = append(outmessage.Messages, MessageModel{ID: 5, UserId: 3, Text: "Второе сообщение чата", Time: newTime})
	outmessage.Messages = append(outmessage.Messages, MessageModel{ID: 4, UserId: 3, Text: "Второе сообщение чата", Time: newTime})

	request, err := messageStorage.GetLast(lastSequence)
	if !reflect.DeepEqual(request, outmessage.Messages) {
		t.Errorf("Не корректный список сообщений \n%v\n%v", request, outmessage.Messages)
	}
	if err != nil {
		t.Errorf("Не получили список последних сообщений %v", err)
	}
}

func TestIMessageStorage(t *testing.T) {
	var inter IMessageStorage = NewMessageStorageMemory()
	_, err := inter.Create(1, "text")
	if err != nil {
		t.Errorf("Сообщение не создалось%v", err)
	}
	_, err = inter.GetLast(1)
	if err != nil {
		t.Errorf("Ну получили последние сообщения %v", err)
	}

}
