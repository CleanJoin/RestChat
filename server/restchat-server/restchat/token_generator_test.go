package restchat

import (
	"testing"
)

func TestCreateTokenGenerator(t *testing.T) {
	var itg = new(TokenGeneratorUUID)
	uuid := itg.Create()
	if uuid == "" {
		t.Errorf("Ошибка не верный вывод из функции")
	}
}
func TestITokenGenerator(t *testing.T) {
	inter := new(TokenGeneratorUUID)
	inter.Create()
}
