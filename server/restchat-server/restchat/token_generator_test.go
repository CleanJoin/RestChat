package restchat

import (
	"testing"
)

func TestCreateTokenGenerator(t *testing.T) {
	var itg = new(UuidSession)
	outuuid := itg.Create()
	if outuuid == "" {
		t.Errorf("Ошибка не верный вывод из функции")

	}

}
