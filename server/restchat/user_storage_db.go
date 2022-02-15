package restchat

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

type UserStorageDB struct {
	Users          []UserModel
	passwordHasher IPasswordHasher
	connect        *pgx.Conn
}

func NewUserStorageDB() *UserStorageDB {
	msdb := new(UserStorageDB)
	msdb.connect = connectDB()
	return msdb
}

type IUserStorageDB interface {
	CreateUser(username string, password string) (UserModel, error)
	GetByName(name string) (UserModel, error)
	GetById(id uint) (UserModel, error)
	GetByIds(ids []uint) ([]UserModel, error)
}

func (userStorageDB *UserStorageDB) CreateUser(username string, password string) (UserModel, error) {
	var text string
	var userId int
	query := `select text,user_id from "UserModel".messages u where id=$1`
	err := userStorageDB.connect.QueryRow(context.Background(), query, 1).Scan(&text, &userId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	defer userStorageDB.connect.Close(context.Background())
	fmt.Println(text, userId)
	return UserModel{ID: 1, Username: "username", PasswordHash: "password"}, fmt.Errorf(text, userId)
}
