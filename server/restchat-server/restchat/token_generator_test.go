package restchat

import (
	"fmt"
	"testing"
)

func TestCreateUuid(t *testing.T) {
	var itg = new(UuidSession)
	outuuid, err := itg.CreateUuid()
	if outuuid == "" {
		t.Errorf("Ошибка не верный вывод из функции")

	}
	fmt.Println(err)
}
