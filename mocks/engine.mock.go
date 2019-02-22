// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	"github.com/RTradeLtd/Lens/engine"
)

type FakeEngineSearcher struct {
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	IndexStub        func(engine.Document) error
	indexMutex       sync.RWMutex
	indexArgsForCall []struct {
		arg1 engine.Document
	}
	indexReturns struct {
		result1 error
	}
	indexReturnsOnCall map[int]struct {
		result1 error
	}
	IsIndexedStub        func(string) bool
	isIndexedMutex       sync.RWMutex
	isIndexedArgsForCall []struct {
		arg1 string
	}
	isIndexedReturns struct {
		result1 bool
	}
	isIndexedReturnsOnCall map[int]struct {
		result1 bool
	}
	RemoveStub        func(string)
	removeMutex       sync.RWMutex
	removeArgsForCall []struct {
		arg1 string
	}
	SearchStub        func(context.Context, engine.Query) ([]engine.Result, error)
	searchMutex       sync.RWMutex
	searchArgsForCall []struct {
		arg1 context.Context
		arg2 engine.Query
	}
	searchReturns struct {
		result1 []engine.Result
		result2 error
	}
	searchReturnsOnCall map[int]struct {
		result1 []engine.Result
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEngineSearcher) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		fake.CloseStub()
	}
}

