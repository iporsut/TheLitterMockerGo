package authorizer

type AcceptingAuthorizerStub struct {
}

func (stub *AcceptingAuthorizerStub) Authorize(username, password string) (bool, error) {
	return true, nil
}
