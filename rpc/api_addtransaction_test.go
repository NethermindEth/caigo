package rpc

import (
	_ "embed"
	"testing"

	"github.com/dontpanicdao/caigo/rpc/types"
)

// TestAddDeployTransaction tests AddDeployTransaction
func TestAddDeployTransaction(t *testing.T) {
	_ = beforeEach(t)

	type testSetType struct {
		BroadcastedDeployTxn    types.BroadcastedDeployTxn
		ExpectedTransactionHash string
		ExpectedContractAddress string
	}
	testSet := map[string][]testSetType{
		"mock":    {},
		"testnet": {},
		"devnet":  {},
		"mainnet": {},
	}[testEnv]

	for _, test := range testSet {
		_ = test
	}
}

// TestAddDeclareTransaction tests AddDeclareTransaction
func TestAddDeclareTransaction(t *testing.T) {
	_ = beforeEach(t)

	type testSetType struct {
		BroadcastedDeclareTxn   types.BroadcastedDeclareTxn
		ExpectedTransactionHash string
		ExpectedClassHash       string
	}
	testSet := map[string][]testSetType{
		"mock":    {},
		"testnet": {},
		// TODO: add tests for mainnet when possible or when figure out how to
		// create a white-listed contract.
		"mainnet": {},
	}[testEnv]

	for _, test := range testSet {
		_ = test
	}
}

// TestAddInvokeTransaction tests AddInvokeTransaction
func TestAddInvokeTransaction(t *testing.T) {
	_ = beforeEach(t)

	type testSetType struct {
		BroadcastedInvokeTxn    types.BroadcastedInvokeTxn
		ExpectedTransactionHash string
	}
	testSet := map[string][]testSetType{
		"mock":    {},
		"testnet": {},
		"mainnet": {},
	}[testEnv]

	for _, test := range testSet {
		_ = test
	}
}
