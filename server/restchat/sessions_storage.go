package restchat

import (
	"fmt"
	"sort"
)

type SessionStorageMemory struct {
	Sessions       []SessionModel
	tokenGenerator ITokenGenerator
}

type ISessionStorage interface {
	GetOnlineUserIds() ([]uint, error)
	Delete(apiToken string) error
	Create(userId uint) (SessionModel, error)
	GetUserId(apiToken string) (uint, error)
}

func NewSessionStorageMemory(tokenGenerator ITokenGenerator) *SessionStorageMemory {
	sessionStorage := new(SessionStorageMemory)
	sessionStorage.tokenGenerator = tokenGenerator

	return sessionStorage
}

func (sessionStorage *SessionStorageMemory) GetOnlineUserIds() ([]uint, error) {
	if sessionStorage == nil || len(sessionStorage.Sessions) == 0 {
		return nil, fmt.Errorf("пользователи не в сети")
	}
	var onlineUsers []uint
	for _, r := range sessionStorage.Sessions {
		onlineUsers = append(onlineUsers, r.UserId)
	}
	return onlineUsers, nil
}

func (sessionStorage *SessionStorageMemory) GetUserId(apiToken string) (uint, error) {
	if sessionStorage == nil || len(sessionStorage.Sessions) == 0 {
		return 0, fmt.Errorf("пользователи не в сети")
	}

	for _, ss := range sessionStorage.Sessions {
		if apiToken == ss.ApiToken {
			return ss.UserId, nil
		}
	}
	return 0, fmt.Errorf("по токену нет такого пользователя")
}

func deleteSessionByIndex(sm []SessionModel, index int) []SessionModel {
	return append(sm[:index], sm[index+1:]...)
}

func getLastSessionId(sessionStorage *SessionStorageMemory) uint {
	if sessionStorage == nil || len(sessionStorage.Sessions) == 0 {
		return 0
	}
	sort.Slice(sessionStorage.Sessions, func(i, j int) (less bool) {
		return sessionStorage.Sessions[i].ID > sessionStorage.Sessions[j].ID
	})
	return sessionStorage.Sessions[0].ID
}

func (sessionStorage *SessionStorageMemory) Delete(apiToken string) error {
	index := 0
	status := false
	for i, r := range sessionStorage.Sessions {
		if r.ApiToken == apiToken {
			index = i
			status = true
			break
		}
	}
	if !status {
		return fmt.Errorf("не удалось удалить сессию пользователя")
	}
	sessionStorage.Sessions = deleteSessionByIndex(sessionStorage.Sessions, index)
	return nil
}

func (sessionStorage *SessionStorageMemory) Create(userId uint) (SessionModel, error) {

	ApiToken := sessionStorage.tokenGenerator.Create()
	for id, session := range sessionStorage.Sessions {
		if session.UserId == userId && len(sessionStorage.Sessions) != 1 {
			sessionStorage.Sessions = append(sessionStorage.Sessions[:id], sessionStorage.Sessions[id+1:]...)
		}
		if session.UserId == userId && len(sessionStorage.Sessions) == 1 {
			sessionStorage.Sessions = []SessionModel{}
			break
		}

	}
	lenCurrentMessages := len(sessionStorage.Sessions)
	sessionId := getLastSessionId(sessionStorage)
	sessionId++
	sessionStorage.Sessions = append(sessionStorage.Sessions, SessionModel{ID: sessionId, UserId: userId, ApiToken: ApiToken})
	if lenCurrentMessages >= len(sessionStorage.Sessions) {
		return SessionModel{}, fmt.Errorf("не удалось добавить сессию")
	}
	return sessionStorage.Sessions[len(sessionStorage.Sessions)-1], nil
}
