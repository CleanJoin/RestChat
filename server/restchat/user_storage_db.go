package restchat

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type UserStorageDB struct {
	Users          []UserModel
	passwordHasher IPasswordHasher
	connect        *pgxpool.Pool
}

func NewUserStorageDB(passwordHasher IPasswordHasher) *UserStorageDB {
	msdb := new(UserStorageDB)
	msdb.connect = connectDB()
	msdb.passwordHasher = passwordHasher
	return msdb
}

func (userStorageDB *UserStorageDB) Create(username string, password string) (UserModel, error) {
	password = userStorageDB.passwordHasher.CalculateHash(password)
	query := `INSERT INTO "UserModel".users (username,passwordhash) VALUES($1, $2)`
	commandTag, err := userStorageDB.connect.Exec(context.Background(), query, username, password)
	if err != nil {
		fmt.Println(err)
	}
	commandTag.Insert()
	dd, _ := userStorageDB.connect.Query(context.Background(), query, username, password)
	dddd, _ := dd.Values()
	fmt.Println(dddd)
	return userStorageDB.Users[len(userStorageDB.Users)-1], nil
}

func (userStorageDB *UserStorageDB) GetByName(name string) (UserModel, error) {
	userModel := new(UserModel)
	query := `select * from "UserModel".users u where username =$1`
	commandTag := userStorageDB.connect.QueryRow(context.Background(), query, name)
	commandTag.Scan(userModel.Username, userModel.PasswordHash, userModel.ID)
	return UserModel{userModel.ID, userModel.Username, userModel.PasswordHash}, fmt.Errorf("не нашелся пользователь по указанному Username: %v", name)
}

func (userStorageDB *UserStorageDB) GetById(id uint) (UserModel, error) {
	userModel := new(UserModel)
	getUserByRequest(userStorageDB, userModel, id)
	return UserModel{userModel.ID, userModel.Username, userModel.PasswordHash}, fmt.Errorf("пользователь с указанным id:%v не нашелся ", id)
}

func (userStorageDB *UserStorageDB) GetByIds(ids []uint) ([]UserModel, error) {
	newUserStorage := new(UserStorageDB)
	userModel := new(UserModel)
	for _, id := range ids {
		getUserByRequest(userStorageDB, userModel, id)
		newUserStorage.Users = append(newUserStorage.Users, UserModel{ID: id, Username: userModel.Username, PasswordHash: userModel.PasswordHash})
	}
	return newUserStorage.Users, nil
}

func getUserByRequest(userStorageDB *UserStorageDB, userModel *UserModel, id uint) {
	query := `select * from "UserModel".users u where id =$1`
	commandTag := userStorageDB.connect.QueryRow(context.Background(), query, id)
	commandTag.Scan(userModel.Username, userModel.PasswordHash, userModel.ID)
}
