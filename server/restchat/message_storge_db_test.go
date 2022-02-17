package restchat

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

func TestCreateDB(t *testing.T) {
	godotenv.Load(".env")
	connectDB := NewConnectDB(5432)
	var inter IMessageStorage = NewMessageStorageDB(connectDB)
	_, err := inter.Create(1, "новое")
	if err != nil {
		t.Errorf(err.Error())
	}
}

func TestGetLastDB(t *testing.T) {
	godotenv.Load(".env")
	connectDB := NewConnectDB(5432)
	var inter IMessageStorage = NewMessageStorageDB(connectDB)
	messageModel, err := inter.GetLast(2)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(messageModel)
}
