// Code generated by counterfeiter. DO NOT EDIT.
package vaultfakes

import (
	sync "sync"

	vault "github.com/cappyzawa/cf-push-with-vault/vault"
	template "github.com/cloudfoundry/bosh-cli/director/template"
)

type FakeVariables struct {
	GetStub        func(template.VariableDefinition) (interface{}, bool, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 template.VariableDefinition
	}
	getReturns struct {
		result1 interface{}
		result2 bool
		result3 error
	}
	getReturnsOnCall map[int]struct {
		result1 interface{}
		result2 bool
		result3 error
	}
	ListStub        func() ([]template.VariableDefinition, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
	}
	listReturns struct {
		result1 []template.VariableDefinition
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 []template.VariableDefinition
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVariables) Get(arg1 template.VariableDefinition) (interface{}, bool, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 template.VariableDefinition
	}{arg1})
	fake.recordInvocation("Get", []interface{}{arg1})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeVariables) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeVariables) GetCalls(stub func(template.VariableDefinition) (interface{}, bool, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeVariables) GetArgsForCall(i int) template.VariableDefinition {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeVariables) GetReturns(result1 interface{}, result2 bool, result3 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 interface{}
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVariables) GetReturnsOnCall(i int, result1 interface{}, result2 bool, result3 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 bool
			result3 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 interface{}
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeVariables) List() ([]template.VariableDefinition, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
	}{})
	fake.recordInvocation("List", []interface{}{})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVariables) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeVariables) ListCalls(stub func() ([]template.VariableDefinition, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeVariables) ListReturns(result1 []template.VariableDefinition, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []template.VariableDefinition
		result2 error
	}{result1, result2}
}

func (fake *FakeVariables) ListReturnsOnCall(i int, result1 []template.VariableDefinition, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 []template.VariableDefinition
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 []template.VariableDefinition
		result2 error
	}{result1, result2}
}

func (fake *FakeVariables) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVariables) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ vault.Variables = new(FakeVariables)
