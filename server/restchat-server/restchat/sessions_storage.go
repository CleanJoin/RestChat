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
}

func NewSessionStorageMemory(tokenGenerator ITokenGenerator) *SessionStorageMemory {
	ssm := new(SessionStorageMemory)
	ssm.tokenGenerator = tokenGenerator

	return ssm
}

func (ssm *SessionStorageMemory) GetOnlineUserIds() ([]uint, error) {
	if ssm == nil || len(ssm.Sessions) == 0 {
		return nil, fmt.Errorf("все пользователи не в сети")
	}
	var onlineUsers []uint
	for _, r := range ssm.Sessions {
		onlineUsers = append(onlineUsers, r.UserId)
	}
	return onlineUsers, fmt.Errorf("список пользователей онлайн:%v", onlineUsers)
}
func deleteSessionByIndex(sm []SessionModel, index int) []SessionModel {
	return append(sm[:index], sm[index+1:]...)
}

func getLastSessionId(ssm *SessionStorageMemory) uint {
	if ssm == nil || len(ssm.Sessions) == 0 {
		return 0
	}
	sort.Slice(ssm.Sessions, func(i, j int) (less bool) {
		return ssm.Sessions[i].ID > ssm.Sessions[j].ID
	})
	return ssm.Sessions[0].ID
}

func (ssm *SessionStorageMemory) Delete(apiToken string) error {
	index := 0
	for i, r := range ssm.Sessions {
		if r.AuthToken == apiToken {
			index = i
			break
		}
	}
	ssm.Sessions = deleteSessionByIndex(ssm.Sessions, index)
	return fmt.Errorf("удалили сессию с токеном %s", apiToken)
}

func (ssm *SessionStorageMemory) Create(userId uint) (SessionModel, error) {
	sessionId := getLastSessionId(ssm)
	sessionId++
	lenCurrentMessages := len(ssm.Sessions)
	authToken := ssm.tokenGenerator.Create()
	fmt.Println(authToken)
	ssm.Sessions = append(ssm.Sessions, SessionModel{ID: sessionId, UserId: userId, AuthToken: authToken})

	if lenCurrentMessages >= len(ssm.Sessions) {
		return SessionModel{ID: 0, UserId: 0, AuthToken: ""}, fmt.Errorf("не удалось добавить сообщение")
	}

	return ssm.Sessions[len(ssm.Sessions)-1], fmt.Errorf("сообщение создалось: %v", ssm.Sessions[len(ssm.Sessions)-1])
}
