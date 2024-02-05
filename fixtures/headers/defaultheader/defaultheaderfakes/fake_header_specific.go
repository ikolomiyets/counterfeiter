// This is a specific header for only some of the fakes in this package
//

// Code generated by counterfeiter. DO NOT EDIT.
package defaultheaderfakes

import (
	"sync"

	"github.com/ikolomiyets/counterfeiter/v6/fixtures/headers/defaultheader"
)

type FakeHeaderSpecific struct {
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHeaderSpecific) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHeaderSpecific) recordInvocation(key string, args []interface{}) {
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

var _ defaultheader.HeaderSpecific = new(FakeHeaderSpecific)
