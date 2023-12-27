// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/xianfuhui/fabric/orderer/consensus/etcdraft"
	"github.com/xianfuhui/fabric/protos/orderer"
)

type FakeRPC struct {
	SendConsensusStub        func(uint64, *orderer.ConsensusRequest) error
	sendConsensusMutex       sync.RWMutex
	sendConsensusArgsForCall []struct {
		arg1 uint64
		arg2 *orderer.ConsensusRequest
	}
	sendConsensusReturns struct {
		result1 error
	}
	sendConsensusReturnsOnCall map[int]struct {
		result1 error
	}
	SendSubmitStub        func(uint64, *orderer.SubmitRequest) error
	sendSubmitMutex       sync.RWMutex
	sendSubmitArgsForCall []struct {
		arg1 uint64
		arg2 *orderer.SubmitRequest
	}
	sendSubmitReturns struct {
		result1 error
	}
	sendSubmitReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRPC) SendConsensus(arg1 uint64, arg2 *orderer.ConsensusRequest) error {
	fake.sendConsensusMutex.Lock()
	ret, specificReturn := fake.sendConsensusReturnsOnCall[len(fake.sendConsensusArgsForCall)]
	fake.sendConsensusArgsForCall = append(fake.sendConsensusArgsForCall, struct {
		arg1 uint64
		arg2 *orderer.ConsensusRequest
	}{arg1, arg2})
	fake.recordInvocation("SendConsensus", []interface{}{arg1, arg2})
	fake.sendConsensusMutex.Unlock()
	if fake.SendConsensusStub != nil {
		return fake.SendConsensusStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sendConsensusReturns
	return fakeReturns.result1
}

func (fake *FakeRPC) SendConsensusCallCount() int {
	fake.sendConsensusMutex.RLock()
	defer fake.sendConsensusMutex.RUnlock()
	return len(fake.sendConsensusArgsForCall)
}

func (fake *FakeRPC) SendConsensusCalls(stub func(uint64, *orderer.ConsensusRequest) error) {
	fake.sendConsensusMutex.Lock()
	defer fake.sendConsensusMutex.Unlock()
	fake.SendConsensusStub = stub
}

func (fake *FakeRPC) SendConsensusArgsForCall(i int) (uint64, *orderer.ConsensusRequest) {
	fake.sendConsensusMutex.RLock()
	defer fake.sendConsensusMutex.RUnlock()
	argsForCall := fake.sendConsensusArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRPC) SendConsensusReturns(result1 error) {
	fake.sendConsensusMutex.Lock()
	defer fake.sendConsensusMutex.Unlock()
	fake.SendConsensusStub = nil
	fake.sendConsensusReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRPC) SendConsensusReturnsOnCall(i int, result1 error) {
	fake.sendConsensusMutex.Lock()
	defer fake.sendConsensusMutex.Unlock()
	fake.SendConsensusStub = nil
	if fake.sendConsensusReturnsOnCall == nil {
		fake.sendConsensusReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendConsensusReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRPC) SendSubmit(arg1 uint64, arg2 *orderer.SubmitRequest) error {
	fake.sendSubmitMutex.Lock()
	ret, specificReturn := fake.sendSubmitReturnsOnCall[len(fake.sendSubmitArgsForCall)]
	fake.sendSubmitArgsForCall = append(fake.sendSubmitArgsForCall, struct {
		arg1 uint64
		arg2 *orderer.SubmitRequest
	}{arg1, arg2})
	fake.recordInvocation("SendSubmit", []interface{}{arg1, arg2})
	fake.sendSubmitMutex.Unlock()
	if fake.SendSubmitStub != nil {
		return fake.SendSubmitStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.sendSubmitReturns
	return fakeReturns.result1
}

func (fake *FakeRPC) SendSubmitCallCount() int {
	fake.sendSubmitMutex.RLock()
	defer fake.sendSubmitMutex.RUnlock()
	return len(fake.sendSubmitArgsForCall)
}

func (fake *FakeRPC) SendSubmitCalls(stub func(uint64, *orderer.SubmitRequest) error) {
	fake.sendSubmitMutex.Lock()
	defer fake.sendSubmitMutex.Unlock()
	fake.SendSubmitStub = stub
}

func (fake *FakeRPC) SendSubmitArgsForCall(i int) (uint64, *orderer.SubmitRequest) {
	fake.sendSubmitMutex.RLock()
	defer fake.sendSubmitMutex.RUnlock()
	argsForCall := fake.sendSubmitArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRPC) SendSubmitReturns(result1 error) {
	fake.sendSubmitMutex.Lock()
	defer fake.sendSubmitMutex.Unlock()
	fake.SendSubmitStub = nil
	fake.sendSubmitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRPC) SendSubmitReturnsOnCall(i int, result1 error) {
	fake.sendSubmitMutex.Lock()
	defer fake.sendSubmitMutex.Unlock()
	fake.SendSubmitStub = nil
	if fake.sendSubmitReturnsOnCall == nil {
		fake.sendSubmitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendSubmitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeRPC) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.sendConsensusMutex.RLock()
	defer fake.sendConsensusMutex.RUnlock()
	fake.sendSubmitMutex.RLock()
	defer fake.sendSubmitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRPC) recordInvocation(key string, args []interface{}) {
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

var _ etcdraft.RPC = new(FakeRPC)
