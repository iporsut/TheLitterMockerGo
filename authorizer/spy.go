package authorizer

type AcceptingAuthorizerSpy struct {
	AuthorizeWasCalled bool
}

func (spy *AcceptingAuthorizerSpy) Authorize(username, password string) (bool, error) {
	AuthorizeWasCalled = true
	return true, nil
}
