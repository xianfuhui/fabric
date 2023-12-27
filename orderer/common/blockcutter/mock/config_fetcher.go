// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	sync "sync"

	channelconfig "github.com/xianfuhui/fabric/common/channelconfig"
)

type OrdererConfigFetcher struct {
	OrdererConfigStub        func() (channelconfig.Orderer, bool)
	ordererConfigMutex       sync.RWMutex
	ordererConfigArgsForCall []struct {
	}
	ordererConfigReturns struct {
		result1 channelconfig.Orderer
		result2 bool
	}
	ordererConfigReturnsOnCall map[int]struct {
		result1 channelconfig.Orderer
		result2 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *OrdererConfigFetcher) OrdererConfig() (channelconfig.Orderer, bool) {
	fake.ordererConfigMutex.Lock()
	ret, specificReturn := fake.ordererConfigReturnsOnCall[len(fake.ordererConfigArgsForCall)]
	fake.ordererConfigArgsForCall = append(fake.ordererConfigArgsForCall, struct {
	}{})
	fake.recordInvocation("OrdererConfig", []interface{}{})
	fake.ordererConfigMutex.Unlock()
	if fake.OrdererConfigStub != nil {
		return fake.OrdererConfigStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.ordererConfigReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *OrdererConfigFetcher) OrdererConfigCallCount() int {
	fake.ordererConfigMutex.RLock()
	defer fake.ordererConfigMutex.RUnlock()
	return len(fake.ordererConfigArgsForCall)
}

func (fake *OrdererConfigFetcher) OrdererConfigCalls(stub func() (channelconfig.Orderer, bool)) {
	fake.ordererConfigMutex.Lock()
	defer fake.ordererConfigMutex.Unlock()
	fake.OrdererConfigStub = stub
}

func (fake *OrdererConfigFetcher) OrdererConfigReturns(result1 channelconfig.Orderer, result2 bool) {
	fake.ordererConfigMutex.Lock()
	defer fake.ordererConfigMutex.Unlock()
	fake.OrdererConfigStub = nil
	fake.ordererConfigReturns = struct {
		result1 channelconfig.Orderer
		result2 bool
	}{result1, result2}
}

func (fake *OrdererConfigFetcher) OrdererConfigReturnsOnCall(i int, result1 channelconfig.Orderer, result2 bool) {
	fake.ordererConfigMutex.Lock()
	defer fake.ordererConfigMutex.Unlock()
	fake.OrdererConfigStub = nil
	if fake.ordererConfigReturnsOnCall == nil {
		fake.ordererConfigReturnsOnCall = make(map[int]struct {
			result1 channelconfig.Orderer
			result2 bool
		})
	}
	fake.ordererConfigReturnsOnCall[i] = struct {
		result1 channelconfig.Orderer
		result2 bool
	}{result1, result2}
}

func (fake *OrdererConfigFetcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.ordererConfigMutex.RLock()
	defer fake.ordererConfigMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *OrdererConfigFetcher) recordInvocation(key string, args []interface{}) {
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
