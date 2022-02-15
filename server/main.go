package main

import (
	"fmt"
	"os"
	"restchat-server/restchat"
	"strconv"
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
	serverPort, _ := strconv.Atoi(os.Getenv("SERVER_PORT"))
	maxMessagesNum, _ := strconv.Atoi(os.Getenv("SERVER_MAX_MESSAGES"))

	fmt.Println("Server port env:", serverPort)
	fmt.Println("Max messages num env:", maxMessagesNum)

	chatServerGin := restchat.NewChatServerGin("localhost", serverPort, uint(maxMessagesNum))
	sessionStorage := restchat.NewSessionStorageMemory(new(restchat.TokenGeneratorUUID))
	usersStorage := restchat.NewUserStorageMemory(new(restchat.PasswordHasherSha1))
	messageStorage := restchat.NewMessageStorageMemory()
	chatServerGin.Use(sessionStorage, usersStorage, messageStorage)
	chatServerGin.Run()
}
