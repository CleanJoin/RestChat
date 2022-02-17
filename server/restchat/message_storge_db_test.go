package restchat

import (
	"fmt"
	"testing"
)

func TestCreateDB(t *testing.T) {

	var inter IMessageStorage = NewMessageStorageDB()
	_, err := inter.Create(1, "новое")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetLastDB(t *testing.T) {
	var inter IMessageStorage = NewMessageStorageDB()
	messageModel, err := inter.GetLast(2)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(messageModel)
}
