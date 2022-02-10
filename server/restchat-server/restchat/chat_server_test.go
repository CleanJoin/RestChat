package restchat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestUse(t *testing.T) {
	chatServer := NewChatServerGin("localhost", 300, 8080)
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))
	usersstorage := NewUserStorageMemory(new(PasswordHasherSha1))
	messageStorage := NewMessageStorageMemory()
	chatServer.Use(sessionStorage, usersstorage, messageStorage)
	if chatServer.router == nil {
		t.Errorf("router не сконфигурирован  %v", chatServer.router)
	}

}

func TestRun(t *testing.T) {
	chatServer := NewChatServerGin("localhost", 300, 8080)
	chatServer.Run()
	if chatServer.router != nil {
		t.Errorf("router сконфигурирован  %v", chatServer.router)
	}
}

func TestLogoutHandler(t *testing.T) {
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))
	usersstorage := NewUserStorageMemory(new(PasswordHasherSha1))
	messageStorage := NewMessageStorageMemory()

	chatServer := NewChatServerGin("localhost", 300, 8080)
	chatServer.Use(sessionStorage, usersstorage, messageStorage)
	usersstorage.Create("Andrey", "fghfghfghfgh")
	sessionStorage.Create(1)
	values := map[string]string{"api_token": sessionStorage.Sessions[0].AuthToken}
	jsonValue, _ := json.Marshal(values)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/logout", bytes.NewBuffer(jsonValue))

	chatServer.router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{}`, w.Body.String())
}
func TestLogoutHandlerNotDelete(t *testing.T) {
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))
	usersstorage := NewUserStorageMemory(new(PasswordHasherSha1))
	messageStorage := NewMessageStorageMemory()

	chatServer := NewChatServerGin("localhost", 300, 8080)
	chatServer.Use(sessionStorage, usersstorage, messageStorage)

	values := map[string]string{"api_token": "uy324uy34uyi234"}
	jsonValue, _ := json.Marshal(values)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/logout", bytes.NewBuffer(jsonValue))

	chatServer.router.ServeHTTP(w, req)
	if http.StatusNotFound != w.Code {
		t.Errorf("%v", w.Body) //{"error": "не удалось удалить сессию пользователя"}
	}
}

func TestLogoutHandlerApiTokenBadRequest(t *testing.T) {
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))
	usersstorage := NewUserStorageMemory(new(PasswordHasherSha1))
	messageStorage := NewMessageStorageMemory()

	chatServer := NewChatServerGin("localhost", 300, 8080)
	chatServer.Use(sessionStorage, usersstorage, messageStorage)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/login", strings.NewReader(`{"api_token": }`))

	chatServer.router.ServeHTTP(w, req)
	if http.StatusBadRequest != w.Code {
		t.Errorf("%v", w.Body) //{"error":"Не содержит поля в запросе"} api_token
	}
}

func TestUserHandler(t *testing.T) {
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))
	usersstorage := NewUserStorageMemory(new(PasswordHasherSha1))
	messageStorage := NewMessageStorageMemory()

	chatServer := NewChatServerGin("localhost", 300, 8080)
	chatServer.Use(sessionStorage, usersstorage, messageStorage)

	values := map[string]string{"username": "Andrey", "password": "fghfghfghfgh"}
	jsonValue, _ := json.Marshal(values)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/user", bytes.NewBuffer(jsonValue))

	chatServer.router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, fmt.Sprintf(`{"username":"%s"}`, usersstorage.Users[0].Username), w.Body.String())
}

func TestUserHandlerInStorage(t *testing.T) {
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))
	usersstorage := NewUserStorageMemory(new(PasswordHasherSha1))
	messageStorage := NewMessageStorageMemory()
	usersstorage.Create("Andrey", "fghfghfghfgh")
	chatServer := NewChatServerGin("localhost", 300, 8080)
	chatServer.Use(sessionStorage, usersstorage, messageStorage)

	values := map[string]string{"username": "Andrey", "password": "fghfghfghfgh"}
	jsonValue, _ := json.Marshal(values)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/user", bytes.NewBuffer(jsonValue))

	chatServer.router.ServeHTTP(w, req)
	if http.StatusForbidden == w.Code {
		t.Errorf("%v", w.Body) //{"error":"Пользователь существует в базе"}
	}
}

func TestUserHandlerPasswordBadRequest(t *testing.T) {
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))
	usersstorage := NewUserStorageMemory(new(PasswordHasherSha1))
	messageStorage := NewMessageStorageMemory()

	chatServer := NewChatServerGin("localhost", 300, 8080)
	chatServer.Use(sessionStorage, usersstorage, messageStorage)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/login", strings.NewReader(`{"username": "Andrey1", "password": }`))

	chatServer.router.ServeHTTP(w, req)
	if http.StatusBadRequest != w.Code {
		t.Errorf("%v", w.Body) //{"error":"Не содержит поля в запросе"} password
	}
}

func TestUserHandlerUserNameBadRequest(t *testing.T) {
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))
	usersstorage := NewUserStorageMemory(new(PasswordHasherSha1))
	messageStorage := NewMessageStorageMemory()

	chatServer := NewChatServerGin("localhost", 300, 8080)
	chatServer.Use(sessionStorage, usersstorage, messageStorage)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/login", strings.NewReader(`{"username": "Andrey1", "password": }`))

	chatServer.router.ServeHTTP(w, req)
	if http.StatusBadRequest != w.Code {
		t.Errorf("%v", w.Body) //{"error":"Не содержит поля в запросе"} UserName
	}
}

func TestLoginHandler(t *testing.T) {
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))
	usersstorage := NewUserStorageMemory(new(PasswordHasherSha1))
	messageStorage := NewMessageStorageMemory()
	usersstorage.Create("Andrey", "fghfghfghfgh")

	chatServer := NewChatServerGin("localhost", 300, 8080)
	chatServer.Use(sessionStorage, usersstorage, messageStorage)

	values := map[string]string{"username": "Andrey", "password": "fghfghfghfgh"}
	jsonValue, _ := json.Marshal(values)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonValue))

	chatServer.router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(`{"auth_token":"%s","member":{"id":%d,"name":"%s"}}`, sessionStorage.Sessions[0].AuthToken, 1, usersstorage.Users[0].Username), w.Body.String())
}

func TestLoginHandlerPassword(t *testing.T) {
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))
	usersstorage := NewUserStorageMemory(new(PasswordHasherSha1))
	messageStorage := NewMessageStorageMemory()
	usersstorage.Create("Andrey", "fghfghfghfgh")

	chatServer := NewChatServerGin("localhost", 300, 8080)
	chatServer.Use(sessionStorage, usersstorage, messageStorage)

	values := map[string]string{"username": "Andrey", "password": "fghfghfghfghdfdf"}
	jsonValue, _ := json.Marshal(values)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonValue))

	chatServer.router.ServeHTTP(w, req)
	if http.StatusUnauthorized != w.Code {
		t.Errorf("%v", w.Body) //{"error":"Не правильно введен пароль"}
	}
}

func TestLoginHandlerUserName(t *testing.T) {
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))
	usersstorage := NewUserStorageMemory(new(PasswordHasherSha1))
	messageStorage := NewMessageStorageMemory()
	usersstorage.Create("Andrey", "fghfghfghfgh")

	chatServer := NewChatServerGin("localhost", 300, 8080)
	chatServer.Use(sessionStorage, usersstorage, messageStorage)

	values := map[string]string{"username": "Andrey1", "password": "fghfghfghfgh"}
	jsonValue, _ := json.Marshal(values)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(jsonValue))

	chatServer.router.ServeHTTP(w, req)
	if http.StatusUnauthorized != w.Code {
		t.Errorf("%v", w.Body) //{"error":"не нашелся пользователь по указанному Username: "}
	}
}

func TestLoginHandlerPasswordBadRequest(t *testing.T) {
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))
	usersstorage := NewUserStorageMemory(new(PasswordHasherSha1))
	messageStorage := NewMessageStorageMemory()
	usersstorage.Create("Andrey", "fghfghfghfgh")

	chatServer := NewChatServerGin("localhost", 300, 8080)
	chatServer.Use(sessionStorage, usersstorage, messageStorage)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/login", strings.NewReader(`{"username": "Andrey1", "password": }`))

	chatServer.router.ServeHTTP(w, req)
	if http.StatusBadRequest != w.Code {
		t.Errorf("%v", w.Body) //{"error":"Не содержит поля в запросе"} password
	}
}

func TestLoginHandlerUserNameBadRequest(t *testing.T) {
	sessionStorage := NewSessionStorageMemory(new(TokenGeneratorUUID))
	usersstorage := NewUserStorageMemory(new(PasswordHasherSha1))
	messageStorage := NewMessageStorageMemory()
	usersstorage.Create("Andrey", "fghfghfghfgh")

	chatServer := NewChatServerGin("localhost", 300, 8080)
	chatServer.Use(sessionStorage, usersstorage, messageStorage)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/login", strings.NewReader(`{"username": , "password":"fghfghfghfgh"}`))

	chatServer.router.ServeHTTP(w, req)
	if http.StatusBadRequest != w.Code {
		t.Errorf("%v", w.Body) //{"error":"Не содержит поля в запросе"} UserName
	}
}

func TestValidatenUserName(t *testing.T) {
	userName := "Andrey"
	if !validatenUserName(userName) {
		t.Errorf("не корректный userName")
	}
}

func TestValidatenUserNameMore(t *testing.T) {
	userName := "Andreydsfdsfsdfdsfsdfdsfdsfdsfdsfdsfsdfdsfdsfdsfdsfdsfdsfdsfdsfdsfdsfds"
	if validatenUserName(userName) {
		t.Errorf("Длина userName меньше 16 символов")
	}
}

func TestValidatenUserNameNotCorrect(t *testing.T) {
	userName := "Andreyd!"
	if validatenUserName(userName) {
		t.Errorf("userName корректный")
	}
}

func TestValidateMessageLess(t *testing.T) {
	text := "Jdsfsdfdsfsdfdsfdsfsdfs"
	if !validateMessage(text) {
		t.Errorf("Длина сообщения больше 1024 символов")
	}
}

func TestValidateMessageMore(t *testing.T) {
	text := `Jdsfsdfdsfsdfdsfdsfsdfsкегщшуегущшегукшцукгнуцгшкуцгкшщуцгкшуцгшщкгуцщкшгуцкгуцшщкгуцшкгуцшщкгцущшкгуцшщгк
	цушщкввгвыкшгкшгыукгшщуцшгщцушгщкуцшщгкцугшщкуцшгщкуцшгщкшгуцкшгщуцщкгшуцгшщкугцшщкшгуцщкшщгуцщкгшуцшгщкуцгшщку
	гцшщкшгщуцкгшщуцщгшкуцшгщкугцшщкшгущцшкгщJdsfsdfdsfsdfdsfdsfsdfsкегщшуегущшегукшцукгнуцгшкуцгкшщуцгкшуцгшщкгуцщкшгуцкгуцшщкгуцшкгуцшщкгцущшкгуцшщгк
	цушщкввгвыкшгкшгыукгшщуцшгщцушгщкуцшщгкцугшщкуцшгщкуцшгщкшгуцкшгщуцщкгшуцгшщкугцшщкшгуцщкшщгуцщкгшуцшгщкуцгшщку
	гцшщкшгщуцкгшщуцщгшкуцшгщкугцшщкшгущцшкгщJdsfsdfdsfsdfdsfdsfsdfsкегщшуегущшегукшцукгнуцгшкуцгкшщуцгкшуцгшщкгуцщкшгуцкгуцшщкгуцшкгуцшщкгцущшкгуцшщгк
	цушщкввгвыкшгкшгыукгшщуцшгщцушгщкуцшщгкцугшщкуцшгщкуцшгщкшгуцкшгщуцщкгшуцгшщкугцшщкшгуцщкшщгуцщкгшуцшгщкуцгшщку
	гцшщкшгщуцкгшщуцщгшкуцшгщкугцшщкшгущцшкгщJdsfsdfdsfsdfdsfdsfsdfsкегщшуегущшегукшцукгнуцгшкуцгкшщуцгкшуцгшщкгуцщкшгуцкгуцшщкгуцшкгуцшщкгцущшкгуцшщгк
	цушщкввгвыкшгкшгыукгшщуцшгщцушгщкуцшщгкцугшщкуцшгщкуцшгщкшгуцкшгщуцщкгшуцгшщкугцшщкшгуцщкшщгуцщкгшуцшгщкуцгшщку
	гцшщкшгщуцкгшщуцщгшкуцшгщкугцшщкшгущцшкгщ`
	if validateMessage(text) {
		t.Errorf("Длина сообщения меньше 1024 символов")
	}
}

func TestValidatePasswordLess(t *testing.T) {
	password := "Jdsfsdfdsfsdfdsfdsfsdfs"
	if !validatePassword(password) {
		t.Errorf("Длина пароля больше 32 символов")
	}
}

func TestValidatePasswordMore(t *testing.T) {
	password := "Jdsfsdfdsfsdfdsfdsfsdfsrertyerutuyertyeruitrey"
	if validatePassword(password) {
		t.Errorf("Длина меньше 32 символов")
	}
}

func TestCheckUserPassword(t *testing.T) {
	usersstorage := NewUserStorageMemory(new(PasswordHasherSha1))
	usersstorage.Create("Андрей", "fghfghfghfgh")
	userName := "Андрей"
	password := "fghfghfghfgh"
	if !checkUserPassword(userName, password, usersstorage) {
		t.Errorf("Не корректный пароль")
	}
}

func TestCheckUserPasswordEmpty(t *testing.T) {
	usersstorage := NewUserStorageMemory(new(PasswordHasherSha1))
	userName := "Андрей"
	password := "fghfghfghfgh"
	if checkUserPassword(userName, password, usersstorage) {
		t.Errorf("Пароль корректный")
	}
}
