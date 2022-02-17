package restchat

import (
	"fmt"
	"testing"
)

func TestUserCreateDB(t *testing.T) {
	var passwordHash IPasswordHasher = new(PasswordHasherSha1)
	var inter IUserStorage = NewUserStorageDB(passwordHash)
	userModel, err := inter.Create("Andrey", "qweasd123")
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(userModel)
}

func TestGetByIdDB(t *testing.T) {
	var passwordHash IPasswordHasher = new(PasswordHasherSha1)
	var inter IUserStorage = NewUserStorageDB(passwordHash)
	userModel, err := inter.GetById(1)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(userModel)
}

func TestGetByIdsDB(t *testing.T) {
	var passwordHash IPasswordHasher = new(PasswordHasherSha1)
	var inter IUserStorage = NewUserStorageDB(passwordHash)
	var ids = []uint{1, 2, 3}
	userModel, err := inter.GetByIds(ids)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(userModel)
}

func TestGetByNameDB(t *testing.T) {
	var passwordHash IPasswordHasher = new(PasswordHasherSha1)
	var inter IUserStorage = NewUserStorageDB(passwordHash)
	userModel, err := inter.GetByName("Andrey")
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(userModel)
}
