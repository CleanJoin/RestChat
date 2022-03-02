package restchat

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

func TestUserCreateDB(t *testing.T) {
	godotenv.Load(".env")
	connectDB := NewConnectDB(5432)
	var passwordHash IPasswordHasher = new(PasswordHasherSha1)
	var inter IUserStorage = NewUserStorageDB(passwordHash, connectDB)
	userModel, err := inter.Create("Andrey", "qweasd123")
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(userModel)
}

func TestGetByIdDBNotSearch(t *testing.T) {
	godotenv.Load(".env")

	connectDB := NewConnectDB(5432)
	var passwordHash IPasswordHasher = new(PasswordHasherSha1)
	var inter IUserStorage = NewUserStorageDB(passwordHash, connectDB)
	userModel, err := inter.GetById(200)
	if userModel.Username != "" {
		t.Errorf(err.Error())
	}

}

func TestGetByIdsDB(t *testing.T) {

	connectDB := NewConnectDB(5432)
	var passwordHash IPasswordHasher = new(PasswordHasherSha1)
	var inter IUserStorage = NewUserStorageDB(passwordHash, connectDB)
	var ids = []uint{1, 2, 3}
	userModel, err := inter.GetByIds(ids)
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(userModel)
}

func TestGetByNameDB(t *testing.T) {
	godotenv.Load(".env")
	connectDB := NewConnectDB(5432)
	var passwordHash IPasswordHasher = new(PasswordHasherSha1)
	var inter IUserStorage = NewUserStorageDB(passwordHash, connectDB)
	userModel, err := inter.GetByName("Andrey")
	if err != nil {
		t.Errorf(err.Error())
	}
	fmt.Println(userModel)
}
