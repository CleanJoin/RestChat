package restchat

import (
	"fmt"
	"sort"
)

type UserStorageMemory struct {
	Users          []UserModel
	passwordHasher IPasswordHasher
	nextId uint
}

type IUserStorage interface {
	Create(username string, password string) (UserModel, error)
	GetByName(name string) (UserModel, error)
	GetById(id uint) (UserModel, error)
	GetByIds(ids []uint) ([]UserModel, error)
}

func NewUserStorageMemory(passwordHasher IPasswordHasher) *UserStorageMemory {
	ssm := new(UserStorageMemory)
	ssm.passwordHasher = passwordHasher
	ssm.nextId = 0
	return ssm
}

func (userStorage *UserStorageMemory) Create(username string, password string) (UserModel, error) {
	lenlastuser := len(userStorage.Users)
	passwordHash := userStorage.passwordHasher.CalculateHash(password)
	userStorage.Users = append(userStorage.Users, UserModel{ID: userStorage.nextId, Username: username, PasswordHash: passwordHash})
	if lenlastuser >= len(userStorage.Users) {
		return UserModel{}, fmt.Errorf("%s", "Не удалось добавить сообщение")
	}
	userStorage.nextId++
	return userStorage.Users[len(userStorage.Users)-1], nil
}

func (userStorage *UserStorageMemory) GetByName(username string) (UserModel, error) {
	for i, r := range userStorage.Users {
		if r.Username == username {
			return userStorage.Users[i], nil
		}
	}
	return UserModel{}, fmt.Errorf("не нашелся пользователь по указанному Username: %v", username)
}

func (userStorage *UserStorageMemory) GetById(id uint) (UserModel, error) {
	for i, r := range userStorage.Users {
		if r.ID == id {
			return userStorage.Users[i], nil
		}
	}
	return UserModel{}, fmt.Errorf("пользователь с указанным id:%v не нашелся ", id)
}

func (userStorage *UserStorageMemory) GetByIds(ids []uint) ([]UserModel, error) {
	idIndexer := make(map[uint]uint)
	newUserStorage := new(UserStorageMemory)
	for index, user := range userStorage.Users {
		idIndexer[user.ID] = uint(index)
	}
	for _, id := range ids {
		if val, ok := idIndexer[id]; ok {
			newUserStorage.Users = append(newUserStorage.Users, UserModel{ID: id, Username: userStorage.Users[val].Username, PasswordHash: userStorage.Users[val].PasswordHash})
		}

	}
	return newUserStorage.Users, nil
}
