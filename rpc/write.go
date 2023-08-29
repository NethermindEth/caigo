package rpc

import (
	"context"
	"errors"
	"strings"
)

type BroadcastedInvokeTransaction interface{}

// AddInvokeTransaction estimates the fee for a given StarkNet transaction.
func (provider *Provider) AddInvokeTransaction(ctx context.Context, broadcastedInvoke BroadcastedInvokeTransaction) (*AddInvokeTransactionResponse, error) {
	// TODO: EntryPointSelector now part of calldata
	// tx, ok := broadcastedInvoke.(BroadcastedInvokeV0Transaction)
	// if ok {
	// 	tx.EntryPointSelector = fmt.Sprintf("0x%x", types.GetSelectorFromName(tx.EntryPointSelector))
	// 	broadcastedInvoke = tx
	// }
	var output AddInvokeTransactionResponse
	switch invoke := broadcastedInvoke.(type) {
	case BroadcastedInvokeV1Transaction:
		if err := do(ctx, provider.c, "starknet_addInvokeTransaction", &output, invoke); err != nil {
			return nil, err
		}
		return &output, nil
	}
	return nil, errors.New("invalid invoke type")
}

// AddDeclareTransaction submits a new class declaration transaction.
func (provider *Provider) AddDeclareTransaction(ctx context.Context, declareTransaction BroadcastedDeclareTransaction) (*AddDeclareTransactionResponse, error) {
	var result AddDeclareTransactionResponse
	if err := do(ctx, provider.c, "starknet_addDeclareTransaction", &result, declareTransaction); err != nil {
		if strings.Contains(err.Error(), "Invalid contract class") {
			return nil, ErrInvalidContractClass
		}
		return nil, err
	}
	return &result, nil
}

// AddDeployAccountTransaction manages the DEPLOY_ACCOUNT syscall
func (provider *Provider) AddDeployAccountTransaction(ctx context.Context, deployAccountTransaction BroadcastedDeployAccountTransaction) (*AddDeployTransactionResponse, error) {
	var result AddDeployTransactionResponse
	if err := do(ctx, provider.c, "starknet_addDeployAccountTransaction", &result, deployAccountTransaction); err != nil {
		if strings.Contains(err.Error(), "Class hash not found") {
			return nil, ErrClassHashNotFound
		}
		return nil, err
	}
	return &result, nil
}
