package restchat

import (
	"reflect"
	"strconv"
	"testing"
)

func TestIUserStorage(t *testing.T) {
	inter := NewUserStorageMemory(new(PasswordHasherSha1))

	user, err := inter.Create("Андрей", "Андрей")
	if err != nil {
		t.Errorf("Не создался пользователь %v", err)
	}
	_, err = inter.GetById(user.ID)
	if err != nil {
		t.Errorf("Не получили пользователя по id %v", err)
	}
	_, err = inter.GetByName(user.Username)
	if err != nil {
		t.Errorf("Не получили пользователя по username %v", err)
	}
	ids := []uint{1, 2, 3, 4, 4}
	_, err = inter.GetByIds(ids)
	if err != nil {
		t.Errorf("Не получили пользователя по указанным идентификаторам %v", err)
	}

}
func TestCreateUser(t *testing.T) {
	userStorage := NewUserStorageMemory(new(PasswordHasherSha1))
	_, err := userStorage.Create("Андрей", "Андрей")
	if err != nil {
		t.Errorf("Не создался пользователь %v", err)
	}
}
func TestGetLastUserId(t *testing.T) {
	userStorage := NewUserStorageMemory(new(PasswordHasherSha1))
	_, err := userStorage.Create("Андрей", "Андрей")
	if err != nil {
		t.Errorf("Не создался пользователь %v", err)
	}
	if getLastUserId(userStorage) != userStorage.Users[0].ID {
		t.Errorf("Не получили идентификатор последнего пользователя")
	}
}

func TestGetByName(t *testing.T) {
	userStorage := NewUserStorageMemory(new(PasswordHasherSha1))
	user, err := userStorage.Create("Андрей", "Андрей")
	if err != nil {
		t.Errorf("Не создался пользователь %v", err)
	}
	_, err = userStorage.GetByName(user.Username)
	if err != nil {
		t.Errorf("не нашелся пользователь по указанному Username %v", err)
	}
}
func TestGetById(t *testing.T) {
	userStorage := NewUserStorageMemory(new(PasswordHasherSha1))
	user, err := userStorage.Create("Андрей", "Андрей")
	if err != nil {
		t.Errorf("Не создался пользователь %v", err)
	}
	_, err = userStorage.GetById(user.ID)
	if err != nil {
		t.Errorf("пользователь с указанным id не нашелся %v", err)
	}
}
func TestGetByIds(t *testing.T) {
	userStorage := NewUserStorageMemory(new(PasswordHasherSha1))
	for i := 1; i < 10; i++ {
		userStorage.Create("Андрей"+strconv.Itoa(i), "Андрей"+strconv.Itoa(i))
	}
	ids := []uint{2, 3, 6, 7}
	request, err := userStorage.GetByIds(ids)
	if err != nil {
		t.Errorf("не нашелся пользователь по указанным идентификаторам %v", err)
	}
	outUserModel := []UserModel{{ID: 2, Username: "Андрей2", PasswordHash: request[0].PasswordHash}, {ID: 3, Username: "Андрей3", PasswordHash: request[1].PasswordHash}, {ID: 6, Username: "Андрей6", PasswordHash: request[2].PasswordHash}, {ID: 7, Username: "Андрей7", PasswordHash: request[3].PasswordHash}}
	if reflect.DeepEqual(request, outUserModel) != true {
		t.Errorf("Не верный список пользователей %v %v", request, outUserModel)
	}
}
