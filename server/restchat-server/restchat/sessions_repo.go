package restchat

import (
	"fmt"
	"sort"
)

type UserSessionsMemRepo struct {
	Sessions       []SessionModel
	tokenGenerator ITokenGenerator
}

type IUserSessions interface {
	GetOnlineUserIds() ([]uint, error)
	DeleteSession(api_token string) error
	CreateSession(user_id uint) (SessionModel, error)
}

func NewUserSessionsMemRepo(tokenGenerator ITokenGenerator) *UserSessionsMemRepo {
	usmr := new(UserSessionsMemRepo)
	usmr.tokenGenerator = tokenGenerator

	return usmr
}

//
func (usmr *UserSessionsMemRepo) GetOnlineUserIds() ([]uint, error) {
	var d []uint
	return d, fmt.Errorf("Test")
}
func deleteSessionByIndex(sm []SessionModel, index int) []SessionModel {
	return append(sm[:index], sm[index+1:]...)
}

func getLastSessionId(usmr *UserSessionsMemRepo) uint {
	if usmr == nil || len(usmr.Sessions) == 0 {
		return 0
	}
	sort.Slice(usmr.Sessions, func(i, j int) (less bool) {
		return usmr.Sessions[i].ID > usmr.Sessions[j].ID
	})
	return usmr.Sessions[0].ID
}

func (usmr *UserSessionsMemRepo) DeleteSession(api_token string) (UserSessionsMemRepo, error) {
	if api_token == "" {
		return *usmr, fmt.Errorf("%s%v", "Не удалось удалить сессию, пустой токен", *usmr)
	}
	index := 0
	for i, r := range usmr.Sessions {
		if r.Auth_token == api_token {
			index = i
			break
		}
	}
	usmr.Sessions = deleteSessionByIndex(usmr.Sessions, index)
	return *usmr, fmt.Errorf("%s", "Не удалось удалить сессию, пустой токен")
}

func (usmr *UserSessionsMemRepo) Create(user_id uint) (SessionModel, error) {
	if user_id == 0 {
		return SessionModel{ID: 0, UserId: 0, Auth_token: ""}, fmt.Errorf("%s", "Не удалось создать сессию, так как user_id пустой")
	}
	id := getLastSessionId(usmr)
	id++
	lenlastmessages := len(usmr.Sessions)
	// uuid := new(UuidSession)
	// authtoken, err := uuid.Create()
	// fmt.Println(err)
	// usermemrep := new(UsersMemRepo)
	// usermodel, err := usermemrep.GetById(user_id)
	// fmt.Println(err)
	authToken := usmr.tokenGenerator.Create()
	usmr.Sessions = append(usmr.Sessions, SessionModel{ID: id, UserId: 0, Auth_token: authToken})

	if lenlastmessages >= len(usmr.Sessions) {
		return SessionModel{ID: 0, UserId: 0, Auth_token: ""}, fmt.Errorf("%s", "Не удалось добавить сообщение")
	}

	return usmr.Sessions[len(usmr.Sessions)-1], fmt.Errorf("сообщение создалось: %v", usmr.Sessions[len(usmr.Sessions)-1])
}
