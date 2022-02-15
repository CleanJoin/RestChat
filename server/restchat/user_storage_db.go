package postgresDB

import (
	"context"
	"fmt"
	"os"
	"restchat-server/restchat"

	"github.com/jackc/pgx/v4"
)

type UserStorageDB struct {
	Users          []restchat.UserModel
	passwordHasher restchat.IPasswordHasher
	connect        *pgx.Conn
}

func NewUserStorageDB() *UserStorageDB {
	msdb := new(UserStorageDB)
	msdb.connect = connectDB()
	return msdb
}

type IUserStorage interface {
	CreateUser(username string, password string) (restchat.UserModel, error)
	GetByName(name string) (restchat.UserModel, error)
	GetById(id uint) (restchat.UserModel, error)
	GetByIds(ids []uint) ([]restchat.UserModel, error)
}

func (userStorageDB *UserStorageDB) CreateUser(username string, password string) (restchat.UserModel, error) {
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
	return restchat.UserModel{ID: 1, Username: "username", PasswordHash: "password"}, fmt.Errorf(text, userId)
}
