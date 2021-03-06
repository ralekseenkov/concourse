// Code generated by counterfeiter. DO NOT EDIT.
package enginefakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/engine"
)

type FakeEngine struct {
	LookupBuildStub        func(lager.Logger, db.Build) (engine.Build, error)
	lookupBuildMutex       sync.RWMutex
	lookupBuildArgsForCall []struct {
		arg1 lager.Logger
		arg2 db.Build
	}
	lookupBuildReturns struct {
		result1 engine.Build
		result2 error
	}
	lookupBuildReturnsOnCall map[int]struct {
		result1 engine.Build
		result2 error
	}
	ReleaseAllStub        func(lager.Logger)
	releaseAllMutex       sync.RWMutex
	releaseAllArgsForCall []struct {
		arg1 lager.Logger
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEngine) LookupBuild(arg1 lager.Logger, arg2 db.Build) (engine.Build, error) {
	fake.lookupBuildMutex.Lock()
	ret, specificReturn := fake.lookupBuildReturnsOnCall[len(fake.lookupBuildArgsForCall)]
	fake.lookupBuildArgsForCall = append(fake.lookupBuildArgsForCall, struct {
		arg1 lager.Logger
		arg2 db.Build
	}{arg1, arg2})
	fake.recordInvocation("LookupBuild", []interface{}{arg1, arg2})
	fake.lookupBuildMutex.Unlock()
	if fake.LookupBuildStub != nil {
		return fake.LookupBuildStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.lookupBuildReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEngine) LookupBuildCallCount() int {
	fake.lookupBuildMutex.RLock()
	defer fake.lookupBuildMutex.RUnlock()
	return len(fake.lookupBuildArgsForCall)
}

func (fake *FakeEngine) LookupBuildCalls(stub func(lager.Logger, db.Build) (engine.Build, error)) {
	fake.lookupBuildMutex.Lock()
	defer fake.lookupBuildMutex.Unlock()
	fake.LookupBuildStub = stub
}

func (fake *FakeEngine) LookupBuildArgsForCall(i int) (lager.Logger, db.Build) {
	fake.lookupBuildMutex.RLock()
	defer fake.lookupBuildMutex.RUnlock()
	argsForCall := fake.lookupBuildArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEngine) LookupBuildReturns(result1 engine.Build, result2 error) {
	fake.lookupBuildMutex.Lock()
	defer fake.lookupBuildMutex.Unlock()
	fake.LookupBuildStub = nil
	fake.lookupBuildReturns = struct {
		result1 engine.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeEngine) LookupBuildReturnsOnCall(i int, result1 engine.Build, result2 error) {
	fake.lookupBuildMutex.Lock()
	defer fake.lookupBuildMutex.Unlock()
	fake.LookupBuildStub = nil
	if fake.lookupBuildReturnsOnCall == nil {
		fake.lookupBuildReturnsOnCall = make(map[int]struct {
			result1 engine.Build
			result2 error
		})
	}
	fake.lookupBuildReturnsOnCall[i] = struct {
		result1 engine.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeEngine) ReleaseAll(arg1 lager.Logger) {
	fake.releaseAllMutex.Lock()
	fake.releaseAllArgsForCall = append(fake.releaseAllArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("ReleaseAll", []interface{}{arg1})
	fake.releaseAllMutex.Unlock()
	if fake.ReleaseAllStub != nil {
		fake.ReleaseAllStub(arg1)
	}
}

func (fake *FakeEngine) ReleaseAllCallCount() int {
	fake.releaseAllMutex.RLock()
	defer fake.releaseAllMutex.RUnlock()
	return len(fake.releaseAllArgsForCall)
}

func (fake *FakeEngine) ReleaseAllCalls(stub func(lager.Logger)) {
	fake.releaseAllMutex.Lock()
	defer fake.releaseAllMutex.Unlock()
	fake.ReleaseAllStub = stub
}

func (fake *FakeEngine) ReleaseAllArgsForCall(i int) lager.Logger {
	fake.releaseAllMutex.RLock()
	defer fake.releaseAllMutex.RUnlock()
	argsForCall := fake.releaseAllArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEngine) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.lookupBuildMutex.RLock()
	defer fake.lookupBuildMutex.RUnlock()
	fake.releaseAllMutex.RLock()
	defer fake.releaseAllMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEngine) recordInvocation(key string, args []interface{}) {
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

var _ engine.Engine = new(FakeEngine)
