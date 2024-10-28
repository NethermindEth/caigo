package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/account"
	"github.com/NethermindEth/starknet.go/rpc"
	"github.com/NethermindEth/starknet.go/utils"

	setup "github.com/NethermindEth/starknet.go/examples/internal"
)

// starkli class-by-hash --network sepolia 0x060309f37b2b47b167c41810f0d95b9018ec36a0f1f65b4d5d6bf2f0b7f1fc89 > ./pitchlakeclass.json

func main() {
	// Load variables from '.env' file
	rpcProviderUrl := setup.GetRpcProviderUrl()
	accountAddress := setup.GetAccountAddress()
	accountCairoVersion := setup.GetAccountCairoVersion()
	privateKey := setup.GetPrivateKey()
	publicKey := setup.GetPublicKey()

	// Initialize connection to RPC provider
	client, err := rpc.NewProvider(rpcProviderUrl)
	if err != nil {
		panic(fmt.Sprintf("Error dialing the RPC provider: %s", err))
	}
	fmt.Println(client.SpecVersion(context.Background()))

	// Initialize the account memkeyStore (set public and private keys)
	ks := account.NewMemKeystore()
	privKeyBI, ok := new(big.Int).SetString(privateKey, 0)
	if !ok {
		panic("Fail to convert privKey to bitInt")
	}
	ks.Put(publicKey, privKeyBI)

	// Here we are converting the account address to felt
	accountAddressInFelt, err := utils.HexToFelt(accountAddress)
	if err != nil {
		fmt.Println("Failed to transform the account address, did you give the hex address?")
		panic(err)
	}
	// Initialize the account
	accnt, err := account.NewAccount(client, accountAddressInFelt, publicKey, ks, accountCairoVersion)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println("Established connection with the client")

	// Getting the nonce from the account
	nonce, err := accnt.Nonce(context.Background(), rpc.BlockID{Tag: "latest"}, accnt.AccountAddress)
	if err != nil {
		panic(err)
	}
	file, err := os.Open("pitchlakeclass.json")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Read the file contents
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var classs rpc.ContractClass

	err = json.Unmarshal(data, &classs)
	if err != nil {
		log.Fatalf("Failed toUnmarshalUnmarshale: %v", err)
	}
	compclasshash, _ := new(felt.Felt).SetString("0x05fb022cff754b819f96b3f7c7731eeb69735bb8f9abe21192d06fd1e34c1c0b") // doesn't matter??
	// Building the InvokeTx struct
	declareTxn := rpc.BroadcastDeclareTxnV2{
		Type:              rpc.TransactionType_Declare,
		SenderAddress:     accountAddressInFelt,
		CompiledClassHash: compclasshash,
		MaxFee:            new(felt.Felt).SetUint64(1234567),
		Version:           rpc.TransactionV2,
		Nonce:             nonce,
		ContractClass:     classs,
	}
	classhash, _ := new(felt.Felt).SetString("0x060309f37b2b47b167c41810f0d95b9018ec36a0f1f65b4d5d6bf2f0b7f1fc89")
	// Signing of the transaction that is done by the account
	tmpdeclare := rpc.DeclareTxnV2{
		Type:              declareTxn.Type,
		SenderAddress:     declareTxn.SenderAddress,
		CompiledClassHash: declareTxn.CompiledClassHash,
		MaxFee:            declareTxn.MaxFee,
		Version:           declareTxn.Version,
		Signature:         declareTxn.Signature,
		Nonce:             declareTxn.Nonce,
		ClassHash:         classhash,
	}
	err = accnt.SignDeclareTransaction(context.Background(), &tmpdeclare)
	if err != nil {
		panic(err)
	}
	declareTxn.Signature = tmpdeclare.Signature

	// After the signing we finally call the AddInvokeTransaction in order to invoke the contract function
	resp, err := accnt.AddDeclareTransaction(context.Background(), declareTxn)
	if err != nil {
		setup.PanicRPC(err)
	}

	fmt.Println("Waiting for the transaction status...")
	time.Sleep(time.Second * 3) // Waiting 3 seconds

	//Getting the transaction status
	txStatus, err := client.GetTransactionStatus(context.Background(), resp.TransactionHash)
	if err != nil {
		setup.PanicRPC(err)
	}

	// This returns us with the transaction hash and status
	fmt.Printf("Transaction hash response: %v\n", resp.TransactionHash)
	fmt.Printf("Transaction execution status: %s\n", txStatus.ExecutionStatus)
	fmt.Printf("Transaction status: %s\n", txStatus.FinalityStatus)

}
