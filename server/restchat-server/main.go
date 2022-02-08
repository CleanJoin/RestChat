package main

import (
	"restchat-server/restchat"
)

func main() {

	chatServerGin := restchat.NewChatServerGin("localhost", 8080, 300)
	sessionStorage1 := restchat.NewSessionStorageMemory(new(restchat.TokenGeneratorUUID))
	usersstorage := restchat.NewUserStorageMemory(new(restchat.PasswordHasherSha1))
	messageStorage := restchat.NewMessageStorageMemory()
	chatServerGin.Use(sessionStorage1, usersstorage, messageStorage)
	chatServerGin.Run()

}
