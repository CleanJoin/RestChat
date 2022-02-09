package restchat

import (
	"fmt"
	"net/http"
	"regexp"
	"time"

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

type RequestUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type RequestApiToken struct {
	ApiToken string `json:"api_token"`
}
type RequestMessage struct {
	ApiToken string `json:"api_token"`
	Text     string `json:"text"`
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
	chat.router.POST("/api/login", loginHandler(chat.sessionStorage, chat.userStorage))
	chat.router.POST("/api/logout", logoutHandler(chat.sessionStorage))
	chat.router.GET("/api/members", membersHandler(chat.sessionStorage))
	chat.router.GET("/api/messages", messagesHandler(chat.messageStorage, chat.maxLastMessages))
	chat.router.POST("/api/message", messageHandler(chat.messageStorage))
	chat.router.GET("/api/health", heathHandler())

}

func heathHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"time":    time.Now().Format(time.RFC3339),
		})

	}
}
func (chat *ChatServerGin) Run() error {
	if chat.router == nil {
		return fmt.Errorf("gin не сконфигурирован %v", chat.router)
	}
	return chat.router.Run()
}

func loginHandler(sessionStorage ISessionStorage, userStorage IUserStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestUser := new(RequestUser)

		statusCode, ctx2, checkBadRequest := validateBadRequest(ctx, requestUser)

		if !checkBadRequest {
			ctx.JSON(statusCode, ctx2)
			return
		}

		if !validatenUserName(requestUser.Username) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Username"})
			return
		}

		if !validatePassword(requestUser.Password) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Password"})
			return
		}
		userMode, err := userStorage.GetByName(requestUser.Username)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		if !checkUserPassword(requestUser.Username, requestUser.Password, userStorage) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Не правильно введен пароль"})
			return
		}

		sessionModel, err := sessionStorage.Create(userMode.ID)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"auth_token": sessionModel.AuthToken, "member": gin.H{"id": sessionModel.ID, "name": userMode.Username}})
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
		requestUser := new(RequestUser)

		statusCode, ctx2, checkBadRequest := validateBadRequest(ctx, requestUser)

		if !checkBadRequest {
			ctx.JSON(statusCode, ctx2)
			return
		}

		if !validatenUserName(requestUser.Username) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Invalid Username"})
			return
		}

		if !validatePassword(requestUser.Password) {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Invalid Password"})
			return
		}
		userMode, err := userStorage.GetByName(requestUser.Username)
		if err == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		userMode, err = userStorage.Create(requestUser.Username, requestUser.Password)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"username": userMode.Username})
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

func validatenUserName(userName string) bool {
	var validPath = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString
	return validPath(userName) && len(userName) < 16
}

func validateMessage(text string) bool {
	return len(text) < 1024
}
func validatePassword(password string) bool {
	return len(password) < 32
}

func checkUserPassword(userName string, password string, userStorage IUserStorage) bool {
	userModel, err := userStorage.GetByName(userName)
	return err == nil && userModel.PasswordHash == new(PasswordHasherSha1).CalculateHash(password)
}

func validateBadRequest(ctx *gin.Context, requestData interface{}) (int, interface{}, bool) {

	err := ctx.BindJSON(&requestData)
	if err != nil {

		return http.StatusBadRequest, gin.H{"error": "Не содержит поля в запросе"}, false
	}
	return http.StatusOK, gin.H{"error": ""}, true
}
