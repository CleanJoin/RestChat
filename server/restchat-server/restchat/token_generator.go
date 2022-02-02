package restchat

import (
	"strings"

	"github.com/google/uuid"
)

type TokenGeneratorUUID struct {
}
type ITokenGenerator interface {
	Create() string
}

func (*TokenGeneratorUUID) Create() string {
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return uuid
}
