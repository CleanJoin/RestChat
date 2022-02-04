package restchat

import (
	"reflect"
	"testing"
	"time"
)

func TestGetLastMessageIdEmpty(t *testing.T) {
	msm := new(MessageStorageMemory)
	if getLastMessageId(msm) != 0 {
		t.Errorf("Есть идентификатор сообщения: %v\n", getLastMessageId(msm))
	}
}

func TestGetLastMessageId(t *testing.T) {
	msm := new(MessageStorageMemory)
	msm.Messages = append(msm.Messages, MessageModel{ID: 4, UserId: 1, Text: "Первое сообщение чата", Time: time.Now()})
	msm.Messages = append(msm.Messages, MessageModel{ID: 2, UserId: 3, Text: "Второе сообщение чата", Time: time.Now()})
	if getLastMessageId(msm) != 4 {
		t.Errorf("Не верное выводиться идентификатор последнего сообщения%v\n", getLastMessageId(msm))
	}
}

func TestCreateMessage(t *testing.T) {
	msm := new(MessageStorageMemory)
	msm.Messages = append(msm.Messages, MessageModel{ID: 1, UserId: 1, Text: "Первое сообщение чата", Time: time.Now()})
	msm.Messages = append(msm.Messages, MessageModel{ID: 4, UserId: 3, Text: "Второе сообщение чата", Time: time.Now()})
	user_id := uint(1)
	text := "Первое сообщение"
	outmessage := new(MessageStorageMemory)
	outmessage.Messages = append(outmessage.Messages, MessageModel{ID: 5, UserId: 1, Text: "Первое сообщение", Time: time.Now()})
	request, err := msm.Create(user_id, text)

	if request != outmessage.Messages[0] {
		t.Errorf("Ошибка не верный вывод из функции")
	}
	if err != nil {
		t.Errorf("Сообщение не создалось %v", err)
	}
}

func TestGetLastEmpty(t *testing.T) {
	msm := new(MessageStorageMemory)
	lastSequence := uint(3)
	request, err := msm.GetLast(lastSequence)
	if request != nil || len(request) != 0 {
		t.Errorf("Массив сообщений не пустой")
	}
	if err == nil {
		t.Errorf("Не получили список последних сообщений %v", err)
	}
}

func TestGetLast(t *testing.T) {
	msm := new(MessageStorageMemory)
	newTime := time.Now()

	msm.Messages = append(msm.Messages, MessageModel{ID: 1, UserId: 1, Text: "Первое сообщение чата", Time: newTime})
	msm.Messages = append(msm.Messages, MessageModel{ID: 4, UserId: 3, Text: "Второе сообщение чата", Time: newTime})
	msm.Messages = append(msm.Messages, MessageModel{ID: 5, UserId: 3, Text: "Второе сообщение чата", Time: newTime})
	msm.Messages = append(msm.Messages, MessageModel{ID: 7, UserId: 3, Text: "Второе сообщение чата", Time: newTime})

	lastSequence := uint(3)
	outmessage := new(MessageStorageMemory)
	outmessage.Messages = append(outmessage.Messages, MessageModel{ID: 7, UserId: 3, Text: "Второе сообщение чата", Time: newTime})
	outmessage.Messages = append(outmessage.Messages, MessageModel{ID: 5, UserId: 3, Text: "Второе сообщение чата", Time: newTime})
	outmessage.Messages = append(outmessage.Messages, MessageModel{ID: 4, UserId: 3, Text: "Второе сообщение чата", Time: newTime})

	request, err := msm.GetLast(lastSequence)
	if reflect.DeepEqual(request, outmessage.Messages) != true {
		t.Errorf("Не корректный список сообщений \n%v\n%v", request, outmessage.Messages)
	}
	if err != nil {
		t.Errorf("Не получили список последних сообщений %v", err)
	}
}

func TestIMessageStorage(t *testing.T) {
	inter := NewMessageStorageMemory()
	_, err := inter.Create(1, "text")
	if err != nil {
		t.Errorf("Сообщение не создалось%v", err)
	}
	_, err = inter.GetLast(1)
	if err != nil {
		t.Errorf("Ну получили последние сообщения %v", err)
	}

}
