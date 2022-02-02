package restchat

import (
	"strings"

	"github.com/google/uuid"
)

type UuidSession struct {
}
type ITokenGenerator interface {
	Create() string
}

func (*UuidSession) Create() string {
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return uuid
}
