package authorizer

type AcceptingAuthorizerSpy struct {
	AuthorizeWasCalled bool
}

func (spy *AcceptingAuthorizerSpy) Authorize(username, password string) (bool, error) {
	spy.AuthorizeWasCalled = true
	return true, nil
}
