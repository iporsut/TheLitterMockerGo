package system

import (
	"lab/mock/authorizer"
)

type System struct {
	Authorizer authorizer.Authorizer
}

func (system *System) LoginCount() int {
	return 0
}

func New(a authorizer.Authorizer) *System {
	return &System{
		Authorizer: a,
	}
}
