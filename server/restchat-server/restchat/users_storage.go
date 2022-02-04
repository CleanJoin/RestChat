package restchat

import (
	"fmt"
	"sort"
)

type UserStorageMemory struct {
	Users          []UserModel
	passwordHasher IPasswordHasher
}

type IUserStorage interface {
	Create(username string, password string) UserModel
	GetByName(name string) UserModel
	GetById(id uint) UserModel
	GetByIds(ids []uint) []UserModel
}

func NewUserStorageMemory(passwordHasher IPasswordHasher) *UserStorageMemory {
	ssm := new(UserStorageMemory)
	ssm.passwordHasher = passwordHasher
	return ssm
}
func getLastUserId(usm *UserStorageMemory) uint {
	if usm == nil || len(usm.Users) == 0 {
		return 0
	}
	sort.Slice(usm.Users, func(i, j int) (less bool) {
		return usm.Users[i].ID > usm.Users[j].ID
	})
	return usm.Users[0].ID
}

func (usm *UserStorageMemory) Create(username string, password string) (UserModel, error) {
	id := getLastUserId(usm)
	id++
	lenlastuser := len(usm.Users)
	passwordHash := usm.passwordHasher.CalculateHash(password)
	usm.Users = append(usm.Users, UserModel{ID: id, Username: username, PasswordHash: passwordHash})
	if lenlastuser >= len(usm.Users) {
		return UserModel{ID: 0, Username: "", PasswordHash: ""}, fmt.Errorf("%s", "Не удалось добавить сообщение")
	}
	return usm.Users[len(usm.Users)-1], nil
}

func (usm *UserStorageMemory) GetByName(username string) (UserModel, error) {
	for i, r := range usm.Users {
		if r.Username == username {
			return usm.Users[i], nil
		}
	}
	return UserModel{ID: 0, Username: "", PasswordHash: ""}, fmt.Errorf("не нашелся пользователь по указанному Username: %v", username)
}

func (usm *UserStorageMemory) GetById(id uint) (UserModel, error) {
	for i, r := range usm.Users {
		if r.ID == id {
			return usm.Users[i], nil
		}
	}
	return UserModel{ID: 0, Username: "", PasswordHash: ""}, fmt.Errorf("пользователь с указанным id:%v не нашелся ", id)
}

func (usm *UserStorageMemory) GetByIds(ids []uint) ([]UserModel, error) {
	idIndexer := make(map[uint]uint)
	newUserStorage := new(UserStorageMemory)
	for index, user := range usm.Users {
		idIndexer[user.ID] = uint(index)
	}
	for _, id := range ids {
		if val, ok := idIndexer[id]; ok {
			newUserStorage.Users = append(newUserStorage.Users, UserModel{ID: id, Username: usm.Users[val].Username, PasswordHash: usm.Users[val].PasswordHash})
		}

	}
	return newUserStorage.Users, nil
}
