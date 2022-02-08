package restchat

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUse(t *testing.T) {
	chatServer := NewChatServerGin("localhost", 300, 8080)
	sessionStorage1 := NewSessionStorageMemory(new(TokenGeneratorUUID))
	usersstorage := NewUserStorageMemory(new(PasswordHasherSha1))
	messageStorage := NewMessageStorageMemory()
	chatServer.Use(sessionStorage1, usersstorage, messageStorage)

}
func TestRun(t *testing.T) {
	chatServer := NewChatServerGin("localhost", 300, 8080)
	router := chatServer.Run()
	if router != nil {
		t.Errorf("Gin не сконфигурирован  %v", router)
	}
}

func TestMessagesHandler(t *testing.T) {
	chatServer := NewChatServerGin("localhost", 300, 8080)
	sessionStorage1 := NewSessionStorageMemory(new(TokenGeneratorUUID))
	usersstorage := NewUserStorageMemory(new(PasswordHasherSha1))
	messageStorage := NewMessageStorageMemory()
	chatServer.Use(sessionStorage1, usersstorage, messageStorage)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:10000/api/messages", nil)
	fmt.Println(w, req)

}
