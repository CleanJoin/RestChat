package restchat

import (
	"fmt"
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
		t.Errorf("Не верное выводиться id последего сообщения")
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
