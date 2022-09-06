package rpc

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TestBlockNumber tests BlockNumber and check the returned value is strictly positive
func TestBlockNumber(t *testing.T) {
	testConfig := beforeEach(t)

	type testSetType struct{}

	testSet := map[string][]testSetType{
		"mock":    {},
		"testnet": {{}},
		"mainnet": {{}},
		"devnet":  {},
	}[testEnv]

	for range testSet {
		spy := NewSpy(testConfig.client.c)
		testConfig.client.c = spy
		blockNumber, err := testConfig.client.BlockNumber(context.Background())
		if err != nil {
			t.Fatal("BlockWithTxHashes match the expected error:", err)
		}
		if diff, err := spy.Compare(blockNumber, false); err != nil || diff != "FullMatch" {
			t.Fatal("expecting to match", err)
		}
		if blockNumber <= 300000 {
			t.Fatal("Block number should be > 3000, instead: ", blockNumber)
		}
	}
}

// TestBlockHashAndNumber tests BlockHashAndNumber and check the returned value is strictly positive
func TestBlockHashAndNumber(t *testing.T) {
	testConfig := beforeEach(t)

	type testSetType struct{}

	testSet := map[string][]testSetType{
		"mock":    {},
		"testnet": {{}},
		"mainnet": {{}},
		"devnet":  {},
	}[testEnv]

	for range testSet {
		spy := NewSpy(testConfig.client.c)
		testConfig.client.c = spy
		blockHashAndNumber, err := testConfig.client.BlockHashAndNumber(context.Background())
		if err != nil {
			t.Fatal("BlockHashAndNumber match the expected error:", err)
		}
		if diff, err := spy.Compare(blockHashAndNumber, false); err != nil || diff != "FullMatch" {
			t.Fatal("expecting to match", err)
		}
		if blockHashAndNumber.BlockNumber < 3000 {
			t.Fatal("Block number should be > 3000, instead: ", blockHashAndNumber.BlockNumber)
		}
		if !strings.HasPrefix(blockHashAndNumber.BlockHash, "0x") {
			t.Fatal("current block hash should return a string starting with 0x")
		}
	}
}

// TestPendingBlockWithTxHashes tests TestPendingBlockWithTxHashes
func TestPendingBlockWithTxHashes(t *testing.T) {
	testConfig := beforeEach(t)

	type testSetType struct {
	}
	testSet := map[string][]testSetType{
		"mock":    {},
		"testnet": {},
		"mainnet": {},
		"devnet":  {},
	}[testEnv]

	for range testSet {
		pending, err := testConfig.client.BlockWithTxHashes(context.Background(), WithBlockTag("pending"))
		if err == nil || !strings.Contains(err.Error(), "Pending data not supported in this configuration") {
			t.Fatal("PendingBlockWithTxHashes should not yet be supported")
		}
		if _, ok := pending.(PendingBlockWithTxHashes); !ok {
			t.Fatalf("expecting PendingBlockWithTxs, instead %T", pending)
		}
	}
}

// TestBlockWithTxHashes tests TestBlockWithTxHashes
func TestBlockWithTxHashes(t *testing.T) {
	testConfig := beforeEach(t)

	type testSetType struct {
		BlockID                   BlockID
		ExpectedError             error
		ExpectedBlockWithTxHashes BlockWithTxHashes
	}

	testSet := map[string][]testSetType{
		"mock": {},
		"testnet": {
			{
				BlockID:       WithBlockTag("latest"),
				ExpectedError: nil,
			},
			{
				BlockID:       WithBlockTag("error"),
				ExpectedError: errInvalidBlockID,
			},
			{
				BlockID:                   WithBlockHash(BlockHash("0x6c2fe3db009a2e008c2d65fca14204f3405cb74742fcf685f02473acaf70c72")),
				ExpectedError:             nil,
				ExpectedBlockWithTxHashes: blockGoerli310370,
			},
			{
				BlockID:                   WithBlockNumber(BlockNumber(310370)),
				ExpectedError:             nil,
				ExpectedBlockWithTxHashes: blockGoerli310370,
			},
		},
		"mainnet": {},
	}[testEnv]

	for _, test := range testSet {
		spy := NewSpy(testConfig.client.c)
		testConfig.client.c = spy
		blockWithTxHashesInterface, err := testConfig.client.BlockWithTxHashes(context.Background(), test.BlockID)
		if err != test.ExpectedError {
			t.Fatal("BlockWithTxHashes match the expected error:", err)
		}
		if test.ExpectedError != nil && blockWithTxHashesInterface == nil {
			continue
		}
		blockWithTxHashes, ok := blockWithTxHashesInterface.(*BlockWithTxHashes)
		if !ok {
			t.Fatalf("expecting BlockWithTxHashes, instead %T", blockWithTxHashesInterface)
		}
		if diff, err := spy.Compare(blockWithTxHashes, false); err != nil || diff != "FullMatch" {
			t.Fatal("expecting to match", err)
		}
		if !strings.HasPrefix(string(blockWithTxHashes.BlockHash), "0x") {
			t.Fatal("Block Hash should start with \"0x\", instead", blockWithTxHashes.BlockHash)
		}

		if len(blockWithTxHashes.Transactions) == 0 {
			t.Fatal("the number of transaction should not be 0")
		}
		if test.ExpectedBlockWithTxHashes.BlockHash == "" {
			continue
		}
		if !cmp.Equal(test.ExpectedBlockWithTxHashes, *blockWithTxHashes) {
			t.Fatalf("the expected transaction blocks to match, instead: %s", cmp.Diff(test.ExpectedBlockWithTxHashes, blockWithTxHashes))
		}
	}
}

