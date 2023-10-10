package rpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/NethermindEth/juno/core/felt"
)

var (
	feltZero = new(felt.Felt).SetUint64(0)
	feltOne  = new(felt.Felt).SetUint64(1)
	feltTwo  = new(felt.Felt).SetUint64(2)
)

// adaptTransaction adapts a TXN to a Transaction and returns it, along with any error encountered.
//
// t - the TXN to be adapted to a Transaction.
// Returns a Transaction, or an error if the adaptation failed.
func adaptTransaction(t TXN) (Transaction, error) {
	txMarshalled, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	switch t.Type {
	case TransactionType_Invoke:
		var tx InvokeTxnV1
		json.Unmarshal(txMarshalled, &tx)
		return tx, nil
	case TransactionType_Declare:
		switch {
		case t.Version.Equal(feltZero):
			var tx DeclareTxnV0
			json.Unmarshal(txMarshalled, &tx)
			return tx, nil
		case t.Version.Equal(feltOne):
			var tx DeclareTxnV1
			json.Unmarshal(txMarshalled, &tx)
			return tx, nil
		case t.Version.Equal(feltTwo):
			var tx DeclareTxnV2
			json.Unmarshal(txMarshalled, &tx)
			return tx, nil
		}
	case TransactionType_Deploy:
		var tx DeployTxn
		json.Unmarshal(txMarshalled, &tx)
		return tx, nil
	case TransactionType_DeployAccount:
		var tx DeployAccountTxn
		json.Unmarshal(txMarshalled, &tx)
		return tx, nil
	case TransactionType_L1Handler:
		var tx L1HandlerTxn
		json.Unmarshal(txMarshalled, &tx)
		return tx, nil
	}
	return nil, errors.New(fmt.Sprint("internal error with adaptTransaction() : unknown transaction type ", t.Type))

}

// TransactionByHash retrieves the details and status of a transaction by its hash.
//
// It takes a context.Context and a *felt.Felt as input parameters.
// The function returns a Transaction and an error.
func (provider *Provider) TransactionByHash(ctx context.Context, hash *felt.Felt) (Transaction, error) {
	// todo: update to return a custom Transaction type, then use adapt function
	var tx TXN
	if err := do(ctx, provider.c, "starknet_getTransactionByHash", &tx, hash); err != nil {
		if errors.Is(err, ErrHashNotFound) {
			return nil, ErrHashNotFound
		}
		return nil, err
	}
	return adaptTransaction(tx)
}

// TransactionByBlockIdAndIndex retrieves a transaction by its block ID and index.
//
// ctx: The context.Context object for the request.
// blockID: The ID of the block containing the transaction.
// index: The index of the transaction within the block.
// Returns: The retrieved Transaction object and an error, if any.
func (provider *Provider) TransactionByBlockIdAndIndex(ctx context.Context, blockID BlockID, index uint64) (Transaction, error) {
	var tx TXN
	if err := do(ctx, provider.c, "starknet_getTransactionByBlockIdAndIndex", &tx, blockID, index); err != nil {
		switch {
		case errors.Is(err, ErrInvalidTxnIndex):
			return nil, ErrInvalidTxnIndex
		case errors.Is(err, ErrBlockNotFound):
			return nil, ErrBlockNotFound
		}
		return nil, err
	}
	return adaptTransaction(tx)
}

// PendingTransaction returns a list of pending transactions in the transaction pool, recognized by this sequencer..
//
// ctx - The context.Context object for controlling the lifespan of the request.
// It carries deadlines, cancellation signals, and other request-scoped values.
//
// []Transaction - A slice of Transaction objects representing the pending transactions.
// error - An error that occurred during the execution of the function, if any.
func (provider *Provider) PendingTransaction(ctx context.Context) ([]Transaction, error) {
	txs := []Transaction{}
	if err := do(ctx, provider.c, "starknet_pendingTransactions", &txs, []interface{}{}); err != nil {
		return nil, err
	}
	return txs, nil
}

// TransactionReceipt fetches the transaction receipt for a given transaction hash.
//
// It takes a context and transaction hash as parameters and returns a TransactionReceipt and an error.
func (provider *Provider) TransactionReceipt(ctx context.Context, transactionHash *felt.Felt) (TransactionReceipt, error) {
	var receipt UnknownTransactionReceipt
	err := do(ctx, provider.c, "starknet_getTransactionReceipt", &receipt, transactionHash)
	if err != nil {
		if errors.Is(err, ErrHashNotFound) {
			return nil, ErrHashNotFound
		}
		return nil, err
	}
	return receipt.TransactionReceipt, nil
}
