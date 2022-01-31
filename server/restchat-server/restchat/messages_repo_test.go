package restchat

import (
	"fmt"
	"testing"
	"time"
)

func TestReceivelastIDMessageempty(t *testing.T) {
	mmr := new(*MessagesMemRepo)
	if ReceivelastIDMessage(*mmr) != 0 {
		t.Errorf("Массив сообщений не пустой")
	}
	fmt.Printf("Все хорошо!! массив пустой %v\n", ReceivelastIDMessage(*mmr))
}

func TestReceivelastIDMessage(t *testing.T) {
	mmr := MessagesMemRepo{}
	mmr.Messages = append(mmr.Messages, MessageModel{ID: 1, UserId: 1, Text: "Первое сообщение чата", TimeMessage: time.Now()})
	mmr.Messages = append(mmr.Messages, MessageModel{ID: 2, UserId: 3, Text: "Второу сообщение чата", TimeMessage: time.Now()})
	if ReceivelastIDMessage(&mmr) != 2 {
		t.Errorf("Не верное выводиться id последего сообщения")
	}
	fmt.Printf("Все хорошо!! Идентификатор последнего сообщения: %v\n", ReceivelastIDMessage(&mmr))
}
