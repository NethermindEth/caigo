package rpc

import (
	"context"
	"flag"
	"fmt"
	"math/big"
	"os"
	"strings"
	"testing"

	ethrpc "github.com/ethereum/go-ethereum/rpc"
	"github.com/joho/godotenv"
	"github.com/test-go/testify/require"
)

const (
	TestPublicKey            = "0x783318b2cc1067e5c06d374d2bb9a0382c39aabd009b165d7a268b882971d6"
	DevNetETHAddress         = "0x62230ea046a9a5fbc261ac77d03c8d41e5d442db2284587570ab46455fd2488"
	TestNetETHAddress        = "0x049d36570d4e46f48e99674bd3fcc84644ddd6b96f7c741b1562b82f9e004dc7"
	DevNetAccount032Address  = "0x06bb9425718d801fd06f144abb82eced725f0e81db61d2f9f4c9a26ece46a829"
	TestNetAccount032Address = "0x6ca4fdd437dffde5253ba7021ef7265c88b07789aa642eafda37791626edf00"
	DevNetAccount040Address  = "0x080dff79c6216ad300b872b73ff41e271c63f213f8a9dc2017b164befa53b9"
	TestNetAccount040Address = "0x6cbfa37f409610fee26eeb427ed854b3a4b24580d9b9ef6c3e38db7b3f7322c"
)

// testConfiguration is a type that is used to configure tests
type testConfiguration struct {
	provider *Provider
	base     string
}

var (
	// set the environment for the test, default: mock
	testEnv = "mock"

	// testConfigurations are predefined test configurations
	testConfigurations = map[string]testConfiguration{
		// Requires a Mainnet Starknet JSON-RPC compliant node (e.g. pathfinder)
		// (ref: https://github.com/eqlabs/pathfinder)
		"mainnet": {
			base: "http://localhost:9545/v0.2/rpc",
		},
		// Requires a Testnet Starknet JSON-RPC compliant node (e.g. pathfinder)
		// (ref: https://github.com/eqlabs/pathfinder)
		"testnet": {
			base: "http://localhost:9545/v0.2/rpc",
		},
		// Requires a Devnet configuration running locally
		// (ref: https://github.com/Shard-Labs/starknet-devnet)
		"devnet": {
			base: "http://localhost:5050/rpc",
		},
		// Used with a mock as a standard configuration, see `mock_test.go``
		"mock": {},
	}
)

// TestMain is a Go function that serves as the entry point for running tests.
//
// It takes a pointer to the testing.M struct as its parameter and returns nothing.
// The purpose of this function is to set up any necessary test environment
// variables before running the tests and to clean up any resources afterwards.
// It also parses command line flags and exits with the exit code returned by
// the testing.M.Run() function.
func TestMain(m *testing.M) {
	flag.StringVar(&testEnv, "env", "mock", "set the test environment")
	flag.Parse()

	os.Exit(m.Run())
}

// beforeEach initializes the test environment configuration before running the script.
//
// t: The testing.T object for testing purposes.
// Returns a pointer to the testConfiguration struct.
func beforeEach(t *testing.T) *testConfiguration {
	t.Helper()
	godotenv.Load(fmt.Sprintf(".env.%s", testEnv), ".env")
	testConfig, ok := testConfigurations[testEnv]
	if !ok {
		t.Fatal("env supports mock, testnet, mainnet or devnet")
	}
	if testEnv == "mock" {
		testConfig.provider = &Provider{
			c: &rpcMock{},
		}
		return &testConfig
	}

	testConfig.base = "https://starknet-goerli.cartridge.gg"
	base := os.Getenv("INTEGRATION_BASE")
	if base != "" {
		testConfig.base = base
	}
	c, err := ethrpc.DialContext(context.Background(), testConfig.base)
	if err != nil {
		t.Fatal("connect should succeed, instead:", err)
	}
	client := NewProvider(c)
	testConfig.provider = client
	t.Cleanup(func() {
		testConfig.provider.c.Close()
	})
	return &testConfig
}

