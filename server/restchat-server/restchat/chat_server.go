package restchat

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChatServerGin struct {
	sessionStorage  ISessionStorage
	userStorage     IUserStorage
	messageStorage  IMessageStorage
	router          *gin.Engine
	host            string
	port            int
	maxLastMessages uint
}

type RequestTask struct {
	ApiToken string `json:"api_token,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Text     string `json:"text,omitempty"`
}

type IChatServer interface {
	Use(sessionStorage ISessionStorage, userStorage IUserStorage, messageStorage IMessageStorage)
	Run()
}

func NewChatServerGin(host string, port int, maxLastMessages uint) *ChatServerGin {
	chatServerGin := new(ChatServerGin)
	chatServerGin.host = host
	chatServerGin.port = port
	chatServerGin.maxLastMessages = maxLastMessages
	return chatServerGin
}

func (chat *ChatServerGin) Use(sessionStorage ISessionStorage, userStorage IUserStorage, messageStorage IMessageStorage) {
	// ChatServerGin := NewChatServerGin("localhost", 300, 8080)
	chat.router = gin.Default()
	chat.sessionStorage = sessionStorage
	chat.messageStorage = messageStorage
	chat.userStorage = userStorage
	// Конфигурируем все эндпоинты
	chat.router.POST("/api/user", userHandler(chat.userStorage))
	chat.router.GET("/api/login", loginHandler(chat.sessionStorage))
	chat.router.POST("/api/logout", logoutHandler(chat.sessionStorage))
	chat.router.GET("/api/members", membersHandler(chat.sessionStorage))
	chat.router.GET("/api/messages", messagesHandler(chat.messageStorage, chat.maxLastMessages))
	chat.router.POST("/api/message", messageHandler(chat.messageStorage))
}

func (chat *ChatServerGin) Run() error {
	if chat.router == nil {
		return fmt.Errorf("gin не сконфигурирован %v", chat.router)
	}
	return chat.router.Run()

}

func loginHandler(sessionStorage ISessionStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestTask := new(RequestTask)
		ctx.BindJSON(&requestTask)
		fmt.Println(requestTask.ApiToken)
		session, _ := sessionStorage.Create(4000)
		fmt.Println("handler func:", session.AuthToken)
		ctx.JSON(http.StatusOK, gin.H{
			"api_token": requestTask.ApiToken,
		})
	}
}
func logoutHandler(sessionStorage ISessionStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session, _ := sessionStorage.Create(4000)
		fmt.Println("handler func:", session.AuthToken)
		ctx.JSON(http.StatusOK, gin.H{
			"api_token": session.AuthToken,
		})
	}
}
func userHandler(userStorage IUserStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{
			"api_token": "session.AuthToken",
		})
	}
}
func membersHandler(sessionStorage ISessionStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		members, _ := sessionStorage.GetOnlineUserIds()
		if len(members) == 0 {
			ctx.IndentedJSON(http.StatusNotFound, members)
			return
		}
		ctx.IndentedJSON(http.StatusOK, members)
	}
}
func messageHandler(messageStorage IMessageStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.JSON(http.StatusOK, gin.H{
			"api_token": "MessageStorageMemory",
		})
	}
}
func messagesHandler(messageStorage IMessageStorage, maxLastMessages uint) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		messages, _ := messageStorage.GetLast(maxLastMessages)
		if len(messages) == 0 {
			ctx.IndentedJSON(http.StatusNotFound, messages)
			return
		}
		ctx.IndentedJSON(http.StatusOK, messages)
	}
}
