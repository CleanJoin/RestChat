package restchat

import (
	"fmt"
	"testing"
)

func TestCreateTokenGenerator(t *testing.T) {
	var itg = new(TokenGeneratorUUID)
	uuid := itg.Create()
	if uuid == "" {
		t.Errorf("Ошибка не верный вывод из функции")
	}
	fmt.Printf("UUID Готов: %s", uuid)
}
