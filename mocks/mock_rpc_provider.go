// Code generated by MockGen. DO NOT EDIT.
// Source: provider.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	felt "github.com/NethermindEth/juno/core/felt"
	rpc "github.com/NethermindEth/starknet.go/rpc"
	gomock "github.com/golang/mock/gomock"
)

// MockRpcProvider is a mock of RpcProvider interface.
type MockRpcProvider struct {
	ctrl     *gomock.Controller
	recorder *MockRpcProviderMockRecorder
}

// MockRpcProviderMockRecorder is the mock recorder for MockRpcProvider.
type MockRpcProviderMockRecorder struct {
	mock *MockRpcProvider
}

// NewMockRpcProvider returns a new instance of MockRpcProvider.
//
// Parameters:
// - ctrl: The gomock.Controller used for creating the mock.
//
// Returns:
// - *MockRpcProvider: The newly created instance of MockRpcProvider.
func NewMockRpcProvider(ctrl *gomock.Controller) *MockRpcProvider {
	mock := &MockRpcProvider{ctrl: ctrl}
	mock.recorder = &MockRpcProviderMockRecorder{mock}
	return mock
}

// EXPECT returns a pointer to a MockRpcProviderMockRecorder (allows the caller to indicate expected use).
//
// No parameters.
// Returns a pointer to a MockRpcProviderMockRecorder.
func (m *MockRpcProvider) EXPECT() *MockRpcProviderMockRecorder {
	return m.recorder
}

