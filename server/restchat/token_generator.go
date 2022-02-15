package restchat

import (
	"github.com/google/uuid"
)

type TokenGeneratorUUID struct {
}
type ITokenGenerator interface {
	Create() string
}

func (*TokenGeneratorUUID) Create() string {
	return uuid.New().String()
}
