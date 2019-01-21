
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//伪造者生成的代码。不要编辑。
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/statedb"
	"github.com/hyperledger/fabric/core/ledger/kvledger/txmgmt/version"
)

type VersionedDB struct {
	GetStateStub        func(namespace string, key string) (*statedb.VersionedValue, error)
	getStateMutex       sync.RWMutex
	getStateArgsForCall []struct {
		namespace string
		key       string
	}
	getStateReturns struct {
		result1 *statedb.VersionedValue
		result2 error
	}
	getStateReturnsOnCall map[int]struct {
		result1 *statedb.VersionedValue
		result2 error
	}
	GetVersionStub        func(namespace string, key string) (*version.Height, error)
	getVersionMutex       sync.RWMutex
	getVersionArgsForCall []struct {
		namespace string
		key       string
	}
	getVersionReturns struct {
		result1 *version.Height
		result2 error
	}
	getVersionReturnsOnCall map[int]struct {
		result1 *version.Height
		result2 error
	}
	GetStateMultipleKeysStub        func(namespace string, keys []string) ([]*statedb.VersionedValue, error)
	getStateMultipleKeysMutex       sync.RWMutex
	getStateMultipleKeysArgsForCall []struct {
		namespace string
		keys      []string
	}
	getStateMultipleKeysReturns struct {
		result1 []*statedb.VersionedValue
		result2 error
	}
	getStateMultipleKeysReturnsOnCall map[int]struct {
		result1 []*statedb.VersionedValue
		result2 error
	}
	GetStateRangeScanIteratorStub        func(namespace string, startKey string, endKey string) (statedb.ResultsIterator, error)
	getStateRangeScanIteratorMutex       sync.RWMutex
	getStateRangeScanIteratorArgsForCall []struct {
		namespace string
		startKey  string
		endKey    string
	}
	getStateRangeScanIteratorReturns struct {
		result1 statedb.ResultsIterator
		result2 error
	}
	getStateRangeScanIteratorReturnsOnCall map[int]struct {
		result1 statedb.ResultsIterator
		result2 error
	}
	GetStateRangeScanIteratorWithMetadataStub        func(namespace string, startKey string, endKey string, metadata map[string]interface{}) (statedb.QueryResultsIterator, error)
	getStateRangeScanIteratorWithMetadataMutex       sync.RWMutex
	getStateRangeScanIteratorWithMetadataArgsForCall []struct {
		namespace string
		startKey  string
		endKey    string
		metadata  map[string]interface{}
	}
	getStateRangeScanIteratorWithMetadataReturns struct {
		result1 statedb.QueryResultsIterator
		result2 error
	}
	getStateRangeScanIteratorWithMetadataReturnsOnCall map[int]struct {
		result1 statedb.QueryResultsIterator
		result2 error
	}
	ExecuteQueryStub        func(namespace, query string) (statedb.ResultsIterator, error)
	executeQueryMutex       sync.RWMutex
	executeQueryArgsForCall []struct {
		namespace string
		query     string
	}
	executeQueryReturns struct {
		result1 statedb.ResultsIterator
		result2 error
	}
	executeQueryReturnsOnCall map[int]struct {
		result1 statedb.ResultsIterator
		result2 error
	}
	ExecuteQueryWithMetadataStub        func(namespace, query string, metadata map[string]interface{}) (statedb.QueryResultsIterator, error)
	executeQueryWithMetadataMutex       sync.RWMutex
	executeQueryWithMetadataArgsForCall []struct {
		namespace string
		query     string
		metadata  map[string]interface{}
	}
	executeQueryWithMetadataReturns struct {
		result1 statedb.QueryResultsIterator
		result2 error
	}
	executeQueryWithMetadataReturnsOnCall map[int]struct {
		result1 statedb.QueryResultsIterator
		result2 error
	}
	ApplyUpdatesStub        func(batch *statedb.UpdateBatch, height *version.Height) error
	applyUpdatesMutex       sync.RWMutex
	applyUpdatesArgsForCall []struct {
		batch  *statedb.UpdateBatch
		height *version.Height
	}
	applyUpdatesReturns struct {
		result1 error
	}
	applyUpdatesReturnsOnCall map[int]struct {
		result1 error
	}
	GetLatestSavePointStub        func() (*version.Height, error)
	getLatestSavePointMutex       sync.RWMutex
	getLatestSavePointArgsForCall []struct{}
	getLatestSavePointReturns     struct {
		result1 *version.Height
		result2 error
	}
	getLatestSavePointReturnsOnCall map[int]struct {
		result1 *version.Height
		result2 error
	}
	ValidateKeyValueStub        func(key string, value []byte) error
	validateKeyValueMutex       sync.RWMutex
	validateKeyValueArgsForCall []struct {
		key   string
		value []byte
	}
	validateKeyValueReturns struct {
		result1 error
	}
	validateKeyValueReturnsOnCall map[int]struct {
		result1 error
	}
	BytesKeySupportedStub        func() bool
	bytesKeySupportedMutex       sync.RWMutex
	bytesKeySupportedArgsForCall []struct{}
	bytesKeySupportedReturns     struct {
		result1 bool
	}
	bytesKeySupportedReturnsOnCall map[int]struct {
		result1 bool
	}
	OpenStub        func() error
	openMutex       sync.RWMutex
	openArgsForCall []struct{}
	openReturns     struct {
		result1 error
	}
	openReturnsOnCall map[int]struct {
		result1 error
	}
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct{}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *VersionedDB) GetState(namespace string, key string) (*statedb.VersionedValue, error) {
	fake.getStateMutex.Lock()
	ret, specificReturn := fake.getStateReturnsOnCall[len(fake.getStateArgsForCall)]
	fake.getStateArgsForCall = append(fake.getStateArgsForCall, struct {
		namespace string
		key       string
	}{namespace, key})
	fake.recordInvocation("GetState", []interface{}{namespace, key})
	fake.getStateMutex.Unlock()
	if fake.GetStateStub != nil {
		return fake.GetStateStub(namespace, key)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getStateReturns.result1, fake.getStateReturns.result2
}

func (fake *VersionedDB) GetStateCallCount() int {
	fake.getStateMutex.RLock()
	defer fake.getStateMutex.RUnlock()
	return len(fake.getStateArgsForCall)
}

func (fake *VersionedDB) GetStateArgsForCall(i int) (string, string) {
	fake.getStateMutex.RLock()
	defer fake.getStateMutex.RUnlock()
	return fake.getStateArgsForCall[i].namespace, fake.getStateArgsForCall[i].key
}

func (fake *VersionedDB) GetStateReturns(result1 *statedb.VersionedValue, result2 error) {
	fake.GetStateStub = nil
	fake.getStateReturns = struct {
		result1 *statedb.VersionedValue
		result2 error
	}{result1, result2}
}

func (fake *VersionedDB) GetStateReturnsOnCall(i int, result1 *statedb.VersionedValue, result2 error) {
	fake.GetStateStub = nil
	if fake.getStateReturnsOnCall == nil {
		fake.getStateReturnsOnCall = make(map[int]struct {
			result1 *statedb.VersionedValue
			result2 error
		})
	}
	fake.getStateReturnsOnCall[i] = struct {
		result1 *statedb.VersionedValue
		result2 error
	}{result1, result2}
}

func (fake *VersionedDB) GetVersion(namespace string, key string) (*version.Height, error) {
	fake.getVersionMutex.Lock()
	ret, specificReturn := fake.getVersionReturnsOnCall[len(fake.getVersionArgsForCall)]
	fake.getVersionArgsForCall = append(fake.getVersionArgsForCall, struct {
		namespace string
		key       string
	}{namespace, key})
	fake.recordInvocation("GetVersion", []interface{}{namespace, key})
	fake.getVersionMutex.Unlock()
	if fake.GetVersionStub != nil {
		return fake.GetVersionStub(namespace, key)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getVersionReturns.result1, fake.getVersionReturns.result2
}

func (fake *VersionedDB) GetVersionCallCount() int {
	fake.getVersionMutex.RLock()
	defer fake.getVersionMutex.RUnlock()
	return len(fake.getVersionArgsForCall)
}

func (fake *VersionedDB) GetVersionArgsForCall(i int) (string, string) {
	fake.getVersionMutex.RLock()
	defer fake.getVersionMutex.RUnlock()
	return fake.getVersionArgsForCall[i].namespace, fake.getVersionArgsForCall[i].key
}

func (fake *VersionedDB) GetVersionReturns(result1 *version.Height, result2 error) {
	fake.GetVersionStub = nil
	fake.getVersionReturns = struct {
		result1 *version.Height
		result2 error
	}{result1, result2}
}

func (fake *VersionedDB) GetVersionReturnsOnCall(i int, result1 *version.Height, result2 error) {
	fake.GetVersionStub = nil
	if fake.getVersionReturnsOnCall == nil {
		fake.getVersionReturnsOnCall = make(map[int]struct {
			result1 *version.Height
			result2 error
		})
	}
	fake.getVersionReturnsOnCall[i] = struct {
		result1 *version.Height
		result2 error
	}{result1, result2}
}

func (fake *VersionedDB) GetStateMultipleKeys(namespace string, keys []string) ([]*statedb.VersionedValue, error) {
	var keysCopy []string
	if keys != nil {
		keysCopy = make([]string, len(keys))
		copy(keysCopy, keys)
	}
	fake.getStateMultipleKeysMutex.Lock()
	ret, specificReturn := fake.getStateMultipleKeysReturnsOnCall[len(fake.getStateMultipleKeysArgsForCall)]
	fake.getStateMultipleKeysArgsForCall = append(fake.getStateMultipleKeysArgsForCall, struct {
		namespace string
		keys      []string
	}{namespace, keysCopy})
	fake.recordInvocation("GetStateMultipleKeys", []interface{}{namespace, keysCopy})
	fake.getStateMultipleKeysMutex.Unlock()
	if fake.GetStateMultipleKeysStub != nil {
		return fake.GetStateMultipleKeysStub(namespace, keys)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getStateMultipleKeysReturns.result1, fake.getStateMultipleKeysReturns.result2
}

func (fake *VersionedDB) GetStateMultipleKeysCallCount() int {
	fake.getStateMultipleKeysMutex.RLock()
	defer fake.getStateMultipleKeysMutex.RUnlock()
	return len(fake.getStateMultipleKeysArgsForCall)
}

func (fake *VersionedDB) GetStateMultipleKeysArgsForCall(i int) (string, []string) {
	fake.getStateMultipleKeysMutex.RLock()
	defer fake.getStateMultipleKeysMutex.RUnlock()
	return fake.getStateMultipleKeysArgsForCall[i].namespace, fake.getStateMultipleKeysArgsForCall[i].keys
}

func (fake *VersionedDB) GetStateMultipleKeysReturns(result1 []*statedb.VersionedValue, result2 error) {
	fake.GetStateMultipleKeysStub = nil
	fake.getStateMultipleKeysReturns = struct {
		result1 []*statedb.VersionedValue
		result2 error
	}{result1, result2}
}

func (fake *VersionedDB) GetStateMultipleKeysReturnsOnCall(i int, result1 []*statedb.VersionedValue, result2 error) {
	fake.GetStateMultipleKeysStub = nil
	if fake.getStateMultipleKeysReturnsOnCall == nil {
		fake.getStateMultipleKeysReturnsOnCall = make(map[int]struct {
			result1 []*statedb.VersionedValue
			result2 error
		})
	}
	fake.getStateMultipleKeysReturnsOnCall[i] = struct {
		result1 []*statedb.VersionedValue
		result2 error
	}{result1, result2}
}

func (fake *VersionedDB) GetStateRangeScanIterator(namespace string, startKey string, endKey string) (statedb.ResultsIterator, error) {
	fake.getStateRangeScanIteratorMutex.Lock()
	ret, specificReturn := fake.getStateRangeScanIteratorReturnsOnCall[len(fake.getStateRangeScanIteratorArgsForCall)]
	fake.getStateRangeScanIteratorArgsForCall = append(fake.getStateRangeScanIteratorArgsForCall, struct {
		namespace string
		startKey  string
		endKey    string
	}{namespace, startKey, endKey})
	fake.recordInvocation("GetStateRangeScanIterator", []interface{}{namespace, startKey, endKey})
	fake.getStateRangeScanIteratorMutex.Unlock()
	if fake.GetStateRangeScanIteratorStub != nil {
		return fake.GetStateRangeScanIteratorStub(namespace, startKey, endKey)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getStateRangeScanIteratorReturns.result1, fake.getStateRangeScanIteratorReturns.result2
}

func (fake *VersionedDB) GetStateRangeScanIteratorCallCount() int {
	fake.getStateRangeScanIteratorMutex.RLock()
	defer fake.getStateRangeScanIteratorMutex.RUnlock()
	return len(fake.getStateRangeScanIteratorArgsForCall)
}

func (fake *VersionedDB) GetStateRangeScanIteratorArgsForCall(i int) (string, string, string) {
	fake.getStateRangeScanIteratorMutex.RLock()
	defer fake.getStateRangeScanIteratorMutex.RUnlock()
	return fake.getStateRangeScanIteratorArgsForCall[i].namespace, fake.getStateRangeScanIteratorArgsForCall[i].startKey, fake.getStateRangeScanIteratorArgsForCall[i].endKey
}

func (fake *VersionedDB) GetStateRangeScanIteratorReturns(result1 statedb.ResultsIterator, result2 error) {
	fake.GetStateRangeScanIteratorStub = nil
	fake.getStateRangeScanIteratorReturns = struct {
		result1 statedb.ResultsIterator
		result2 error
	}{result1, result2}
}

func (fake *VersionedDB) GetStateRangeScanIteratorReturnsOnCall(i int, result1 statedb.ResultsIterator, result2 error) {
	fake.GetStateRangeScanIteratorStub = nil
	if fake.getStateRangeScanIteratorReturnsOnCall == nil {
		fake.getStateRangeScanIteratorReturnsOnCall = make(map[int]struct {
			result1 statedb.ResultsIterator
			result2 error
		})
	}
	fake.getStateRangeScanIteratorReturnsOnCall[i] = struct {
		result1 statedb.ResultsIterator
		result2 error
	}{result1, result2}
}

func (fake *VersionedDB) GetStateRangeScanIteratorWithMetadata(namespace string, startKey string, endKey string, metadata map[string]interface{}) (statedb.QueryResultsIterator, error) {
	fake.getStateRangeScanIteratorWithMetadataMutex.Lock()
	ret, specificReturn := fake.getStateRangeScanIteratorWithMetadataReturnsOnCall[len(fake.getStateRangeScanIteratorWithMetadataArgsForCall)]
	fake.getStateRangeScanIteratorWithMetadataArgsForCall = append(fake.getStateRangeScanIteratorWithMetadataArgsForCall, struct {
		namespace string
		startKey  string
		endKey    string
		metadata  map[string]interface{}
	}{namespace, startKey, endKey, metadata})
	fake.recordInvocation("GetStateRangeScanIteratorWithMetadata", []interface{}{namespace, startKey, endKey, metadata})
	fake.getStateRangeScanIteratorWithMetadataMutex.Unlock()
	if fake.GetStateRangeScanIteratorWithMetadataStub != nil {
		return fake.GetStateRangeScanIteratorWithMetadataStub(namespace, startKey, endKey, metadata)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getStateRangeScanIteratorWithMetadataReturns.result1, fake.getStateRangeScanIteratorWithMetadataReturns.result2
}

func (fake *VersionedDB) GetStateRangeScanIteratorWithMetadataCallCount() int {
	fake.getStateRangeScanIteratorWithMetadataMutex.RLock()
	defer fake.getStateRangeScanIteratorWithMetadataMutex.RUnlock()
	return len(fake.getStateRangeScanIteratorWithMetadataArgsForCall)
}

func (fake *VersionedDB) GetStateRangeScanIteratorWithMetadataArgsForCall(i int) (string, string, string, map[string]interface{}) {
	fake.getStateRangeScanIteratorWithMetadataMutex.RLock()
	defer fake.getStateRangeScanIteratorWithMetadataMutex.RUnlock()
	return fake.getStateRangeScanIteratorWithMetadataArgsForCall[i].namespace, fake.getStateRangeScanIteratorWithMetadataArgsForCall[i].startKey, fake.getStateRangeScanIteratorWithMetadataArgsForCall[i].endKey, fake.getStateRangeScanIteratorWithMetadataArgsForCall[i].metadata
}

func (fake *VersionedDB) GetStateRangeScanIteratorWithMetadataReturns(result1 statedb.QueryResultsIterator, result2 error) {
	fake.GetStateRangeScanIteratorWithMetadataStub = nil
	fake.getStateRangeScanIteratorWithMetadataReturns = struct {
		result1 statedb.QueryResultsIterator
		result2 error
	}{result1, result2}
}

func (fake *VersionedDB) GetStateRangeScanIteratorWithMetadataReturnsOnCall(i int, result1 statedb.QueryResultsIterator, result2 error) {
	fake.GetStateRangeScanIteratorWithMetadataStub = nil
	if fake.getStateRangeScanIteratorWithMetadataReturnsOnCall == nil {
		fake.getStateRangeScanIteratorWithMetadataReturnsOnCall = make(map[int]struct {
			result1 statedb.QueryResultsIterator
			result2 error
		})
	}
	fake.getStateRangeScanIteratorWithMetadataReturnsOnCall[i] = struct {
		result1 statedb.QueryResultsIterator
		result2 error
	}{result1, result2}
}

func (fake *VersionedDB) ExecuteQuery(namespace string, query string) (statedb.ResultsIterator, error) {
	fake.executeQueryMutex.Lock()
	ret, specificReturn := fake.executeQueryReturnsOnCall[len(fake.executeQueryArgsForCall)]
	fake.executeQueryArgsForCall = append(fake.executeQueryArgsForCall, struct {
		namespace string
		query     string
	}{namespace, query})
	fake.recordInvocation("ExecuteQuery", []interface{}{namespace, query})
	fake.executeQueryMutex.Unlock()
	if fake.ExecuteQueryStub != nil {
		return fake.ExecuteQueryStub(namespace, query)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.executeQueryReturns.result1, fake.executeQueryReturns.result2
}

func (fake *VersionedDB) ExecuteQueryCallCount() int {
	fake.executeQueryMutex.RLock()
	defer fake.executeQueryMutex.RUnlock()
	return len(fake.executeQueryArgsForCall)
}

func (fake *VersionedDB) ExecuteQueryArgsForCall(i int) (string, string) {
	fake.executeQueryMutex.RLock()
	defer fake.executeQueryMutex.RUnlock()
	return fake.executeQueryArgsForCall[i].namespace, fake.executeQueryArgsForCall[i].query
}

func (fake *VersionedDB) ExecuteQueryReturns(result1 statedb.ResultsIterator, result2 error) {
	fake.ExecuteQueryStub = nil
	fake.executeQueryReturns = struct {
		result1 statedb.ResultsIterator
		result2 error
	}{result1, result2}
}

func (fake *VersionedDB) ExecuteQueryReturnsOnCall(i int, result1 statedb.ResultsIterator, result2 error) {
	fake.ExecuteQueryStub = nil
	if fake.executeQueryReturnsOnCall == nil {
		fake.executeQueryReturnsOnCall = make(map[int]struct {
			result1 statedb.ResultsIterator
			result2 error
		})
	}
	fake.executeQueryReturnsOnCall[i] = struct {
		result1 statedb.ResultsIterator
		result2 error
	}{result1, result2}
}

func (fake *VersionedDB) ExecuteQueryWithMetadata(namespace string, query string, metadata map[string]interface{}) (statedb.QueryResultsIterator, error) {
	fake.executeQueryWithMetadataMutex.Lock()
	ret, specificReturn := fake.executeQueryWithMetadataReturnsOnCall[len(fake.executeQueryWithMetadataArgsForCall)]
	fake.executeQueryWithMetadataArgsForCall = append(fake.executeQueryWithMetadataArgsForCall, struct {
		namespace string
		query     string
		metadata  map[string]interface{}
	}{namespace, query, metadata})
	fake.recordInvocation("ExecuteQueryWithMetadata", []interface{}{namespace, query, metadata})
	fake.executeQueryWithMetadataMutex.Unlock()
	if fake.ExecuteQueryWithMetadataStub != nil {
		return fake.ExecuteQueryWithMetadataStub(namespace, query, metadata)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.executeQueryWithMetadataReturns.result1, fake.executeQueryWithMetadataReturns.result2
}

func (fake *VersionedDB) ExecuteQueryWithMetadataCallCount() int {
	fake.executeQueryWithMetadataMutex.RLock()
	defer fake.executeQueryWithMetadataMutex.RUnlock()
	return len(fake.executeQueryWithMetadataArgsForCall)
}

func (fake *VersionedDB) ExecuteQueryWithMetadataArgsForCall(i int) (string, string, map[string]interface{}) {
	fake.executeQueryWithMetadataMutex.RLock()
	defer fake.executeQueryWithMetadataMutex.RUnlock()
	return fake.executeQueryWithMetadataArgsForCall[i].namespace, fake.executeQueryWithMetadataArgsForCall[i].query, fake.executeQueryWithMetadataArgsForCall[i].metadata
}

func (fake *VersionedDB) ExecuteQueryWithMetadataReturns(result1 statedb.QueryResultsIterator, result2 error) {
	fake.ExecuteQueryWithMetadataStub = nil
	fake.executeQueryWithMetadataReturns = struct {
		result1 statedb.QueryResultsIterator
		result2 error
	}{result1, result2}
}

func (fake *VersionedDB) ExecuteQueryWithMetadataReturnsOnCall(i int, result1 statedb.QueryResultsIterator, result2 error) {
	fake.ExecuteQueryWithMetadataStub = nil
	if fake.executeQueryWithMetadataReturnsOnCall == nil {
		fake.executeQueryWithMetadataReturnsOnCall = make(map[int]struct {
			result1 statedb.QueryResultsIterator
			result2 error
		})
	}
	fake.executeQueryWithMetadataReturnsOnCall[i] = struct {
		result1 statedb.QueryResultsIterator
		result2 error
	}{result1, result2}
}

func (fake *VersionedDB) ApplyUpdates(batch *statedb.UpdateBatch, height *version.Height) error {
	fake.applyUpdatesMutex.Lock()
	ret, specificReturn := fake.applyUpdatesReturnsOnCall[len(fake.applyUpdatesArgsForCall)]
	fake.applyUpdatesArgsForCall = append(fake.applyUpdatesArgsForCall, struct {
		batch  *statedb.UpdateBatch
		height *version.Height
	}{batch, height})
	fake.recordInvocation("ApplyUpdates", []interface{}{batch, height})
	fake.applyUpdatesMutex.Unlock()
	if fake.ApplyUpdatesStub != nil {
		return fake.ApplyUpdatesStub(batch, height)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.applyUpdatesReturns.result1
}

func (fake *VersionedDB) ApplyUpdatesCallCount() int {
	fake.applyUpdatesMutex.RLock()
	defer fake.applyUpdatesMutex.RUnlock()
	return len(fake.applyUpdatesArgsForCall)
}

func (fake *VersionedDB) ApplyUpdatesArgsForCall(i int) (*statedb.UpdateBatch, *version.Height) {
	fake.applyUpdatesMutex.RLock()
	defer fake.applyUpdatesMutex.RUnlock()
	return fake.applyUpdatesArgsForCall[i].batch, fake.applyUpdatesArgsForCall[i].height
}

func (fake *VersionedDB) ApplyUpdatesReturns(result1 error) {
	fake.ApplyUpdatesStub = nil
	fake.applyUpdatesReturns = struct {
		result1 error
	}{result1}
}

func (fake *VersionedDB) ApplyUpdatesReturnsOnCall(i int, result1 error) {
	fake.ApplyUpdatesStub = nil
	if fake.applyUpdatesReturnsOnCall == nil {
		fake.applyUpdatesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.applyUpdatesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *VersionedDB) GetLatestSavePoint() (*version.Height, error) {
	fake.getLatestSavePointMutex.Lock()
	ret, specificReturn := fake.getLatestSavePointReturnsOnCall[len(fake.getLatestSavePointArgsForCall)]
	fake.getLatestSavePointArgsForCall = append(fake.getLatestSavePointArgsForCall, struct{}{})
	fake.recordInvocation("GetLatestSavePoint", []interface{}{})
	fake.getLatestSavePointMutex.Unlock()
	if fake.GetLatestSavePointStub != nil {
		return fake.GetLatestSavePointStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getLatestSavePointReturns.result1, fake.getLatestSavePointReturns.result2
}

func (fake *VersionedDB) GetLatestSavePointCallCount() int {
	fake.getLatestSavePointMutex.RLock()
	defer fake.getLatestSavePointMutex.RUnlock()
	return len(fake.getLatestSavePointArgsForCall)
}

func (fake *VersionedDB) GetLatestSavePointReturns(result1 *version.Height, result2 error) {
	fake.GetLatestSavePointStub = nil
	fake.getLatestSavePointReturns = struct {
		result1 *version.Height
		result2 error
	}{result1, result2}
}

func (fake *VersionedDB) GetLatestSavePointReturnsOnCall(i int, result1 *version.Height, result2 error) {
	fake.GetLatestSavePointStub = nil
	if fake.getLatestSavePointReturnsOnCall == nil {
		fake.getLatestSavePointReturnsOnCall = make(map[int]struct {
			result1 *version.Height
			result2 error
		})
	}
	fake.getLatestSavePointReturnsOnCall[i] = struct {
		result1 *version.Height
		result2 error
	}{result1, result2}
}

func (fake *VersionedDB) ValidateKeyValue(key string, value []byte) error {
	var valueCopy []byte
	if value != nil {
		valueCopy = make([]byte, len(value))
		copy(valueCopy, value)
	}
	fake.validateKeyValueMutex.Lock()
	ret, specificReturn := fake.validateKeyValueReturnsOnCall[len(fake.validateKeyValueArgsForCall)]
	fake.validateKeyValueArgsForCall = append(fake.validateKeyValueArgsForCall, struct {
		key   string
		value []byte
	}{key, valueCopy})
	fake.recordInvocation("ValidateKeyValue", []interface{}{key, valueCopy})
	fake.validateKeyValueMutex.Unlock()
	if fake.ValidateKeyValueStub != nil {
		return fake.ValidateKeyValueStub(key, value)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.validateKeyValueReturns.result1
}

func (fake *VersionedDB) ValidateKeyValueCallCount() int {
	fake.validateKeyValueMutex.RLock()
	defer fake.validateKeyValueMutex.RUnlock()
	return len(fake.validateKeyValueArgsForCall)
}

func (fake *VersionedDB) ValidateKeyValueArgsForCall(i int) (string, []byte) {
	fake.validateKeyValueMutex.RLock()
	defer fake.validateKeyValueMutex.RUnlock()
	return fake.validateKeyValueArgsForCall[i].key, fake.validateKeyValueArgsForCall[i].value
}

func (fake *VersionedDB) ValidateKeyValueReturns(result1 error) {
	fake.ValidateKeyValueStub = nil
	fake.validateKeyValueReturns = struct {
		result1 error
	}{result1}
}

func (fake *VersionedDB) ValidateKeyValueReturnsOnCall(i int, result1 error) {
	fake.ValidateKeyValueStub = nil
	if fake.validateKeyValueReturnsOnCall == nil {
		fake.validateKeyValueReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateKeyValueReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *VersionedDB) BytesKeySupported() bool {
	fake.bytesKeySupportedMutex.Lock()
	ret, specificReturn := fake.bytesKeySupportedReturnsOnCall[len(fake.bytesKeySupportedArgsForCall)]
	fake.bytesKeySupportedArgsForCall = append(fake.bytesKeySupportedArgsForCall, struct{}{})
	fake.recordInvocation("BytesKeySupported", []interface{}{})
	fake.bytesKeySupportedMutex.Unlock()
	if fake.BytesKeySupportedStub != nil {
		return fake.BytesKeySupportedStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.bytesKeySupportedReturns.result1
}

func (fake *VersionedDB) BytesKeySupportedCallCount() int {
	fake.bytesKeySupportedMutex.RLock()
	defer fake.bytesKeySupportedMutex.RUnlock()
	return len(fake.bytesKeySupportedArgsForCall)
}

func (fake *VersionedDB) BytesKeySupportedReturns(result1 bool) {
	fake.BytesKeySupportedStub = nil
	fake.bytesKeySupportedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *VersionedDB) BytesKeySupportedReturnsOnCall(i int, result1 bool) {
	fake.BytesKeySupportedStub = nil
	if fake.bytesKeySupportedReturnsOnCall == nil {
		fake.bytesKeySupportedReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.bytesKeySupportedReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *VersionedDB) Open() error {
	fake.openMutex.Lock()
	ret, specificReturn := fake.openReturnsOnCall[len(fake.openArgsForCall)]
	fake.openArgsForCall = append(fake.openArgsForCall, struct{}{})
	fake.recordInvocation("Open", []interface{}{})
	fake.openMutex.Unlock()
	if fake.OpenStub != nil {
		return fake.OpenStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.openReturns.result1
}

func (fake *VersionedDB) OpenCallCount() int {
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	return len(fake.openArgsForCall)
}

func (fake *VersionedDB) OpenReturns(result1 error) {
	fake.OpenStub = nil
	fake.openReturns = struct {
		result1 error
	}{result1}
}

func (fake *VersionedDB) OpenReturnsOnCall(i int, result1 error) {
	fake.OpenStub = nil
	if fake.openReturnsOnCall == nil {
		fake.openReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.openReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *VersionedDB) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct{}{})
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		fake.CloseStub()
	}
}

func (fake *VersionedDB) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *VersionedDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getStateMutex.RLock()
	defer fake.getStateMutex.RUnlock()
	fake.getVersionMutex.RLock()
	defer fake.getVersionMutex.RUnlock()
	fake.getStateMultipleKeysMutex.RLock()
	defer fake.getStateMultipleKeysMutex.RUnlock()
	fake.getStateRangeScanIteratorMutex.RLock()
	defer fake.getStateRangeScanIteratorMutex.RUnlock()
	fake.getStateRangeScanIteratorWithMetadataMutex.RLock()
	defer fake.getStateRangeScanIteratorWithMetadataMutex.RUnlock()
	fake.executeQueryMutex.RLock()
	defer fake.executeQueryMutex.RUnlock()
	fake.executeQueryWithMetadataMutex.RLock()
	defer fake.executeQueryWithMetadataMutex.RUnlock()
	fake.applyUpdatesMutex.RLock()
	defer fake.applyUpdatesMutex.RUnlock()
	fake.getLatestSavePointMutex.RLock()
	defer fake.getLatestSavePointMutex.RUnlock()
	fake.validateKeyValueMutex.RLock()
	defer fake.validateKeyValueMutex.RUnlock()
	fake.bytesKeySupportedMutex.RLock()
	defer fake.bytesKeySupportedMutex.RUnlock()
	fake.openMutex.RLock()
	defer fake.openMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *VersionedDB) recordInvocation(key string, args []interface{}) {
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

var _ statedb.VersionedDB = new(VersionedDB)