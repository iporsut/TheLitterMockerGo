package authorizer

type AcceptingAuthorizerFake struct {
}

func (fake *AcceptingAuthorizerFake) Authorize(username, password string) (bool, error) {
	return (username == "Bob"), nil
}
