package authorizer

type AcceptingAuthorizerVerificationMock struct {
	AuthorizeWasCalled bool
}

func (mock *AcceptingAuthorizerVerificationMock) Authorize(username, password string) (bool, error) {
	AuthorizeWasCalled = true
	return true, nil
}

func (mock *AcceptingAuthorizerVerificationMock) Verify() bool {
	return mock.AuthorizeWasCalled
}
