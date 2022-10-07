package types

import (
	"encoding/json"
	"fmt"
	"strconv"

	ctypes "github.com/dontpanicdao/caigo/types"
)

type TransactionHash struct {
	TransactionHash ctypes.Hash `json:"transaction_hash"`
}

func (tx TransactionHash) Hash() ctypes.Hash {
	return tx.TransactionHash
}

func (tx *TransactionHash) UnmarshalJSON(input []byte) error {
	unquoted, err := strconv.Unquote(string(input))
	if err != nil {
		return err
	}
	tx.TransactionHash = ctypes.HexToHash(unquoted)
	return nil
}

func (tx TransactionHash) MarshalJSON() ([]byte, error) {
	b, err := tx.TransactionHash.MarshalText()
	if err != nil {
		return nil, err
	}

	return []byte(strconv.Quote(string(b))), nil
}

type Call struct {
	MaxFee string `json:"max_fee"`
	// Version of the transaction scheme, should be set to 0 or 1
	Version ctypes.NumAsHex `json:"version"`
	// Signature
	Signature []string `json:"signature"`
	// Nonce should only be set with Transaction V1
	Nonce *string `json:"nonce,omitempty"`

	ContractAddress ctypes.Hash `json:"contract_address"`

	// EntryPointSelector should only be set with Transaction V0
	EntryPointSelector *string `json:"entry_point_selector,omitempty"`

	// Calldata The parameters passed to the function
	Calldata []string `json:"calldata"`
}

type CommonTransaction struct {
	TransactionHash ctypes.Hash     `json:"transaction_hash,omitempty"`
	Type            TransactionType `json:"type,omitempty"`
	// MaxFee maximal fee that can be charged for including the transaction
	MaxFee string `json:"max_fee,omitempty"`
	// Version of the transaction scheme
	Version ctypes.NumAsHex `json:"version"`
	// Signature
	Signature []string `json:"signature,omitempty"`
	// Nonce
	Nonce string `json:"nonce,omitempty"`
}

// InvokeTxnDuck is a type used to understand the Invoke Version
type InvokeTxnDuck struct {
	AccountAddress     ctypes.Hash `json:"account_address"`
	ContractAddress    ctypes.Hash `json:"contract_address"`
	EntryPointSelector string      `json:"entry_point_selector"`
}

type InvokeTxnV0 struct {
	CommonTransaction
	ContractAddress    ctypes.Hash `json:"contract_address"`
	EntryPointSelector string      `json:"entry_point_selector"`

	// Calldata The parameters passed to the function
	Calldata []string `json:"calldata"`
}

func (tx InvokeTxnV0) Hash() ctypes.Hash {
	return tx.TransactionHash
}

type InvokeTxnV1 struct {
	CommonTransaction
	InvokeV1
}

func (tx InvokeTxnV1) Hash() ctypes.Hash {
	return tx.TransactionHash
}

type InvokeTxn interface{}

type L1HandlerTxn struct {
	TransactionHash ctypes.Hash     `json:"transaction_hash,omitempty"`
	Type            TransactionType `json:"type,omitempty"`
	// Version of the transaction scheme
	Version ctypes.NumAsHex `json:"version"`
	// Nonce
	Nonce string `json:"nonce,omitempty"`
	// MaxFee maximal fee that can be charged for including the transaction
	MaxFee string `json:"max_fee,omitempty"`
	// Signature
	Signature          []string    `json:"signature,omitempty"`
	ContractAddress    ctypes.Hash `json:"contract_address"`
	EntryPointSelector string      `json:"entry_point_selector"`

	// Calldata The parameters passed to the function
	Calldata []string `json:"calldata"`
}

func (tx L1HandlerTxn) Hash() ctypes.Hash {
	return tx.TransactionHash
}

type DeclareTxn struct {
	CommonTransaction

	// ClassHash the hash of the declared class
	ClassHash string `json:"class_hash"`

	// SenderAddress the address of the account contract sending the declaration transaction
	SenderAddress string `json:"sender_address"`
}

func (tx DeclareTxn) Hash() ctypes.Hash {
	return tx.TransactionHash
}

// DeployTxn The structure of a deploy transaction. Note that this transaction type is deprecated and will no longer be supported in future versions
type DeployTxn struct {
	CommonTransaction
	// ClassHash The hash of the deployed contract's class
	ClassHash string `json:"class_hash"`

	// ContractAddress The address of the deployed contract
	ContractAddress string `json:"contract_address"`

	// ContractAddressSalt The salt for the address of the deployed contract
	ContractAddressSalt string `json:"contract_address_salt"`

	// ConstructorCalldata The parameters passed to the constructor
	ConstructorCalldata []string `json:"constructor_calldata"`
}

func (tx DeployTxn) Hash() ctypes.Hash {
	return tx.TransactionHash
}

type DeployTxnProperties struct {
	// Version of the transaction scheme
	Version ctypes.NumAsHex `json:"version"`
	Type    TransactionType `json:"type"`
	// ContractAddressSalt The salt for the address of the deployed contract
	ContractAddressSalt string `json:"contract_address_salt"`
	// ConstructorCalldata The parameters passed to the constructor
	ConstructorCalldata []string `json:"constructor_calldata"`
}

type Transaction interface {
	Hash() ctypes.Hash
}

// InvokeV0 version 0 invoke transaction
type InvokeV0 ctypes.FunctionCall

// InvokeV1 version 1 invoke transaction
type InvokeV1 struct {
	SenderAddress ctypes.Hash `json:"sender_address"`
	// Calldata The parameters passed to the function
	Calldata []string `json:"calldata"`
}

type Transactions []Transaction

func (txns *Transactions) UnmarshalJSON(data []byte) error {
	var dec []interface{}
	if err := json.Unmarshal(data, &dec); err != nil {
		return err
	}

	unmarshalled := make([]Transaction, len(dec))
	for i, t := range dec {
		var err error
		unmarshalled[i], err = unmarshalTxn(t)
		if err != nil {
			return err
		}
	}

	*txns = unmarshalled
	return nil
}

type UnknownTransaction struct{ Transaction }

func (txn *UnknownTransaction) UnmarshalJSON(data []byte) error {
	var dec map[string]interface{}
	if err := json.Unmarshal(data, &dec); err != nil {
		return err
	}

	t, err := unmarshalTxn(dec)
	if err != nil {
		return err
	}

	*txn = UnknownTransaction{t}
	return nil
}

func unmarshalTxn(t interface{}) (Transaction, error) {
	switch casted := t.(type) {
	case string:
		return TransactionHash{ctypes.HexToHash(casted)}, nil
	case map[string]interface{}:
		switch TransactionType(casted["type"].(string)) {
		case TransactionType_Declare:
			var txn DeclareTxn
			remarshal(casted, &txn)
			return txn, nil
		case TransactionType_Deploy:
			var txn DeployTxn
			remarshal(casted, &txn)
			return txn, nil
		case TransactionType_Invoke:
			if casted["version"].(string) == "0x0" {
				var txn InvokeTxnV0
				remarshal(casted, &txn)
				return txn, nil
			} else {
				var txn InvokeTxnV1
				remarshal(casted, &txn)
				return txn, nil
			}
		case TransactionType_L1Handler:
			var txn L1HandlerTxn
			remarshal(casted, &txn)
			return txn, nil
		}
	}

	return nil, fmt.Errorf("unknown transaction type: %v", t)
}

func remarshal(v interface{}, dst interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, dst); err != nil {
		return err
	}

	return nil
}
