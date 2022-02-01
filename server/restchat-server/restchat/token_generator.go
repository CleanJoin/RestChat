package restchat

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type UuidSession struct {
	UUID string
}
type ITokenGenerator interface {
	CreateUuid() string
}

func (*UuidSession) CreateUuid() (string, error) {
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return uuid, fmt.Errorf("uuid создан: %s", uuid)
}
