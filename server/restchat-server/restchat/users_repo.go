package restchat

import "fmt"

type UserStorageMemory struct {
	Users []UserModel
}

type IUserStorage interface {
	Create(username string, password string) UserModel
	GetByName(name string) UserModel
	GetById(id uint) UserModel
	GetbyIds(ids []uint) []UserModel
}

func NewUserStorageMemory() *UserStorageMemory {
	return &UserStorageMemory{}
}

func (usm *UserStorageMemory) Create(username string, password string) (UserModel, error) {
	return UserModel{ID: 0, Username: "", PasswordHash: ""}, fmt.Errorf("%s", "Не нашелся пользователь:")
}
func (usm *UserStorageMemory) GetByName(name string) (UserModel, error) {
	for i, r := range usm.Users {
		if r.Username == name {
			return usm.Users[i], fmt.Errorf("нашелся пользователь: %v", usm.Users[i])
		}
	}
	return UserModel{ID: 0, Username: "", PasswordHash: ""}, fmt.Errorf("%s", "Не нашелся пользователь:")
}

func (usm *UserStorageMemory) GetById(id uint) (UserModel, error) {
	for i, r := range usm.Users {
		if r.ID == id {
			return usm.Users[i], fmt.Errorf("нашелся пользователь: %v", usm.Users[i])
		}
	}
	return UserModel{ID: 0, Username: "", PasswordHash: ""}, fmt.Errorf("%s", "Не нашелся пользователь:")
}

// func (usm *UserStorageMemory) GetByIds(ids []uint) ([]UserModel, error) {
// 	return  fmt.Errorf("%s", "Не нашелся пользователь:")
// }
