package db

import "testing"

func TestUserCreateDB(t *testing.T) {
	UserStorageDB := NewUserStorageDB()
	UserStorageDB.CreateUser("GHhhb", "dsfdsfs")
}