// AddDeclareTransaction description of the Go function (base method).
//
// AddDeclareTransaction adds a new declaration transaction to the mock RPC provider.
// It takes a context and an rpc.AddDeclareTxnInput as input parameters. The function
// returns an *rpc.AddDeclareTransactionResponse and an error.
func (m *MockRpcProvider) AddDeclareTransaction(ctx context.Context, declareTransaction rpc.AddDeclareTxnInput) (*rpc.AddDeclareTransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDeclareTransaction", ctx, declareTransaction)
	ret0, _ := ret[0].(*rpc.AddDeclareTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddDeclareTransaction description of the Go function.
//
// Parameters:
// - ctx: the context.
// - declareTransaction: the declare transaction.
// 
// Returns:
// - *gomock.Call: the mock recorder.
func (mr *MockRpcProviderMockRecorder) AddDeclareTransaction(ctx, declareTransaction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDeclareTransaction", reflect.TypeOf((*MockRpcProvider)(nil).AddDeclareTransaction), ctx, declareTransaction)
}

// AddDeployAccountTransaction adds a deploy account transaction to the MockRpcProvider.
//
// ctx is the context.Context used for the function call.
// deployAccountTransaction is the rpc.DeployAccountTxn to be added.
// Returns *rpc.AddDeployAccountTransactionResponse and error.
func (m *MockRpcProvider) AddDeployAccountTransaction(ctx context.Context, deployAccountTransaction rpc.DeployAccountTxn) (*rpc.AddDeployAccountTransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDeployAccountTransaction", ctx, deployAccountTransaction)
	ret0, _ := ret[0].(*rpc.AddDeployAccountTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddDeployAccountTransaction description of the expected call of AddDeployAccountTransaction.
//
// Parameters:
// - ctx: The context.
// - deployAccountTransaction: The deploy account transaction.
//
// Returns:
// - The gomock.Call object.
func (mr *MockRpcProviderMockRecorder) AddDeployAccountTransaction(ctx, deployAccountTransaction interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDeployAccountTransaction", reflect.TypeOf((*MockRpcProvider)(nil).AddDeployAccountTransaction), ctx, deployAccountTransaction)
}

// AddInvokeTransaction is a function that adds an invoke transaction.
//
// ctx: the context object.
// invokeTxn: the invoke transaction object.
// Returns an AddInvokeTransactionResponse and an error.
func (m *MockRpcProvider) AddInvokeTransaction(ctx context.Context, invokeTxn rpc.InvokeTxnV1) (*rpc.AddInvokeTransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddInvokeTransaction", ctx, invokeTxn)
	ret0, _ := ret[0].(*rpc.AddInvokeTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddInvokeTransaction description of an expected call of AddInvokeTransaction.
//
// Parameters:
// - ctx: The context.
// - invokeTxn: The invoke transaction.
// Return type: *gomock.Call.
func (mr *MockRpcProviderMockRecorder) AddInvokeTransaction(ctx, invokeTxn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddInvokeTransaction", reflect.TypeOf((*MockRpcProvider)(nil).AddInvokeTransaction), ctx, invokeTxn)
}

// BlockHashAndNumber returns the BlockHashAndNumberOutput and possible error.
//
// It takes a context.Context as a parameter.
// It returns *rpc.BlockHashAndNumberOutput and error.
func (m *MockRpcProvider) BlockHashAndNumber(ctx context.Context) (*rpc.BlockHashAndNumberOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockHashAndNumber", ctx)
	ret0, _ := ret[0].(*rpc.BlockHashAndNumberOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockHashAndNumber description of an expected call of BlockHashAndNumber.
//
// ctx is the context parameter.
// It returns a *gomock.Call type.
func (mr *MockRpcProviderMockRecorder) BlockHashAndNumber(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockHashAndNumber", reflect.TypeOf((*MockRpcProvider)(nil).BlockHashAndNumber), ctx)
}

// BlockNumber returns the block number.
//
// ctx - the context object.
// uint64 - the block number.
// error - an error if any occurred.
func (m *MockRpcProvider) BlockNumber(ctx context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockNumber", ctx)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockNumber is a function that records a call to the BlockNumber method of the MockRpcProvider interface.
//
// ctx is the context object passed as a parameter to the BlockNumber method.
// The function returns a *gomock.Call object.
func (mr *MockRpcProviderMockRecorder) BlockNumber(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockNumber", reflect.TypeOf((*MockRpcProvider)(nil).BlockNumber), ctx)
}

// BlockTransactionCount returns the number of transactions in a block.
//
// ctx: The context for the function.
// blockID: The ID of the block.
// Returns:
//   uint64: The number of transactions in the block.
//   error: An error if the operation fails.
func (m *MockRpcProvider) BlockTransactionCount(ctx context.Context, blockID rpc.BlockID) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockTransactionCount", ctx, blockID)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockTransactionCount description of an expected call of BlockTransactionCount.
//
// ctx: The context.
// blockID: The block ID.
// Return type: *gomock.Call.
func (mr *MockRpcProviderMockRecorder) BlockTransactionCount(ctx, blockID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockTransactionCount", reflect.TypeOf((*MockRpcProvider)(nil).BlockTransactionCount), ctx, blockID)
}

// BlockWithTxHashes is a function that retrieves a block with transaction hashes.
//
// It takes in two parameters: 
// - ctx: the context.Context object used for cancellation and timeouts.
// - blockID: the rpc.BlockID object representing the ID of the block to retrieve.
//
// It returns two values:
// - interface{}: the block with transaction hashes.
// - error: an error object, if any error occurs during the retrieval process.
func (m *MockRpcProvider) BlockWithTxHashes(ctx context.Context, blockID rpc.BlockID) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockWithTxHashes", ctx, blockID)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockWithTxHashes description of an expected call of BlockWithTxHashes.
//
// Takes a context and blockID as parameters and returns a *gomock.Call.
func (mr *MockRpcProviderMockRecorder) BlockWithTxHashes(ctx, blockID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockWithTxHashes", reflect.TypeOf((*MockRpcProvider)(nil).BlockWithTxHashes), ctx, blockID)
}

// BlockWithTxs description of the Go function.
//
// This function takes a context and a block ID as parameters and returns an interface and an error.
func (m *MockRpcProvider) BlockWithTxs(ctx context.Context, blockID rpc.BlockID) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockWithTxs", ctx, blockID)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockWithTxs description of an expected call of BlockWithTxs.
//
// - ctx: description of the ctx parameter.
// - blockID: description of the blockID parameter.
// Return type: *gomock.Call.
func (mr *MockRpcProviderMockRecorder) BlockWithTxs(ctx, blockID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockWithTxs", reflect.TypeOf((*MockRpcProvider)(nil).BlockWithTxs), ctx, blockID)
}

// Call is a method of the MockRpcProvider struct that handles a function call.
//
// It takes in a context.Context object, a rpc.FunctionCall object, and a rpc.BlockID object as parameters.
// It returns a slice of pointers to felt.Felt objects and an error object.
func (m *MockRpcProvider) Call(ctx context.Context, call rpc.FunctionCall, block rpc.BlockID) ([]*felt.Felt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Call", ctx, call, block)
	ret0, _ := ret[0].([]*felt.Felt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Call mocks the Call method of the RpcProvider interface.
//
// It takes in three parameters: ctx, call, and block, all of type interface{}.
// It returns a *gomock.Call object.
func (mr *MockRpcProviderMockRecorder) Call(ctx, call, block interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Call", reflect.TypeOf((*MockRpcProvider)(nil).Call), ctx, call, block)
}

// ChainID returns the chain ID of the RPC provider.
//
// ctx - the context for the function.
// string - the chain ID.
// error - an error if the chain ID cannot be retrieved.
func (m *MockRpcProvider) ChainID(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChainID", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChainID description of an expected call of ChainID.
//
// ctx is the context parameter used for ...
// The function returns a *gomock.Call type.
func (mr *MockRpcProviderMockRecorder) ChainID(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChainID", reflect.TypeOf((*MockRpcProvider)(nil).ChainID), ctx)
}

// Class description of the Go function.
//
// Class retrieves the rpc.ClassOutput for the given blockID and classHash.
// It returns the rpc.ClassOutput and an error, if any.
func (m *MockRpcProvider) Class(ctx context.Context, blockID rpc.BlockID, classHash *felt.Felt) (rpc.ClassOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Class", ctx, blockID, classHash)
	ret0, _ := ret[0].(rpc.ClassOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Class description of an expected call of Class.
//
// ctx - The context.
// blockID - The block ID.
// classHash - The class hash.
// Returns *gomock.Call.
func (mr *MockRpcProviderMockRecorder) Class(ctx, blockID, classHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Class", reflect.TypeOf((*MockRpcProvider)(nil).Class), ctx, blockID, classHash)
}

// ClassAt description of the Go function.
//
// ClassAt is a method of the MockRpcProvider struct that takes in the following parameters:
// - ctx: a context.Context object
// - blockID: an rpc.BlockID object
// - contractAddress: a pointer to a felt.Felt object
//
// It returns an rpc.ClassOutput object and an error.
func (m *MockRpcProvider) ClassAt(ctx context.Context, blockID rpc.BlockID, contractAddress *felt.Felt) (rpc.ClassOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClassAt", ctx, blockID, contractAddress)
	ret0, _ := ret[0].(rpc.ClassOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClassAt description of the Go function.
//
// Parameters:
// - ctx: The context.
// - blockID: The block ID.
// - contractAddress: The contract address.
//
// Returns:
// - *gomock.Call: The gomock.Call object.
func (mr *MockRpcProviderMockRecorder) ClassAt(ctx, blockID, contractAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClassAt", reflect.TypeOf((*MockRpcProvider)(nil).ClassAt), ctx, blockID, contractAddress)
}

// ClassHashAt description of the Go function.
//
// ClassHashAt is a function that takes in a context, blockID, and contractAddress 
// parameters and returns a felt.Felt pointer and an error. It is part of the 
// MockRpcProvider struct.
//
// Parameters:
//   - ctx: The context.Context object representing the context of the function.
//   - blockID: The rpc.BlockID object representing the block ID.
//   - contractAddress: The felt.Felt object representing the contract address.
//
// Returns:
//   - ret0: The felt.Felt pointer returned by the function.
//   - ret1: The error returned by the function.
func (m *MockRpcProvider) ClassHashAt(ctx context.Context, blockID rpc.BlockID, contractAddress *felt.Felt) (*felt.Felt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClassHashAt", ctx, blockID, contractAddress)
	ret0, _ := ret[0].(*felt.Felt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClassHashAt description of an expected call of ClassHashAt.
//
// Parameters:
// - ctx: The context.
// - blockID: The block ID.
// - contractAddress: The contract address.
//
// Returns:
// - *gomock.Call: The gomock call.
func (mr *MockRpcProviderMockRecorder) ClassHashAt(ctx, blockID, contractAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClassHashAt", reflect.TypeOf((*MockRpcProvider)(nil).ClassHashAt), ctx, blockID, contractAddress)
}

// EstimateFee is a function that estimates the fees for a given set of requests and block ID.
//
// ctx: The context.Context object for the function.
// requests: An array of rpc.EstimateFeeInput objects representing the requests for which fees need to be estimated.
// blockID: The rpc.BlockID object representing the block ID for which fees need to be estimated.
// []rpc.FeeEstimate: An array of rpc.FeeEstimate objects representing the estimated fees.
// error: An error object representing any error that occurred during the estimation process.
func (m *MockRpcProvider) EstimateFee(ctx context.Context, requests []rpc.EstimateFeeInput, blockID rpc.BlockID) ([]rpc.FeeEstimate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EstimateFee", ctx, requests, blockID)
	ret0, _ := ret[0].([]rpc.FeeEstimate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateFee description of an expected call of EstimateFee.
//
// - ctx: The context.
// - requests: The requests.
// - blockID: The block ID.
// Return type: *gomock.Call.
func (mr *MockRpcProviderMockRecorder) EstimateFee(ctx, requests, blockID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateFee", reflect.TypeOf((*MockRpcProvider)(nil).EstimateFee), ctx, requests, blockID)
}

// EstimateMessageFee is a function that estimates the fee for a given message in the MockRpcProvider.
//
// ctx: The context in which the function is executed.
// msg: The message from L1.
// blockID: The ID of the block.
// Returns: The estimated fee for the message and an error object.
func (m *MockRpcProvider) EstimateMessageFee(ctx context.Context, msg rpc.MsgFromL1, blockID rpc.BlockID) (*rpc.FeeEstimate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EstimateMessageFee", ctx, msg, blockID)
	ret0, _ := ret[0].(*rpc.FeeEstimate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateMessageFee description of an expected call of EstimateMessageFee.
//
// Parameters:
// - ctx: the context.
// - msg: the message.
// - blockID: the block ID.
// Return type:
// - *gomock.Call: the mock call.
func (mr *MockRpcProviderMockRecorder) EstimateMessageFee(ctx, msg, blockID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateMessageFee", reflect.TypeOf((*MockRpcProvider)(nil).EstimateMessageFee), ctx, msg, blockID)
}

// Events returns the event chunk and error for the MockRpcProvider.
//
// It takes the context and input parameters and returns *rpc.EventChunk and error.
func (m *MockRpcProvider) Events(ctx context.Context, input rpc.EventsInput) (*rpc.EventChunk, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Events", ctx, input)
	ret0, _ := ret[0].(*rpc.EventChunk)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Events description of an expected call of Events.
//
// ctx: the context parameter.
// input: the input parameter.
// Return type: a *gomock.Call.
func (mr *MockRpcProviderMockRecorder) Events(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Events", reflect.TypeOf((*MockRpcProvider)(nil).Events), ctx, input)
}

// Nonce is a function that retrieves the nonce for a given block ID and contract address.
//
// ctx: The context for the function.
// blockID: The block ID for which to retrieve the nonce.
// contractAddress: The contract address for which to retrieve the nonce.
// Returns a pointer to a string and an error.
func (m *MockRpcProvider) Nonce(ctx context.Context, blockID rpc.BlockID, contractAddress *felt.Felt) (*string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Nonce", ctx, blockID, contractAddress)
	ret0, _ := ret[0].(*string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Nonce returns a gomock.Call object that records the call to the Nonce method of the MockRpcProvider interface.
//
// The Nonce method takes three parameters:
// - ctx: the context parameter of type interface{}.
// - blockID: the blockID parameter of type interface{}.
// - contractAddress: the contractAddress parameter of type interface{}.
//
// The function returns a *gomock.Call object.
func (mr *MockRpcProviderMockRecorder) Nonce(ctx, blockID, contractAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nonce", reflect.TypeOf((*MockRpcProvider)(nil).Nonce), ctx, blockID, contractAddress)
}

// SimulateTransactions is a function that simulates transactions.
//
// It takes the following parameters:
// - ctx: the context.Context object for controlling the execution flow.
// - blockID: the rpc.BlockID object representing the ID of the block.
// - txns: a slice of rpc.Transaction objects representing the transactions.
// - simulationFlags: a slice of rpc.SimulationFlag objects representing the simulation flags.
//
// It returns a slice of rpc.SimulatedTransaction objects, which represents the simulated transactions,
// and an error object in case of any error during the simulation.
func (m *MockRpcProvider) SimulateTransactions(ctx context.Context, blockID rpc.BlockID, txns []rpc.Transaction, simulationFlags []rpc.SimulationFlag) ([]rpc.SimulatedTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SimulateTransactions", ctx, blockID, txns, simulationFlags)
	ret0, _ := ret[0].([]rpc.SimulatedTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SimulateTransactions description of an expected call of SimulateTransactions.
//
// Parameters:
//   - ctx: The context of the function.
//   - blockID: The ID of the block.
//   - txns: The transactions to be simulated.
//   - simulationFlags: The flags for simulation.
// Return type:
//   - *gomock.Call: The call object.
func (mr *MockRpcProviderMockRecorder) SimulateTransactions(ctx, blockID, txns, simulationFlags interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SimulateTransactions", reflect.TypeOf((*MockRpcProvider)(nil).SimulateTransactions), ctx, blockID, txns, simulationFlags)
}

// StateUpdate description of the Go function.
//
// StateUpdate is a method that updates the state based on the given block ID.
// It takes a context and a block ID as parameters and returns a pointer to a
// StateUpdateOutput and an error.
func (m *MockRpcProvider) StateUpdate(ctx context.Context, blockID rpc.BlockID) (*rpc.StateUpdateOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StateUpdate", ctx, blockID)
	ret0, _ := ret[0].(*rpc.StateUpdateOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StateUpdate description of an expected call of StateUpdate.
//
// ctx - description of the ctx parameter.
// blockID - description of the blockID parameter.
// Return type - *gomock.Call.
func (mr *MockRpcProviderMockRecorder) StateUpdate(ctx, blockID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateUpdate", reflect.TypeOf((*MockRpcProvider)(nil).StateUpdate), ctx, blockID)
}

// StorageAt description of the base method.
//
// StorageAt retrieves the storage data at a specific key for a given contract address and block ID.
// It takes the following parameters:
// - ctx: The context.Context object for the RPC call.
// - contractAddress: The contract address as a *felt.Felt object.
// - key: The key as a string.
// - blockID: The block ID as a rpc.BlockID object.
// It returns a string representing the storage data and an error if any.
func (m *MockRpcProvider) StorageAt(ctx context.Context, contractAddress *felt.Felt, key string, blockID rpc.BlockID) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StorageAt", ctx, contractAddress, key, blockID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StorageAt description of an expected call of StorageAt.
//
// ctx: description of the parameter.
// contractAddress: description of the parameter.
// key: description of the parameter.
// blockID: description of the parameter.
// Return type: *gomock.Call.
func (mr *MockRpcProviderMockRecorder) StorageAt(ctx, contractAddress, key, blockID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StorageAt", reflect.TypeOf((*MockRpcProvider)(nil).StorageAt), ctx, contractAddress, key, blockID)
}

// Syncing returns the sync status and an error.
//
// It takes a context.Context as a parameter.
// It returns a *rpc.SyncStatus and an error.
func (m *MockRpcProvider) Syncing(ctx context.Context) (*rpc.SyncStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Syncing", ctx)
	ret0, _ := ret[0].(*rpc.SyncStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Syncing description of an expected call of Syncing.
//
// ctx interface{}.
// *gomock.Call.
func (mr *MockRpcProviderMockRecorder) Syncing(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Syncing", reflect.TypeOf((*MockRpcProvider)(nil).Syncing), ctx)
}

// TraceBlockTransactions returns the traces of transactions in a given block.
//
// ctx: the context.Context object.
// blockHash: the hash of the block to trace.
// []rpc.Trace: an array of rpc.Trace objects representing the traces of the transactions in the block.
// error: an error object, if any error occurred during the trace.
func (m *MockRpcProvider) TraceBlockTransactions(ctx context.Context, blockHash *felt.Felt) ([]rpc.Trace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TraceBlockTransactions", ctx, blockHash)
	ret0, _ := ret[0].([]rpc.Trace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TraceBlockTransactions description of an expected call of TraceBlockTransactions.
//
// ctx, blockHash description of its parameter(s).
// *gomock.Call description of its return type(s).
func (mr *MockRpcProviderMockRecorder) TraceBlockTransactions(ctx, blockHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TraceBlockTransactions", reflect.TypeOf((*MockRpcProvider)(nil).TraceBlockTransactions), ctx, blockHash)
}

// TransactionByBlockIdAndIndex returns a transaction by block ID and index.
//
// ctx: the context.Context object for the function.
// blockID: the ID of the block.
// index: the index of the transaction within the block.
// Returns:
// - the rpc.Transaction object representing the transaction.
// - an error if there was an error retrieving the transaction.
func (m *MockRpcProvider) TransactionByBlockIdAndIndex(ctx context.Context, blockID rpc.BlockID, index uint64) (rpc.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionByBlockIdAndIndex", ctx, blockID, index)
	ret0, _ := ret[0].(rpc.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionByBlockIdAndIndex description of an expected call of TransactionByBlockIdAndIndex.
//
// Parameters:
//   - ctx: The context.
//   - blockID: The block ID.
//   - index: The index.
// Return type:
//   - *gomock.Call
func (mr *MockRpcProviderMockRecorder) TransactionByBlockIdAndIndex(ctx, blockID, index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionByBlockIdAndIndex", reflect.TypeOf((*MockRpcProvider)(nil).TransactionByBlockIdAndIndex), ctx, blockID, index)
}

// TransactionByHash is a function that retrieves a transaction by its hash.
//
// It takes the following parameter(s):
// - ctx: the context for the function execution.
// - hash: the hash of the transaction to retrieve.
//
// It returns a value of type rpc.Transaction representing the retrieved transaction,
// and an error indicating any error that occurred during the retrieval process.
func (m *MockRpcProvider) TransactionByHash(ctx context.Context, hash *felt.Felt) (rpc.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionByHash", ctx, hash)
	ret0, _ := ret[0].(rpc.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionByHash description of an expected call of TransactionByHash.
//
// ctx represents the context of the function call.
// hash is the hash value used to find the transaction.
// Returns a *gomock.Call object.
func (mr *MockRpcProviderMockRecorder) TransactionByHash(ctx, hash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionByHash", reflect.TypeOf((*MockRpcProvider)(nil).TransactionByHash), ctx, hash)
}

// TransactionReceipt description of the base method.
//
// TransactionReceipt retrieves the transaction receipt for a given transaction hash.
// It takes a context.Context and a *felt.Felt as parameters.
// It returns an rpc.TransactionReceipt and an error.
func (m *MockRpcProvider) TransactionReceipt(ctx context.Context, transactionHash *felt.Felt) (rpc.TransactionReceipt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionReceipt", ctx, transactionHash)
	ret0, _ := ret[0].(rpc.TransactionReceipt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionReceipt returns a mock call for the TransactionReceipt method of the MockRpcProviderMockRecorder type.
//
// It takes two parameters:
//   - ctx: the context for the transaction
//   - transactionHash: the hash of the transaction
// 
// It returns a *gomock.Call object.
func (mr *MockRpcProviderMockRecorder) TransactionReceipt(ctx, transactionHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionReceipt", reflect.TypeOf((*MockRpcProvider)(nil).TransactionReceipt), ctx, transactionHash)
}

// TransactionTrace is a function that retrieves the transaction trace for a given transaction hash.
//
// It takes in a context and a transaction hash as parameters and returns a rpc.TxnTrace and an error.
func (m *MockRpcProvider) TransactionTrace(ctx context.Context, transactionHash *felt.Felt) (rpc.TxnTrace, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionTrace", ctx, transactionHash)
	ret0, _ := ret[0].(rpc.TxnTrace)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionTrace description of an expected call of TransactionTrace.
//
// ctx: The context.
// transactionHash: The transaction hash.
// Returns *gomock.Call.
func (mr *MockRpcProviderMockRecorder) TransactionTrace(ctx, transactionHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionTrace", reflect.TypeOf((*MockRpcProvider)(nil).TransactionTrace), ctx, transactionHash)
}
