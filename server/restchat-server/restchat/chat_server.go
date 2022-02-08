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
	ChatServerGin := NewChatServerGin("localhost", 300, 8080)
	ChatServerGin.router = gin.Default()
	chat.sessionStorage = sessionStorage
	chat.messageStorage = messageStorage
	chat.userStorage = userStorage
	// Конфигурируем все эндпоинты
	ChatServerGin.router.POST("/api/user", userHandler(chat.userStorage))
	ChatServerGin.router.POST("/api/login", loginHandler(chat.sessionStorage))
	ChatServerGin.router.POST("/api/logout", logoutHandler(chat.sessionStorage))
	ChatServerGin.router.GET("/api/members", membersHandler(chat.sessionStorage))
	ChatServerGin.router.GET("/api/messages", messagesHandler(chat.messageStorage, ChatServerGin.maxLastMessages))
	ChatServerGin.router.POST("/api/message", messageHandler(chat.messageStorage))
}

func (chat *ChatServerGin) Run() error {
	if chat.router != nil {
		return fmt.Errorf("gin не сконфигурирован %v", chat.router)
	}
	chat.router.Run()
	return nil
}

func loginHandler(sessionStorage ISessionStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestTask := new(RequestTask)
		ctx.BindJSON(&requestTask)
		fmt.Println(requestTask.ApiToken)
		if err := ctx.BindJSON(&requestTask); err != nil {
			fmt.Errorf("Пустой JSON")
		}
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
		// session, _ := userStorage.Create(4000)
		// fmt.Println("handler func:", session.AuthToken)
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
		// urlParams:=ctx.Request.URL.Query()
		// message, _ := messageStorage.Create()
		// fmt.Println("handler func:", session.AuthToken)
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

// {api_token: "string"}
// {username: "string", password: "string"}
// {api_token: "string", text: "string"}
type RequestTask struct {
	ApiToken string `json:"api_token,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Text     string `json:"text,omitempty"`
}
