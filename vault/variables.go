package vault

import "github.com/cloudfoundry/bosh-cli/director/template"

//go:generate counterfeiter . VariablesFactory

// VariablesFactory has intialize method.
type VariablesFactory interface {
	NewVariables() Variables
}

//go:generate counterfeiter . Variables

// Variables has reference methods.
type Variables interface {
	Get(template.VariableDefinition) (interface{}, bool, error)
	List() ([]template.VariableDefinition, error)
}
