package v1

import (
	"path"

	"github.com/cloudfoundry/bosh-cli/director/template"
	vaultapi "github.com/hashicorp/vault/api"
)

// A SecretReader reads a vault secret from the given path. It should
// be thread safe!
type SecretReader interface {
	Read(path string) (*vaultapi.Secret, error)
}

type V1 struct {
	SecretReader SecretReader
	PathPrefix   string
}

func (v V1) Get(varDef template.VariableDefinition) (interface{}, bool, error) {
	var secret *vaultapi.Secret
	var found bool
	var err error

	secret, found, err = v.findSecret(v.path(varDef.Name))
	if err != nil {
		return nil, false, err
	}

	if !found {
		return nil, false, nil
	}

	val, found := secret.Data["value"]
	if found {
		return val, true, nil
	}

	return nil, false, nil
}

func (v V1) findSecret(path string) (*vaultapi.Secret, bool, error) {
	secret, err := v.SecretReader.Read(path)
	if err != nil {
		return nil, false, err
	}

	if secret != nil {
		return secret, true, nil
	}

	return nil, false, nil
}

func (v V1) path(segments ...string) string {
	return path.Join(append([]string{v.PathPrefix}, segments...)...)
}

func (v V1) List() ([]template.VariableDefinition, error) {
	// Don't think this works with vault.. if we need it to we'll figure it out
	// var defs []template.VariableDefinition

	// secret, err := v.vaultClient.List(v.PathPrefix)
	// if err != nil {
	// 	return defs, err
	// }

	// var def template.VariableDefinition
	// for name, _ := range secret.Data {
	// 	defs := append(defs, template.VariableDefinition{
	// 		Name: name,
	// 	})
	// }

	return []template.VariableDefinition{}, nil
}
