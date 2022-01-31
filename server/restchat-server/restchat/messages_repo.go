package restchat

import (
	"fmt"
	"sort"
	"time"
)

type MessagesMemRepo struct {
	Messages []MessageModel
}

type IMessagesRepo interface {
	Create(user_id int, text string) (MessageModel, error)
	GetLast(n int) ([]MessageModel, error)
}

func NewMessagesMemRepo() *MessagesMemRepo {
	return &MessagesMemRepo{}
}

