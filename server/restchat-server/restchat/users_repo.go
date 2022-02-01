package restchat

import "fmt"

type UsersMemRepo struct {
	Users []UserModel
}

type IUsersRepo interface {
	CreateUser(username string, password string) UserModel
	GetByName(name string) UserModel
	GetById(id int) UserModel
	GetbyIds(ids []int) []UserModel
}

func NewUsersMemRepo() *UsersMemRepo {
	return &UsersMemRepo{}
}

func (umr *UsersMemRepo) CreateUser(username string, password string) (UserModel, error) {
	return UserModel{ID: 0, Username: "", PasswordHash: ""}, fmt.Errorf("%s", "Не нашелся пользователь:")
}
func (umr *UsersMemRepo) GetByName(name string) (UserModel, error) {
	for i, r := range umr.Users {
		if r.Username == name {
			return umr.Users[i], fmt.Errorf("нашелся пользователь: %v", umr.Users[i])
		}
	}
	return UserModel{ID: 0, Username: "", PasswordHash: ""}, fmt.Errorf("%s", "Не нашелся пользователь:")
}

func (umr *UsersMemRepo) GetById(id int) (UserModel, error) {
	for i, r := range umr.Users {
		if r.ID == id {
			return umr.Users[i], fmt.Errorf("нашелся пользователь: %v", umr.Users[i])
		}
	}
	return UserModel{ID: 0, Username: "", PasswordHash: ""}, fmt.Errorf("%s", "Не нашелся пользователь:")
}

// func (umr *UsersMemRepo) GetById(ids []int) ([]UserModel, error) {

// }