// TestBlockWithTxsAndInvokeTXNV0 tests block with Invoke TXN V0
func TestBlockWithTxsAndInvokeTXNV0(t *testing.T) {
	testConfig := beforeEach(t)

	type testSetType struct {
		BlockID                     BlockID
		ExpectedError               error
		LookupTxnPositionInOriginal int
		LookupTxnPositionInExpected int
		ExpectedBlockWithTxs        BlockWithTxs
	}
	testSet := map[string][]testSetType{
		"mock": {},
		"testnet": {
			{
				BlockID:       WithBlockTag("latest"),
				ExpectedError: nil,
			},
			{
				BlockID:       WithBlockTag("error"),
				ExpectedError: errInvalidBlockID,
			},
			{
				BlockID:              WithBlockHash(BlockHash("0x6c2fe3db009a2e008c2d65fca14204f3405cb74742fcf685f02473acaf70c72")),
				ExpectedError:        nil,
				ExpectedBlockWithTxs: fullBlockGoerli310370,
			},
			{
				BlockID:              WithBlockNumber(BlockNumber(310370)),
				ExpectedError:        nil,
				ExpectedBlockWithTxs: fullBlockGoerli310370,
			},
		},
		"mainnet": {},
	}[testEnv]

	for _, test := range testSet {
		spy := NewSpy(testConfig.client.c)
		testConfig.client.c = spy
		blockWithTxsInterface, err := testConfig.client.BlockWithTxs(context.Background(), test.BlockID)
		if err != test.ExpectedError {
			t.Fatal("BlockWithTxHashes match the expected error:", err)
		}
		if test.ExpectedError != nil && blockWithTxsInterface == nil {
			continue
		}
		blockWithTxs, ok := blockWithTxsInterface.(*BlockWithTxs)
		if !ok {
			t.Fatalf("expecting BlockWithTxs, instead %T", blockWithTxsInterface)
		}
		diff, err := spy.Compare(blockWithTxs, false)
		if err != nil {
			t.Fatal("expecting to match", err)
		}
		if diff != "FullMatch" {
			spy.Compare(blockWithTxs, true)
			t.Fatal("structure expecting to be FullMatch, instead", diff)
		}
		if !strings.HasPrefix(string(blockWithTxs.BlockHash), "0x") {
			t.Fatal("Block Hash should start with \"0x\", instead", blockWithTxs.BlockHash)
		}

		if len(blockWithTxs.Transactions) == 0 {
			t.Fatal("the number of transaction should not be 0")
		}
		if test.ExpectedBlockWithTxs.BlockHash == "" {
			continue
		}
		if !cmp.Equal(test.ExpectedBlockWithTxs.Transactions[test.LookupTxnPositionInExpected], blockWithTxs.Transactions[test.LookupTxnPositionInOriginal]) {
			t.Fatalf("the expected transaction blocks to match, instead: %s", cmp.Diff(test.ExpectedBlockWithTxs.Transactions[test.LookupTxnPositionInExpected], blockWithTxs.Transactions[test.LookupTxnPositionInOriginal]))
		}
	}
}

