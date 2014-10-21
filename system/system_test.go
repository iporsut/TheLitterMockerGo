package system

import (
	"lab/mock/authorizer"
	"testing"
)

func TestNewlyCreatedSystemHasNoLoggedInUsers(t *testing.T) {
	system := New(authorizer.NewDummy())

	loginCount := system.LoginCount()
	expectedLoginCount := 0

	if loginCount != expectedLoginCount {
		t.Errorf("Expect Login Count as %d but %d", expectedLoginCount, loginCount)
	}
}
