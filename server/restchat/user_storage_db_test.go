package restchat

import "testing"

func TestUserCreateDB(t *testing.T) {
	UserStorageDB := NewUserStorageDB(new(PasswordHasherSha1))
	UserStorageDB.Create("GHhhb", "dsfdsfs")
}
