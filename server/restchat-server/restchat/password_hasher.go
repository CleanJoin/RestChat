package restchat

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

type PasswordHasherSha1 struct {
}
type IPasswordHasher interface {
	CalculateHash(password string) string
}

func (*PasswordHasherSha1) CalculateHash(password string) (string, error) {
	if password == "" {
		return "", fmt.Errorf("%s", "Пароль пустой")
	}
	sha1 := sha1.New()
	sha1.Write([]byte(password))
	sha1_hash := hex.EncodeToString(sha1.Sum(nil))
	return sha1_hash, fmt.Errorf("хэш создан: %s", sha1_hash)
}