// TestChainID is a function that tests the ChainID function in the Go test file.
//
// The function initializes a test configuration and defines a test set with different chain IDs for different environments.
// It then iterates over the test set and for each test, creates a new spy and sets the spy as the provider's client.
// The function calls the ChainID function and compares the returned chain ID with the expected chain ID.
// If there is a mismatch or an error occurs, the function logs a fatal error.
func TestChainID(t *testing.T) {
	testConfig := beforeEach(t)

	type testSetType struct {
		ChainID string
	}
	testSet := map[string][]testSetType{
		"devnet":  {{ChainID: "SN_GOERLI"}},
		"mainnet": {{ChainID: "SN_MAIN"}},
		"mock":    {{ChainID: "SN_GOERLI"}},
		"testnet": {{ChainID: "SN_GOERLI"}},
	}[testEnv]

	for _, test := range testSet {
		spy := NewSpy(testConfig.provider.c)
		testConfig.provider.c = spy
		chain, err := testConfig.provider.ChainID(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if _, err := spy.Compare(chain, false); err != nil {
			t.Fatal("expecting to match", err)
		}
		if chain != test.ChainID {
			t.Fatalf("expecting %s, instead: %s", test.ChainID, chain)
		}
	}
}

// TestSyncing is a test function that tests the syncing functionality.
//
// It initializes the test configuration and defines a testSet of chain IDs.
// It then iterates over the testSet and performs the following steps:
//   - Creates a new spy object.
//   - Sets the provider's context to the spy object.
//   - Calls the Syncing method of the provider.
//   - Checks if there is an error during syncing. If there is, it logs a fatal error.
//   - Checks if the starting block hash is not nil. If it is not nil, it performs the following checks:
//     - Compares the sync object with the spy object, expecting a full match.
//     - Parses the current block number as a big integer and checks if it is a positive number.
//     - Checks if the current block hash is a string starting with "0x".
//   - If the starting block hash is nil, it compares the sync object with the spy object and expects the current block hash to be nil.
//
// The function is a test case for the TestSyncing function.
func TestSyncing(t *testing.T) {
	// Function body
}

// TestSyncing tests the syncing functionality.
//
// It initializes a test configuration and sets up a test set. Then it loops
// through the test set and creates a spy object. It calls the Syncing function
// of the provider using the test configuration. It checks if there is any
// error during syncing, and if so, it fails the test. If the starting block
// hash is not nil, it compares the sync object with the spy object. It checks
// if the current block number is a positive number and if the current block
// hash starts with "0x". If the starting block hash is nil, it compares the
// sync object with the spy object and checks if the current block hash is nil.
func TestSyncing(t *testing.T) {
	testConfig := beforeEach(t)

	type testSetType struct {
		ChainID string
	}

	testSet := map[string][]testSetType{
		"devnet":  {},
		"mainnet": {{ChainID: "SN_MAIN"}},
		"mock":    {{ChainID: "MOCK"}},
		"testnet": {{ChainID: "SN_GOERLI"}},
	}[testEnv]

	for range testSet {
		spy := NewSpy(testConfig.provider.c)
		testConfig.provider.c = spy
		sync, err := testConfig.provider.Syncing(context.Background())
		if err != nil {
			t.Fatal("Syncing error:", err)
		}
		if sync.StartingBlockHash != nil {
			if diff, err := spy.Compare(sync, false); err != nil || diff != "FullMatch" {
				spy.Compare(sync, true)
				t.Fatal("expecting to match", err)
			}
			i, ok := big.NewInt(0).SetString(string(sync.CurrentBlockNum), 0)
			if !ok || i.Cmp(big.NewInt(0)) <= 0 {
				t.Fatal("CurrentBlockNum should be positive number, instead: ", sync.CurrentBlockNum)
			}
			if !strings.HasPrefix(sync.CurrentBlockHash.String(), "0x") {
				t.Fatal("current block hash should return a string starting with 0x")
			}
		} else {
			spy.Compare(sync, false)
			require.Nil(t, sync.CurrentBlockHash)

		}

	}
}
