package restchat

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestGetLastMessageIdEmpty(t *testing.T) {
	msm := new(MessageStorageMemory)
	if getLastMessageId(msm) != 0 {
		t.Errorf("Есть идентификатор сообщения: %v\n", getLastMessageId(msm))
	}
	fmt.Printf("Нет сообщений: %v\n", getLastMessageId(msm))
}

func TestGetLastMessageId(t *testing.T) {
	msm := new(MessageStorageMemory)
	msm.Messages = append(msm.Messages, MessageModel{ID: 4, UserId: 1, Text: "Первое сообщение чата", TimeMessage: time.Now()})
	msm.Messages = append(msm.Messages, MessageModel{ID: 2, UserId: 3, Text: "Второе сообщение чата", TimeMessage: time.Now()})
	if getLastMessageId(msm) != 4 {
		t.Errorf("Не верное выводиться идентификатор последнего сообщения%v\n", getLastMessageId(msm))
	}
	fmt.Printf("Идентификатор последнего сообщения: %v\n", getLastMessageId(msm))
}

func TestCreateMessage(t *testing.T) {
	msm := new(MessageStorageMemory)
	msm.Messages = append(msm.Messages, MessageModel{ID: 1, UserId: 1, Text: "Первое сообщение чата", TimeMessage: time.Now()})
	msm.Messages = append(msm.Messages, MessageModel{ID: 4, UserId: 3, Text: "Второе сообщение чата", TimeMessage: time.Now()})
	user_id := uint(1)
	text := "Первое сообщение"
	outmessage := new(MessageStorageMemory)
	outmessage.Messages = append(outmessage.Messages, MessageModel{ID: 5, UserId: 1, Text: "Первое сообщение", TimeMessage: time.Now()})
	request, err := msm.Create(user_id, text)

	if request != outmessage.Messages[0] {
		t.Errorf("Ошибка не верный вывод из функции")
	}
	fmt.Println(err)
}

func TestGetLastEmpty(t *testing.T) {
	msm := new(MessageStorageMemory)
	lastSequence := uint(3)
	request, err := msm.GetLast(lastSequence)
	if request != nil || len(request) != 0 {
		t.Errorf("Массив сообщений не пустой")
	}
	fmt.Println(err)
}

func TestGetLast(t *testing.T) {
	msm := new(MessageStorageMemory)
	msm.Messages = append(msm.Messages, MessageModel{ID: 1, UserId: 1, Text: "Первое сообщение чата", TimeMessage: time.Now()})
	msm.Messages = append(msm.Messages, MessageModel{ID: 4, UserId: 3, Text: "Второе сообщение чата", TimeMessage: time.Now()})
	msm.Messages = append(msm.Messages, MessageModel{ID: 5, UserId: 3, Text: "Второе сообщение чата", TimeMessage: time.Now()})
	msm.Messages = append(msm.Messages, MessageModel{ID: 7, UserId: 3, Text: "Второе сообщение чата", TimeMessage: time.Now()})

	lastSequence := uint(3)

	outmessage := new(MessageStorageMemory)
	outmessage.Messages = append(outmessage.Messages, MessageModel{ID: 7, UserId: 3, Text: "Второе сообщение чата", TimeMessage: time.Now()})
	outmessage.Messages = append(outmessage.Messages, MessageModel{ID: 5, UserId: 3, Text: "Второе сообщение чата", TimeMessage: time.Now()})
	outmessage.Messages = append(outmessage.Messages, MessageModel{ID: 4, UserId: 3, Text: "Второе сообщение чата", TimeMessage: time.Now()})

	request, err := msm.GetLast(lastSequence)
	fmt.Println(outmessage.Messages)
	if reflect.DeepEqual(request, outmessage.Messages) != true {
		t.Errorf("Не корректный список сообщений \n%v\n%v", request, outmessage.Messages)
	}
	fmt.Println(err)
}
