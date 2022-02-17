package restchat

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lib/pq"
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
	var id uint
	newPasswordHash := userStorageDB.passwordHasher.CalculateHash(password)
	query := `INSERT INTO "UserModel".users (username,"password") VALUES($1, $2) RETURNING id;`
	row := userStorageDB.connect.QueryRow(context.Background(), query, username, newPasswordHash)
	err := row.Scan(&id)
	if err != nil {
		return UserModel{}, fmt.Errorf(err.Error())
	}
	return UserModel{id, username, newPasswordHash}, nil
}

func (userStorageDB *UserStorageDB) GetByName(name string) (UserModel, error) {
	userModel := new(UserModel)
	query := `select * from "UserModel".users u where username=$1`
	commandTag := userStorageDB.connect.QueryRow(context.Background(), query, name)
	err := commandTag.Scan(&userModel.ID, &userModel.Username, &userModel.PasswordHash)
	if err != nil {
		return *userModel, fmt.Errorf(err.Error())
	}
	return *userModel, nil
}

func (userStorageDB *UserStorageDB) GetById(id uint) (UserModel, error) {
	userModel := new(UserModel)
	query := `select * from "UserModel".users u where id=$1`
	commandTag := userStorageDB.connect.QueryRow(context.Background(), query, id)
	err := commandTag.Scan(&userModel.ID, &userModel.Username, &userModel.PasswordHash)
	if err != nil {
		return *userModel, fmt.Errorf(err.Error())
	}

	return *userModel, nil
}

func (userStorageDB *UserStorageDB) GetByIds(ids []uint) ([]UserModel, error) {
	userModel := new(UserModel)
	query := `select * from "UserModel".users u where id = ANY($1)`
	commandTag, err := userStorageDB.connect.Query(context.Background(), query, pq.Array(ids))
	if err != nil {
		return userStorageDB.Users, fmt.Errorf(err.Error())
	}
	for commandTag.Next() {
		err := commandTag.Scan(&userModel.ID, &userModel.Username, &userModel.PasswordHash)
		userStorageDB.Users = append(userStorageDB.Users, UserModel{userModel.ID, userModel.Username, userModel.PasswordHash})
		if err != nil {
			return userStorageDB.Users, fmt.Errorf(err.Error())
		}
	}

	return userStorageDB.Users, nil
}
