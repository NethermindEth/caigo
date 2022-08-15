package types

import (
	"context"

	"github.com/dontpanicdao/caigo/felt"
)

type Provider interface {
	AccountNonce(context.Context, felt.Felt) (*felt.Felt, error)
	BlockByHash(context.Context, felt.Felt, string) (*Block, error)
	BlockByNumber(context.Context, int64, string) (*Block, error)
	Call(context.Context, FunctionCall, string) ([]string, error)
	ChainID(context.Context) (string, error)
	Invoke(context.Context, FunctionInvoke) (*AddTxResponse, error)
	TransactionByHash(context.Context, string) (*Transaction, error)
	TransactionReceipt(context.Context, string) (*TransactionReceipt, error)
	EstimateFee(context.Context, FunctionInvoke, string) (*FeeEstimate, error)
	Class(context.Context, string) (*ContractClass, error)
	ClassHashAt(context.Context, string) (*felt.Felt, error)
	ClassAt(context.Context, string) (*ContractClass, error)
}
