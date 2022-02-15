package restchat

import (
	"crypto/sha1"
	"encoding/hex"
)

type PasswordHasherSha1 struct {
}
type IPasswordHasher interface {
	CalculateHash(password string) string
}

func (*PasswordHasherSha1) CalculateHash(password string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(password))
	sha1_hash := hex.EncodeToString(sha1.Sum(nil))
	return sha1_hash
}
