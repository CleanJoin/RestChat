package restchat

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestgetLastMessageIdEmpty(t *testing.T) {
	mmr := new(MessagesMemRepo)
	if getLastMessageId(mmr) != 0 {
		t.Errorf("Массив сообщений не пустой")
	}
	fmt.Printf("Все хорошо!! массив пустой %v\n", getLastMessageId(mmr))
}

func TestgetLastMessageId(t *testing.T) {
	mmr := new(MessagesMemRepo)
	mmr.Messages = append(mmr.Messages, MessageModel{ID: 4, UserId: 1, Text: "Первое сообщение чата", TimeMessage: time.Now()})
	mmr.Messages = append(mmr.Messages, MessageModel{ID: 2, UserId: 3, Text: "Второу сообщение чата", TimeMessage: time.Now()})
	if getLastMessageId(mmr) != 4 {
		t.Errorf("Не верное выводиться id последнего сообщения")
	}
	fmt.Printf("Все хорошо!! Идентификатор последнего сообщения: %v\n", getLastMessageId(mmr))
}

func TestCreateEmpty(t *testing.T) {
	mmr := new(MessagesMemRepo)
	user_id := uint(1)
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
	user_id := uint(1)
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
	lastsqn := uint(3)
	request, err := mmr.GetLastMessages(lastsqn)
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

	lastsqn := uint(3)

	outmessage := new(MessagesMemRepo)
	outmessage.Messages = append(outmessage.Messages, MessageModel{ID: 7, UserId: 3, Text: "Второу сообщение чата", TimeMessage: time.Now()})
	outmessage.Messages = append(outmessage.Messages, MessageModel{ID: 5, UserId: 3, Text: "Второу сообщение чата", TimeMessage: time.Now()})
	outmessage.Messages = append(outmessage.Messages, MessageModel{ID: 4, UserId: 3, Text: "Второу сообщение чата", TimeMessage: time.Now()})

	request, err := mmr.GetLastMessages(lastsqn)
	fmt.Println(outmessage.Messages)
	if reflect.DeepEqual(request, outmessage.Messages) != true {
		t.Errorf("Не корретный список сообщений \n%v\n%v", request, outmessage.Messages)
	}
	fmt.Println(err)
}