// TestBlockWithTxsAndDeployOrDeclare tests BlockWithTxs with Deploy or Declare TXN
func TestBlockWithTxsAndDeployOrDeclare(t *testing.T) {
	testConfig := beforeEach(t)

	type testSetType struct {
		BlockID                     BlockID
		ExpectedError               error
		LookupTxnPositionInOriginal int
		LookupTxnPositionInExpected int
		ExpectedBlockWithTxs        BlockWithTxs
	}
	testSet := map[string][]testSetType{
		"mock": {},
		"testnet": {
			{
				BlockID:       WithBlockTag("latest"),
				ExpectedError: nil,
			},
			{
				BlockID:       WithBlockTag("error"),
				ExpectedError: errInvalidBlockID,
			},
			{
				BlockID:                     WithBlockHash(BlockHash("0x424fba26a7760b63895abe0c366c2d254cb47090c6f9e91ba2b3fa0824d4fc9")),
				ExpectedError:               nil,
				LookupTxnPositionInOriginal: 14,
				LookupTxnPositionInExpected: 0,
				ExpectedBlockWithTxs:        fullBlockGoerli310843,
			},
			{
				BlockID:                     WithBlockNumber(BlockNumber(310843)),
				ExpectedError:               nil,
				LookupTxnPositionInOriginal: 14,
				LookupTxnPositionInExpected: 0,
				ExpectedBlockWithTxs:        fullBlockGoerli310843,
			},
			{
				BlockID:                     WithBlockNumber(BlockNumber(300114)),
				ExpectedError:               nil,
				LookupTxnPositionInOriginal: 3,
				LookupTxnPositionInExpected: 0,
				ExpectedBlockWithTxs:        fullBlockGoerli300114,
			},
		},
		"mainnet": {},
	}[testEnv]

	for _, test := range testSet {
		spy := NewSpy(testConfig.client.c)
		testConfig.client.c = spy
		blockWithTxsInterface, err := testConfig.client.BlockWithTxs(context.Background(), test.BlockID)
		if err != test.ExpectedError {
			t.Fatal("BlockWithTxHashes match the expected error:", err)
		}
		if test.ExpectedError != nil && blockWithTxsInterface == nil {
			continue
		}
		blockWithTxs, ok := blockWithTxsInterface.(*BlockWithTxs)
		if !ok {
			t.Fatalf("expecting BlockWithTxs, instead %T", blockWithTxsInterface)
		}
		diff, err := spy.Compare(blockWithTxs, false)
		if err != nil {
			t.Fatal("expecting to match", err)
		}
		if diff != "FullMatch" {
			spy.Compare(blockWithTxs, true)
			t.Fatal("structure expecting to be FullMatch, instead", diff)
		}
		if !strings.HasPrefix(string(blockWithTxs.BlockHash), "0x") {
			t.Fatal("Block Hash should start with \"0x\", instead", blockWithTxs.BlockHash)
		}

		if len(blockWithTxs.Transactions) == 0 {
			t.Fatal("the number of transaction should not be 0")
		}
		if test.ExpectedBlockWithTxs.BlockHash == "" {
			continue
		}
		if !cmp.Equal(test.ExpectedBlockWithTxs.Transactions[test.LookupTxnPositionInExpected], blockWithTxs.Transactions[test.LookupTxnPositionInOriginal]) {
			t.Fatalf("the expected transaction blocks to match, instead: %s", cmp.Diff(test.ExpectedBlockWithTxs.Transactions[test.LookupTxnPositionInExpected], blockWithTxs.Transactions[test.LookupTxnPositionInOriginal]))
		}
	}
}

// TestBlockTransactionCount tests BlockTransactionCount
func TestBlockTransactionCount(t *testing.T) {
	testConfig := beforeEach(t)

	type testSetType struct {
		BlockID       BlockID
		ExpectedCount uint64
	}
	testSet := map[string][]testSetType{
		"mock": {
			{
				BlockID:       WithBlockNumber(300000),
				ExpectedCount: 10,
			},
		},
		"testnet": {
			{
				BlockID:       WithBlockNumber(300000),
				ExpectedCount: 23,
			},
		},
		"mainnet": {},
	}[testEnv]
	for _, test := range testSet {
		spy := NewSpy(testConfig.client.c)
		testConfig.client.c = spy
		count, err := testConfig.client.BlockTransactionCount(context.Background(), test.BlockID)
		if err != nil {
			t.Fatal(err)
		}
		diff, err := spy.Compare(count, false)
		if err != nil {
			t.Fatal("expecting to match", err)
		}
		if diff != "FullMatch" {
			spy.Compare(count, true)
			t.Fatal("structure expecting to be FullMatch, instead", diff)
		}
		if count != test.ExpectedCount {
			t.Fatalf("structure expecting %d, instead: %d", test.ExpectedCount, count)
		}
	}
}

