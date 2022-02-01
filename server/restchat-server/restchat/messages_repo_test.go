package restchat

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestReceivelastIDMessageEmpty(t *testing.T) {
	mmr := new(MessagesMemRepo)
	if ReceivelastIDMessage(mmr) != 0 {
		t.Errorf("Массив сообщений не пустой")
	}
	fmt.Printf("Все хорошо!! массив пустой %v\n", ReceivelastIDMessage(mmr))
}

func TestReceivelastIDMessage(t *testing.T) {
	mmr := new(MessagesMemRepo)
	mmr.Messages = append(mmr.Messages, MessageModel{ID: 4, UserId: 1, Text: "Первое сообщение чата", TimeMessage: time.Now()})
	mmr.Messages = append(mmr.Messages, MessageModel{ID: 2, UserId: 3, Text: "Второу сообщение чата", TimeMessage: time.Now()})
	if ReceivelastIDMessage(mmr) != 4 {
		t.Errorf("Не верное выводиться id последнего сообщения")
	}
	fmt.Printf("Все хорошо!! Идентификатор последнего сообщения: %v\n", ReceivelastIDMessage(mmr))
}

func TestCreateEmpty(t *testing.T) {
	mmr := new(MessagesMemRepo)
	user_id := 1
	text := "Первое сообщение"
	outmessage := new(MessagesMemRepo)
	outmessage.Messages = append(outmessage.Messages, MessageModel{ID: 1, UserId: user_id, Text: text, TimeMessage: time.Now()})
	request, err := mmr.Create(user_id, text)
	if request != outmessage.Messages[0] {
		t.Errorf("Ошибка не верный вывод из функции")
	}
	fmt.Println(err)
}

func TestCreate(t *testing.T) {
	mmr := new(MessagesMemRepo)
	mmr.Messages = append(mmr.Messages, MessageModel{ID: 1, UserId: 1, Text: "Первое сообщение чата", TimeMessage: time.Now()})
	mmr.Messages = append(mmr.Messages, MessageModel{ID: 4, UserId: 3, Text: "Второу сообщение чата", TimeMessage: time.Now()})
	user_id := 1
	text := "Первое сообщение"
	outmessage := new(MessagesMemRepo)
	outmessage.Messages = append(outmessage.Messages, MessageModel{ID: 5, UserId: 1, Text: "Первое сообщение", TimeMessage: time.Now()})
	request, err := mmr.Create(user_id, text)

	if request != outmessage.Messages[0] {
		t.Errorf("Ошибка не верный вывод из функции")
	}
	fmt.Println(err)
}

func TestGetLastEmpty(t *testing.T) {
	mmr := new(MessagesMemRepo)
	lastsqn := 3
	request, err := mmr.GetLast(lastsqn)
	if request != nil || len(request) != 0 {
		t.Errorf("Массив сообщений не пустой")
	}
	fmt.Println(err)
}

func TestGetLast(t *testing.T) {
	mmr := new(MessagesMemRepo)
	mmr.Messages = append(mmr.Messages, MessageModel{ID: 1, UserId: 1, Text: "Первое сообщение чата", TimeMessage: time.Now()})
	mmr.Messages = append(mmr.Messages, MessageModel{ID: 4, UserId: 3, Text: "Второу сообщение чата", TimeMessage: time.Now()})
	mmr.Messages = append(mmr.Messages, MessageModel{ID: 5, UserId: 3, Text: "Второу сообщение чата", TimeMessage: time.Now()})
	mmr.Messages = append(mmr.Messages, MessageModel{ID: 7, UserId: 3, Text: "Второу сообщение чата", TimeMessage: time.Now()})

	lastsqn := 3

	outmessage := new(MessagesMemRepo)
	outmessage.Messages = append(outmessage.Messages, MessageModel{ID: 7, UserId: 3, Text: "Второу сообщение чата", TimeMessage: time.Now()})
	outmessage.Messages = append(outmessage.Messages, MessageModel{ID: 5, UserId: 3, Text: "Второу сообщение чата", TimeMessage: time.Now()})
	outmessage.Messages = append(outmessage.Messages, MessageModel{ID: 4, UserId: 3, Text: "Второу сообщение чата", TimeMessage: time.Now()})

	request, err := mmr.GetLast(lastsqn)
	fmt.Println(outmessage.Messages)
	if reflect.DeepEqual(request, outmessage.Messages) != true {
		t.Errorf("Не корретный список сообщений \n%v\n%v", request, outmessage.Messages)
	}
	fmt.Println(err)
}
