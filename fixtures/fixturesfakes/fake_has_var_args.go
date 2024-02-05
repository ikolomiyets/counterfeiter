// Code generated by counterfeiter. DO NOT EDIT.
package fixturesfakes

import (
	"sync"

	"github.com/ikolomiyets/counterfeiter/v6/fixtures"
)

type FakeHasVarArgs struct {
	DoMoreThingsStub        func(int, int, ...string) int
	doMoreThingsMutex       sync.RWMutex
	doMoreThingsArgsForCall []struct {
		arg1 int
		arg2 int
		arg3 []string
	}
	doMoreThingsReturns struct {
		result1 int
	}
	doMoreThingsReturnsOnCall map[int]struct {
		result1 int
	}
	DoThingsStub        func(int, ...string) int
	doThingsMutex       sync.RWMutex
	doThingsArgsForCall []struct {
		arg1 int
		arg2 []string
	}
	doThingsReturns struct {
		result1 int
	}
	doThingsReturnsOnCall map[int]struct {
		result1 int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHasVarArgs) DoMoreThings(arg1 int, arg2 int, arg3 ...string) int {
	fake.doMoreThingsMutex.Lock()
	ret, specificReturn := fake.doMoreThingsReturnsOnCall[len(fake.doMoreThingsArgsForCall)]
	fake.doMoreThingsArgsForCall = append(fake.doMoreThingsArgsForCall, struct {
		arg1 int
		arg2 int
		arg3 []string
	}{arg1, arg2, arg3})
	stub := fake.DoMoreThingsStub
	fakeReturns := fake.doMoreThingsReturns
	fake.recordInvocation("DoMoreThings", []interface{}{arg1, arg2, arg3})
	fake.doMoreThingsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHasVarArgs) DoMoreThingsCallCount() int {
	fake.doMoreThingsMutex.RLock()
	defer fake.doMoreThingsMutex.RUnlock()
	return len(fake.doMoreThingsArgsForCall)
}

func (fake *FakeHasVarArgs) DoMoreThingsCalls(stub func(int, int, ...string) int) {
	fake.doMoreThingsMutex.Lock()
	defer fake.doMoreThingsMutex.Unlock()
	fake.DoMoreThingsStub = stub
}

func (fake *FakeHasVarArgs) DoMoreThingsArgsForCall(i int) (int, int, []string) {
	fake.doMoreThingsMutex.RLock()
	defer fake.doMoreThingsMutex.RUnlock()
	argsForCall := fake.doMoreThingsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeHasVarArgs) DoMoreThingsReturns(result1 int) {
	fake.doMoreThingsMutex.Lock()
	defer fake.doMoreThingsMutex.Unlock()
	fake.DoMoreThingsStub = nil
	fake.doMoreThingsReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeHasVarArgs) DoMoreThingsReturnsOnCall(i int, result1 int) {
	fake.doMoreThingsMutex.Lock()
	defer fake.doMoreThingsMutex.Unlock()
	fake.DoMoreThingsStub = nil
	if fake.doMoreThingsReturnsOnCall == nil {
		fake.doMoreThingsReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.doMoreThingsReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeHasVarArgs) DoThings(arg1 int, arg2 ...string) int {
	fake.doThingsMutex.Lock()
	ret, specificReturn := fake.doThingsReturnsOnCall[len(fake.doThingsArgsForCall)]
	fake.doThingsArgsForCall = append(fake.doThingsArgsForCall, struct {
		arg1 int
		arg2 []string
	}{arg1, arg2})
	stub := fake.DoThingsStub
	fakeReturns := fake.doThingsReturns
	fake.recordInvocation("DoThings", []interface{}{arg1, arg2})
	fake.doThingsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHasVarArgs) DoThingsCallCount() int {
	fake.doThingsMutex.RLock()
	defer fake.doThingsMutex.RUnlock()
	return len(fake.doThingsArgsForCall)
}

func (fake *FakeHasVarArgs) DoThingsCalls(stub func(int, ...string) int) {
	fake.doThingsMutex.Lock()
	defer fake.doThingsMutex.Unlock()
	fake.DoThingsStub = stub
}

func (fake *FakeHasVarArgs) DoThingsArgsForCall(i int) (int, []string) {
	fake.doThingsMutex.RLock()
	defer fake.doThingsMutex.RUnlock()
	argsForCall := fake.doThingsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeHasVarArgs) DoThingsReturns(result1 int) {
	fake.doThingsMutex.Lock()
	defer fake.doThingsMutex.Unlock()
	fake.DoThingsStub = nil
	fake.doThingsReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeHasVarArgs) DoThingsReturnsOnCall(i int, result1 int) {
	fake.doThingsMutex.Lock()
	defer fake.doThingsMutex.Unlock()
	fake.DoThingsStub = nil
	if fake.doThingsReturnsOnCall == nil {
		fake.doThingsReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.doThingsReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeHasVarArgs) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.doMoreThingsMutex.RLock()
	defer fake.doMoreThingsMutex.RUnlock()
	fake.doThingsMutex.RLock()
	defer fake.doThingsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHasVarArgs) recordInvocation(key string, args []interface{}) {
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

var _ fixtures.HasVarArgs = new(FakeHasVarArgs)
