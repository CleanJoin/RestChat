package main

import (
	"fmt"
	"net/http"
	"restchat-server/restchat"

	"github.com/gin-gonic/gin"
)

func main() {

	// session_storage := restchat.NewSessionStorageMemory(
	// 	new(restchat.TokenGeneratorUUID),
	// )

	chatServerGin := restchat.NewChatServerGin("localhost", 8080, 300)
	sessionStorage1 := restchat.NewSessionStorageMemory(new(restchat.TokenGeneratorUUID))
	usersstorage := restchat.NewUserStorageMemory(new(restchat.PasswordHasherSha1))
	messageStorage := restchat.NewMessageStorageMemory()
	chatServerGin.Use(sessionStorage1, usersstorage, messageStorage)
	chatServerGin.Run()
	// router := gin.Default()
	// router.GET("/api/", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"error":   "Not implemented YET",
	// 		"message": "Kill",
	// 	})
	// })

	// router.GET("/api/user", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"error":   "Not implemented YET",
	// 		"message": "What your name?",
	// 	})
	// })

	// router.GET("/api/login", login_handler(session_storage))

	// router.GET("/api/health", func(ctx *gin.Context) {
	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"success": true,
	// 		"time":    time.Now().Format(time.RFC3339),
	// 	})
	// })

	// router.Run()
}

func login_handler(session_storage restchat.ISessionStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session, _ := session_storage.Create(4000)
		fmt.Println("handler func:", session.AuthToken)
		ctx.JSON(http.StatusOK, gin.H{
			"api_token": session.AuthToken,
		})
	}
}
