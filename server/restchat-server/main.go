package main

import (
	"restchat-server/restchat"
)

// @title           Swagger RestChat
// @version         1.0
// @description     This is a sample server Rest API Server Chat.
// @termsOfService  https://github.com/CleanJoin/RestChat/

// @contact.name   Github.com
// @contact.url    https://github.com/CleanJoin/RestChat/

// @host      localhost:8000
// @BasePath  /

func main() {

	chatServerGin := restchat.NewChatServerGin("localhost", 8080, 300)
	sessionStorage1 := restchat.NewSessionStorageMemory(new(restchat.TokenGeneratorUUID))
	usersstorage := restchat.NewUserStorageMemory(new(restchat.PasswordHasherSha1))
	messageStorage := restchat.NewMessageStorageMemory()
	chatServerGin.Use(sessionStorage1, usersstorage, messageStorage)
	chatServerGin.Run()

}
