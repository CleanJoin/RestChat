package restchat

import (
	"fmt"
	"sort"
)

type UserSessionsMemRepo struct {
	Sessions []SessionModel
}

type IUserSessions interface {
	GetOnlineUserIds() ([]int, error)
	DeleteSession(api_token string) error
	CreateSession(user_id int) (SessionModel, error)
}

func NewUserSessionsMemRepo() *UserSessionsMemRepo {
	return &UserSessionsMemRepo{}
}

func (usmr *UserSessionsMemRepo) GetOnlineUserIds() ([]int, error) {
	usermemrep := new(UsersMemRepo)
	var userid []int
	for _, r := range usmr.Sessions {

		usermodel, _ := usermemrep.GetByName(r.Username)
		userid = append(userid, usermodel.ID)

	}
	return userid, fmt.Errorf("вывод список пользователей онлайн: %v", userid)

}
func deleteSessionByIndex(sm []SessionModel, index int) []SessionModel {
	return append(sm[:index], sm[index+1:]...)
}

func getLastSessionId(usmr *UserSessionsMemRepo) int {
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
		return *usmr, fmt.Errorf("%s%v", "Не удалось удалить сесию, пустой токен", *usmr)
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

func (usmr *UserSessionsMemRepo) CreateSession(user_id int) (SessionModel, error) {
	if user_id == 0 {
		return SessionModel{ID: 0, Username: "", Auth_token: ""}, fmt.Errorf("%s", "Не удалось создать сесию, так как user_id пустой")
	}
	id := getLastSessionId(usmr)
	id++
	lenlastmessages := len(usmr.Sessions)
	uuid := new(UuidSession)
	authtoken, err := uuid.CreateUuid()
	fmt.Println(err)
	usermemrep := new(UsersMemRepo)
	usermodel, err := usermemrep.GetById(user_id)
	fmt.Println(err)
	usmr.Sessions = append(usmr.Sessions, SessionModel{ID: id, Username: usermodel.Username, Auth_token: authtoken})

	if lenlastmessages >= len(usmr.Sessions) {
		return SessionModel{ID: 0, Username: "", Auth_token: ""}, fmt.Errorf("%s", "Не удалось добавить сообщение")
	}

	return usmr.Sessions[len(usmr.Sessions)-1], fmt.Errorf("cообщение создалось: %v", usmr.Sessions[len(usmr.Sessions)-1])
}