func TestCaptureUnsupportedBlockTxn(t *testing.T) {
	testConfig := beforeEach(t)

	type testSetType struct {
		StartBlock uint64
		EndBlock   uint64
	}
	testSet := map[string][]testSetType{
		"mock": {},
		"testnet": {
			{
				StartBlock: 310000,
				EndBlock:   310100,
			},
		},
		"mainnet": {},
	}[testEnv]
	for _, test := range testSet {
		for i := test.StartBlock; i < test.EndBlock; i++ {
			blockWithTxsInterface, err := testConfig.client.BlockWithTxs(context.Background(), WithBlockNumber(BlockNumber(i)))
			if err != nil {
				t.Fatal("BlockWithTxHashes match the expected error:", err)
			}
			blockWithTxs, ok := blockWithTxsInterface.(*BlockWithTxs)
			if !ok {
				t.Fatalf("expecting BlockWithTxs, instead %T", blockWithTxsInterface)
			}
			for k, v := range blockWithTxs.Transactions {
				if fmt.Sprintf("%T", v) != "rpc.InvokeTxnV0" &&
					fmt.Sprintf("%T", v) != "rpc.DeployTxn" &&
					fmt.Sprintf("%T", v) != "rpc.DeclareTxn" {
					t.Fatalf("New Type Detected %T at Block(%d)/Txn(%d)", v, i, k)
				}
			}

		}
	}
}

// TODO: Find a block with such a Txn
// TestBlockWithTxsAndInvokeTXNV1 tests BlockWithTxs with Invoke V1
func TestBlockWithTxsAndInvokeTXNV1(t *testing.T) {
	_ = beforeEach(t)

	type testSetType struct {
	}
	testSet := map[string][]testSetType{
		"mock": {},
		"testnet": {
			{},
		},
		"mainnet": {},
	}[testEnv]
	for range testSet {
		t.Fatalf("error running test: %v", errNotImplemented)
	}
}

// TestStateUpdate tests StateUpdateByHash
// TODO: this is not implemented yet with pathfinder as you can see from the
// [code](https://github.com/eqlabs/pathfinder/blob/927183552dad6dcdfebac16c8c1d2baf019127b1/crates/pathfinder/rpc_examples.sh#L37)
// check when it is and test when it is the case.
func TestStateUpdate(t *testing.T) {
	testConfig := beforeEach(t)

	type testSetType struct {
		BlockID                   BlockID
		ExpectedStateUpdateOutput StateUpdateOutput
	}
	testSet := map[string][]testSetType{
		"mock": {
			{
				BlockID: WithBlockNumber(300000),
				ExpectedStateUpdateOutput: StateUpdateOutput{
					BlockHash:    "0x4f1cee281edb6cb31b9ba5a8530694b5527cf05c5ac6502decf3acb1d0cec4",
					NewRoot:      "0x70677cda9269d47da3ff63bc87cf1c87d0ce167b05da295dc7fc68242b250b",
					OldRoot:      "0x19aa982a75263d4c4de4cc4c5d75c3dec32e00b95bef7bbb4d17762a0b138af",
					AcceptedTime: 0,
					StateDiff: StateDiff{
						StorageDiffs: []ContractStorageDiffItem{{
							Address: "0xe5cc6f2b6d34979184b88334eb64173fe4300cab46ecd3229633fcc45c83d4",
							Entries: []StorageEntry{{
								Key:   "0x1813aac5f5e7799684c6dc33e51f44d3627fd748c800724a184ed5be09b713e",
								Value: "0x630b4197",
							}},
						}},
					},
				},
			},
		},
		"testnet": {
			{
				BlockID: WithBlockTag("latest"),
			},
		},
		"mainnet": {},
	}[testEnv]
	for _, test := range testSet {
		spy := NewSpy(testConfig.client.c)
		testConfig.client.c = spy
		stateUpdate, err := testConfig.client.StateUpdate(context.Background(), test.BlockID)
		if err != nil {
			t.Fatal(err)
		}
		diff, err := spy.Compare(stateUpdate, false)
		if err != nil {
			t.Fatal("expecting to match", err)
		}
		if diff != "FullMatch" {
			spy.Compare(stateUpdate, true)
			t.Fatal("structure expecting to be FullMatch, instead", diff)
		}
		if uint64(stateUpdate.AcceptedTime) != 0 {
			t.Fatalf("structure expecting %d, instead: %d", 0, stateUpdate.AcceptedTime)
		}
	}
}
