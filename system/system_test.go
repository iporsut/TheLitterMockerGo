package system

import (
	"lab/mock/authorizer"
	"testing"
)

func TestNewlyCreatedSystemHasNoLoggedInUsers(t *testing.T) {
	system := &System{&authorizer.DummyAuthorizer{}}

	loginCount := system.LoginCount()
	expectedLoginCount := 0

	if loginCount != expectedLoginCount {
		t.Errorf("Expect Login Count as %d but %d", expectedLoginCount, loginCount)
	}
}
