package requests

import (
	"github.com/corpix/uarand"
)

type IUserAgentGenerator interface {
	Generate() string
}
type UserAgentGenerator struct {
}

func NewUserAgentGenerator() *UserAgentGenerator {
	return &UserAgentGenerator{}
}

func (u *UserAgentGenerator) Generate() string {
	return uarand.GetRandom()
}
