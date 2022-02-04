package main

import (
	"fmt"
	"net/http"
	"restchat-server/restchat"

	"github.com/gin-gonic/gin"
)

func main() {

	session_storage := restchat.NewSessionStorageMemory(
		new(restchat.TokenGeneratorUUID),
	)

	session_storage.Create(1000)
	session_storage.Create(2000)

	router := gin.Default()
	router.GET("/api/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"error":   "Not implemented YET",
			"message": "Kill",
		})
	})

	router.GET("/api/user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"error":   "Not implemented YET",
			"message": "What your name?",
		})
	})

	router.GET("/api/login", login_handler(session_storage))

	router.Run()
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
