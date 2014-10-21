package authorizer

import (
	"errors"
)

type DummyAuthorizer struct {
}

func (dummy *DummyAuthorizer) Authorize(username, password string) (bool, error) {
	return false, errors.New("Don't used Dummy")
}

func NewDummy() *DummyAuthorizer {
	return &DummyAuthorizer{}
}
