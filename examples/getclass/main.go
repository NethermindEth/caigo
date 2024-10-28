package main

import (
	"context"
	"fmt"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/rpc"

	setup "github.com/NethermindEth/starknet.go/examples/internal"
)

var (
	someContract               string = "0x049D36570D4e46f48e99674bd3fcc84644DdD6b96F7C741B1562B82f9e004dC7" // Sepolia ETH contract address
	contractMethod             string = "decimals"
	contractMethodWithCalldata string = "balance_of"
)

// main entry point of the program.
//
// It initializes the environment and establishes a connection with the client.
// It then makes two contract calls and prints the responses.
//
// Parameters:
//
//	none
//
// Returns:
//
//	none
func main() {
	fmt.Println("Starting simpleCall example")

	// Load variables from '.env' file
	rpcProviderUrl := setup.GetRpcProviderUrl()

	// Initialize connection to RPC provider
	client, err := rpc.NewProvider(rpcProviderUrl)
	if err != nil {
		panic(fmt.Sprintf("Error dialing the RPC provider: %s", err))
	}
	classhash, _ := new(felt.Felt).SetString("0x060309f37b2b47b167c41810f0d95b9018ec36a0f1f65b4d5d6bf2f0b7f1fc89")
	class, err := client.Class(context.Background(), rpc.BlockID{Tag: "latest"}, classhash)
	fmt.Println(class)
}