func (fake *FakeEngineSearcher) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeEngineSearcher) CloseCalls(stub func()) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeEngineSearcher) Index(arg1 engine.Document) error {
	fake.indexMutex.Lock()
	ret, specificReturn := fake.indexReturnsOnCall[len(fake.indexArgsForCall)]
	fake.indexArgsForCall = append(fake.indexArgsForCall, struct {
		arg1 engine.Document
	}{arg1})
	fake.recordInvocation("Index", []interface{}{arg1})
	fake.indexMutex.Unlock()
	if fake.IndexStub != nil {
		return fake.IndexStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.indexReturns
	return fakeReturns.result1
}

func (fake *FakeEngineSearcher) IndexCallCount() int {
	fake.indexMutex.RLock()
	defer fake.indexMutex.RUnlock()
	return len(fake.indexArgsForCall)
}

func (fake *FakeEngineSearcher) IndexCalls(stub func(engine.Document) error) {
	fake.indexMutex.Lock()
	defer fake.indexMutex.Unlock()
	fake.IndexStub = stub
}

func (fake *FakeEngineSearcher) IndexArgsForCall(i int) engine.Document {
	fake.indexMutex.RLock()
	defer fake.indexMutex.RUnlock()
	argsForCall := fake.indexArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEngineSearcher) IndexReturns(result1 error) {
	fake.indexMutex.Lock()
	defer fake.indexMutex.Unlock()
	fake.IndexStub = nil
	fake.indexReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeEngineSearcher) IndexReturnsOnCall(i int, result1 error) {
	fake.indexMutex.Lock()
	defer fake.indexMutex.Unlock()
	fake.IndexStub = nil
	if fake.indexReturnsOnCall == nil {
		fake.indexReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.indexReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeEngineSearcher) IsIndexed(arg1 string) bool {
	fake.isIndexedMutex.Lock()
	ret, specificReturn := fake.isIndexedReturnsOnCall[len(fake.isIndexedArgsForCall)]
	fake.isIndexedArgsForCall = append(fake.isIndexedArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("IsIndexed", []interface{}{arg1})
	fake.isIndexedMutex.Unlock()
	if fake.IsIndexedStub != nil {
		return fake.IsIndexedStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.isIndexedReturns
	return fakeReturns.result1
}

func (fake *FakeEngineSearcher) IsIndexedCallCount() int {
	fake.isIndexedMutex.RLock()
	defer fake.isIndexedMutex.RUnlock()
	return len(fake.isIndexedArgsForCall)
}

func (fake *FakeEngineSearcher) IsIndexedCalls(stub func(string) bool) {
	fake.isIndexedMutex.Lock()
	defer fake.isIndexedMutex.Unlock()
	fake.IsIndexedStub = stub
}

func (fake *FakeEngineSearcher) IsIndexedArgsForCall(i int) string {
	fake.isIndexedMutex.RLock()
	defer fake.isIndexedMutex.RUnlock()
	argsForCall := fake.isIndexedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEngineSearcher) IsIndexedReturns(result1 bool) {
	fake.isIndexedMutex.Lock()
	defer fake.isIndexedMutex.Unlock()
	fake.IsIndexedStub = nil
	fake.isIndexedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeEngineSearcher) IsIndexedReturnsOnCall(i int, result1 bool) {
	fake.isIndexedMutex.Lock()
	defer fake.isIndexedMutex.Unlock()
	fake.IsIndexedStub = nil
	if fake.isIndexedReturnsOnCall == nil {
		fake.isIndexedReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isIndexedReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeEngineSearcher) Remove(arg1 string) {
	fake.removeMutex.Lock()
	fake.removeArgsForCall = append(fake.removeArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Remove", []interface{}{arg1})
	fake.removeMutex.Unlock()
	if fake.RemoveStub != nil {
		fake.RemoveStub(arg1)
	}
}

func (fake *FakeEngineSearcher) RemoveCallCount() int {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return len(fake.removeArgsForCall)
}

func (fake *FakeEngineSearcher) RemoveCalls(stub func(string)) {
	fake.removeMutex.Lock()
	defer fake.removeMutex.Unlock()
	fake.RemoveStub = stub
}

func (fake *FakeEngineSearcher) RemoveArgsForCall(i int) string {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	argsForCall := fake.removeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEngineSearcher) Search(arg1 context.Context, arg2 engine.Query) ([]engine.Result, error) {
	fake.searchMutex.Lock()
	ret, specificReturn := fake.searchReturnsOnCall[len(fake.searchArgsForCall)]
	fake.searchArgsForCall = append(fake.searchArgsForCall, struct {
		arg1 context.Context
		arg2 engine.Query
	}{arg1, arg2})
	fake.recordInvocation("Search", []interface{}{arg1, arg2})
	fake.searchMutex.Unlock()
	if fake.SearchStub != nil {
		return fake.SearchStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.searchReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEngineSearcher) SearchCallCount() int {
	fake.searchMutex.RLock()
	defer fake.searchMutex.RUnlock()
	return len(fake.searchArgsForCall)
}

func (fake *FakeEngineSearcher) SearchCalls(stub func(context.Context, engine.Query) ([]engine.Result, error)) {
	fake.searchMutex.Lock()
	defer fake.searchMutex.Unlock()
	fake.SearchStub = stub
}

func (fake *FakeEngineSearcher) SearchArgsForCall(i int) (context.Context, engine.Query) {
	fake.searchMutex.RLock()
	defer fake.searchMutex.RUnlock()
	argsForCall := fake.searchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEngineSearcher) SearchReturns(result1 []engine.Result, result2 error) {
	fake.searchMutex.Lock()
	defer fake.searchMutex.Unlock()
	fake.SearchStub = nil
	fake.searchReturns = struct {
		result1 []engine.Result
		result2 error
	}{result1, result2}
}

func (fake *FakeEngineSearcher) SearchReturnsOnCall(i int, result1 []engine.Result, result2 error) {
	fake.searchMutex.Lock()
	defer fake.searchMutex.Unlock()
	fake.SearchStub = nil
	if fake.searchReturnsOnCall == nil {
		fake.searchReturnsOnCall = make(map[int]struct {
			result1 []engine.Result
			result2 error
		})
	}
	fake.searchReturnsOnCall[i] = struct {
		result1 []engine.Result
		result2 error
	}{result1, result2}
}

func (fake *FakeEngineSearcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.indexMutex.RLock()
	defer fake.indexMutex.RUnlock()
	fake.isIndexedMutex.RLock()
	defer fake.isIndexedMutex.RUnlock()
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	fake.searchMutex.RLock()
	defer fake.searchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEngineSearcher) recordInvocation(key string, args []interface{}) {
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

var _ engine.Searcher = new(FakeEngineSearcher)
