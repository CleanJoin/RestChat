package restchat

import (
	"testing"
)

func TestCalculateHash(t *testing.T) {
	phs := new(PasswordHasherSha1)
	password := "hello"
	outshas1 := phs.CalculateHash(password)
	request := "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"
	if outshas1 != request {
		t.Errorf("Ошибка не верный вывод из функции")
	}

}
func TestIPasswordHasher(t *testing.T) {
	inter := new(PasswordHasherSha1)
	password := "hello"
	inter.CalculateHash(password)
}
