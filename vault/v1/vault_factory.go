package v1

import "github.com/cappyzawa/cf-push-with-vault/vault"

// The Factory will return a vault implementation of creds.Variables.
type Factory struct {
	sr     SecretReader
	prefix string
}

func NewFactory(sr SecretReader, prefix string) *Factory {
	factory := &Factory{
		sr:     sr,
		prefix: prefix,
	}

	return factory
}

// NewVariables will block until the loggedIn channel passed to the
// constructor signals a successful login.
func (factory *Factory) NewVariables() vault.Variables {
	return &V1{
		SecretReader: factory.sr,
		PathPrefix:   factory.prefix,
	}
}
