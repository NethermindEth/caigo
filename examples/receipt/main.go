package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/rpc"

	setup "github.com/NethermindEth/starknet.go/examples/internal"
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

	fmt.Println("Established connection with the client")
	qwe, err := new(felt.Felt).SetString("0x3902a3ff7c6d9a0345c449914946f0136ccb15a2298e1ae3f1fda98a0e5370e")
	rec, err := client.TransactionReceipt(context.Background(), qwe)
	asd, err := json.MarshalIndent(rec, " ", " ")
	fmt.Println(string(asd))
}
