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
type RequestMembers struct {
	Members []struct {
		Id   uint   `json:"id"`
		Name string `json:"name"`
	} `json:"members"`
}
type RequestMessages struct {
	Messages []struct {
		Id         uint      `json:"id"`
		MemberName string    `json:"member_name"`
		Text       string    `json:"text"`
		Time       time.Time `json:"time"`
	} `json:"messages"`
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
	chat.router.GET("/api/members", membersHandler(chat.sessionStorage, chat.userStorage))
	chat.router.GET("/api/messages", messagesHandler(chat.messageStorage, chat.userStorage, chat.sessionStorage, chat.maxLastMessages))
	chat.router.POST("/api/message", messageHandler(chat.messageStorage, chat.userStorage, chat.sessionStorage))
	chat.router.GET("/api/health", heathHandler())

}

func heathHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, gin.H{
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

		statusCode, ctx2, checkBadRequest := validateClientRequest(ctx, requestUser)

		if !checkBadRequest {
			ctx.IndentedJSON(statusCode, ctx2)
			return
		}

		if !validatenUserName(requestUser.Username) {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Username"})
			return
		}

		if !validatePassword(requestUser.Password) {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Password"})
			return
		}
		userMode, err := userStorage.GetByName(requestUser.Username)
		if err != nil {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		if !checkUserPassword(requestUser.Username, requestUser.Password, userStorage) {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Не правильно введен пароль"})
			return
		}

		sessionModel, err := sessionStorage.Create(userMode.ID)
		if err != nil {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusOK, gin.H{"auth_token": sessionModel.AuthToken, "member": gin.H{"id": sessionModel.ID, "name": userMode.Username}})
	}
}

func logoutHandler(sessionStorage ISessionStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestApiToken := new(RequestApiToken)

		statusCode, ctx2, checkBadRequest := validateClientRequest(ctx, requestApiToken)
		if !checkBadRequest {
			ctx.IndentedJSON(statusCode, ctx2)
			return
		}
		_, err := sessionStorage.GetUserId(requestApiToken.ApiToken)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		err = sessionStorage.Delete(requestApiToken.ApiToken)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.IndentedJSON(http.StatusOK, gin.H{})
	}
}

func userHandler(userStorage IUserStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestUser := new(RequestUser)

		statusCode, ctx2, checkBadRequest := validateClientRequest(ctx, requestUser)

		if !checkBadRequest {
			ctx.IndentedJSON(statusCode, ctx2)
			return
		}

		if !validatenUserName(requestUser.Username) {
			ctx.IndentedJSON(http.StatusForbidden, gin.H{"error": "Invalid Username"})
			return
		}

		if !validatePassword(requestUser.Password) {
			ctx.IndentedJSON(http.StatusForbidden, gin.H{"error": "Invalid Password"})
			return
		}
		userMode, err := userStorage.GetByName(requestUser.Username)
		if err == nil {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Пользователь существует в базе"})
			return
		}
		userMode, err = userStorage.Create(requestUser.Username, requestUser.Password)
		if err != nil {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusCreated, gin.H{"username": userMode.Username})
	}

}

func membersHandler(sessionStorage ISessionStorage, userStorage IUserStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		requestApiToken := new(RequestApiToken)

		statusCode, ctx2, checkBadRequest := validateClientRequest(ctx, requestApiToken)
		if !checkBadRequest {
			ctx.IndentedJSON(statusCode, ctx2)
			return
		}

		userId, err := sessionStorage.GetOnlineUserIds()
		if err != nil {
			type Newmembers struct {
				Members []string `json:"members"`
			}
			empty := make([]string, 0)
			emptyMembers := Newmembers{Members: empty}
			ctx.IndentedJSON(http.StatusOK, emptyMembers)
			return
		}

		users, err := userStorage.GetByIds(userId)
		if err != nil {
			ctx.IndentedJSON(http.StatusOK, gin.H{"members": userId})
			return
		}
		newMembers := new(RequestMembers)
		for _, u := range users {

			newMembers.Members = append(newMembers.Members, struct {
				Id   uint   "json:\"id\""
				Name string "json:\"name\""
			}{u.ID, u.Username})
		}
		ctx.IndentedJSON(http.StatusOK, newMembers)
	}
}

func messagesHandler(messageStorage IMessageStorage, userStorage IUserStorage, sessionStorage ISessionStorage, maxLastMessages uint) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		requestApiToken := new(RequestMessage)

		statusCode, ctx2, checkBadRequest := validateClientRequest(ctx, requestApiToken)
		if !checkBadRequest {
			ctx.IndentedJSON(statusCode, ctx2)
			return
		}
		_, err := sessionStorage.GetUserId(requestApiToken.ApiToken)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		messageModel, err := messageStorage.GetLast(maxLastMessages)
		if err != nil {
			type NewMessages struct {
				Messages []string `json:"messages"`
			}
			empty := make([]string, 0)
			emptyMembers := NewMessages{Messages: empty}
			ctx.IndentedJSON(http.StatusOK, emptyMembers)
			return
		}

		newMessages := new(RequestMessages)
		for _, u := range messageModel {
			userModel, _ := userStorage.GetById(u.UserId)
			newMessages.Messages = append(newMessages.Messages, struct {
				Id         uint      "json:\"id\""
				MemberName string    "json:\"member_name\""
				Text       string    "json:\"text\""
				Time       time.Time "json:\"time\""
			}{u.ID, userModel.Username, u.Text, u.Time})

		}
		ctx.IndentedJSON(http.StatusOK, newMessages.Messages)
	}
}

func messageHandler(messageStorage IMessageStorage, userStorage IUserStorage, sessionStorage ISessionStorage) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestMessage := new(RequestMessage)
		statusCode, ctx2, checkBadRequest := validateClientRequest(ctx, requestMessage)
		if !checkBadRequest {
			ctx.IndentedJSON(statusCode, ctx2)
			return
		}

		if !validateMessage(requestMessage.Text) {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "сообщение больше 1024 символов"})
			return
		}
		if requestMessage.Text == "" {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "сообщение пустое"})
			return
		}
		userId, err := sessionStorage.GetUserId(requestMessage.ApiToken)
		if err != nil {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		messageModel, err := messageStorage.Create(userId, requestMessage.Text)
		if err != nil {
			ctx.IndentedJSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}

		userModel, err := userStorage.GetById(messageModel.UserId)
		if err != nil {
			ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		newMessages := new(RequestMessages)
		newMessages.Messages = append(newMessages.Messages, struct {
			Id         uint      "json:\"id\""
			MemberName string    "json:\"member_name\""
			Text       string    "json:\"text\""
			Time       time.Time "json:\"time\""
		}{messageModel.ID, userModel.Username, messageModel.Text, messageModel.Time})
		ctx.IndentedJSON(http.StatusOK, newMessages.Messages)
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

func validateClientRequest(ctx *gin.Context, requestData interface{}) (int, interface{}, bool) {

	err := ctx.BindJSON(&requestData)
	if err != nil {

		return http.StatusBadRequest, gin.H{"error": "Не содержит поля в запросе"}, false
	}
	return http.StatusOK, gin.H{"error": ""}, true
}
