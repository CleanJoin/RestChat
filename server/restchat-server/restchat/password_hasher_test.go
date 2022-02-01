package restchat

import (
	"fmt"
	"testing"
)

func TestCalculateHash(t *testing.T) {
	phs := new(PasswordHasherSha1)
	password := "hello"
	outshas1, err := phs.CalculateHash(password)
	request := "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d"
	if outshas1 != request {
		t.Errorf("Ошибка не верный вывод из функции")

	}
	fmt.Println(err)
}
func TestCalculateHashEmpty(t *testing.T) {
	phs := new(PasswordHasherSha1)
	password := ""
	outshas1, err := phs.CalculateHash(password)
	request := ""
	if outshas1 != request {
		t.Errorf("%s", "Пароль не пустой, нужно кодировать")

	}
	fmt.Println(err)
}
