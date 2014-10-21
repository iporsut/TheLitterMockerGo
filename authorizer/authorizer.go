package authorizer

type Authorizer interface {
	Authorize(username, password string) (bool, error)
}
