package rpc

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/NethermindEth/juno/core/felt"
)

// CommonTransactionReceipt Common properties for a transaction receipt
type CommonTransactionReceipt struct {
	// TransactionHash The hash identifying the transaction
	TransactionHash *felt.Felt `json:"transaction_hash"`
	// ActualFee The fee that was charged by the sequencer
	ActualFee       *felt.Felt         `json:"actual_fee"`
	ExecutionStatus TxnExecutionStatus `json:"execution_status"`
	FinalityStatus  TxnFinalityStatus  `json:"finality_status"`
	BlockHash       *felt.Felt         `json:"block_hash"`
	BlockNumber     uint64             `json:"block_number"`
	Type            TransactionType    `json:"type,omitempty"`
	MessagesSent    []MsgToL1          `json:"messages_sent"`
	RevertReason    string             `json:"revert_reason,omitempty"`
	// Events The events emitted as part of this transaction
	Events []Event `json:"events"`
}

func (tr CommonTransactionReceipt) Hash() *felt.Felt {
	return tr.TransactionHash
}

func (tr CommonTransactionReceipt) GetExecutionStatus() TxnExecutionStatus {
	return tr.ExecutionStatus
}

// TODO: check how we can move that type up in starknet.go/types
type TransactionType string

const (
	TransactionType_Declare       TransactionType = "DECLARE"
	TransactionType_DeployAccount TransactionType = "DEPLOY_ACCOUNT"
	TransactionType_Deploy        TransactionType = "DEPLOY"
	TransactionType_Invoke        TransactionType = "INVOKE"
	TransactionType_L1Handler     TransactionType = "L1_HANDLER"
)

// UnmarshalJSON unmarshals the JSON data into a TransactionType.
//
// The data parameter is the JSON data to be unmarshaled.
// It returns an error if the unmarshaling fails.
func (tt *TransactionType) UnmarshalJSON(data []byte) error {
	unquoted, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}

	switch unquoted {
	case "DECLARE":
		*tt = TransactionType_Declare
	case "DEPLOY_ACCOUNT":
		*tt = TransactionType_DeployAccount
	case "DEPLOY":
		*tt = TransactionType_Deploy
	case "INVOKE":
		*tt = TransactionType_Invoke
	case "L1_HANDLER":
		*tt = TransactionType_L1Handler
	default:
		return fmt.Errorf("unsupported type: %s", data)
	}

	return nil
}

// MarshalJSON marshals the TransactionType to JSON.
//
// It returns a byte slice and an error.
func (tt TransactionType) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(string(tt))), nil
}

// InvokeTransactionReceipt Invoke Transaction Receipt
type InvokeTransactionReceipt CommonTransactionReceipt

// Hash returns the transaction hash of the InvokeTransactionReceipt.
//
// It returns a pointer to a felt.Felt.
func (tr InvokeTransactionReceipt) Hash() *felt.Felt {
	return tr.TransactionHash
}

// GetExecutionStatus returns the execution status of the InvokeTransactionReceipt.
//
// No parameters.
// Returns a TxnExecutionStatus.
func (tr InvokeTransactionReceipt) GetExecutionStatus() TxnExecutionStatus {
	return tr.ExecutionStatus
}

// DeclareTransactionReceipt Declare Transaction Receipt
type DeclareTransactionReceipt CommonTransactionReceipt

// Hash returns the transaction hash of the DeclareTransactionReceipt.
//
// No parameters.
// Returns a pointer to a `felt.Felt` object.
func (tr DeclareTransactionReceipt) Hash() *felt.Felt {
	return tr.TransactionHash
}

// GetExecutionStatus returns the execution status of the transaction receipt.
//
// No parameters.
// Returns a TxnExecutionStatus.
func (tr DeclareTransactionReceipt) GetExecutionStatus() TxnExecutionStatus {
	return tr.ExecutionStatus
}

// DeployTransactionReceipt Deploy  Transaction Receipt
type DeployTransactionReceipt struct {
	CommonTransactionReceipt
	// The address of the deployed contract
	ContractAddress *felt.Felt `json:"contract_address"`
}

func (tr DeployTransactionReceipt) Hash() *felt.Felt {
	return tr.TransactionHash
}

// GetExecutionStatus returns the execution status of the deploy transaction receipt.
//
// No parameters.
// Returns the transaction execution status.
func (tr DeployTransactionReceipt) GetExecutionStatus() TxnExecutionStatus {
	return tr.ExecutionStatus
}

// DeployAccountTransactionReceipt Deploy Account Transaction Receipt
type DeployAccountTransactionReceipt struct {
	CommonTransactionReceipt
	// ContractAddress The address of the deployed contract
	ContractAddress *felt.Felt `json:"contract_address"`
}

// Hash returns the transaction hash of the DeployAccountTransactionReceipt.
//
// It does not take any parameters.
// It returns a pointer to a `felt.Felt` object.
func (tr DeployAccountTransactionReceipt) Hash() *felt.Felt {
	return tr.TransactionHash
}

// GetExecutionStatus returns the execution status of the DeployAccountTransactionReceipt.
//
// It does not take any parameters.
// It returns a TxnExecutionStatus.
func (tr DeployAccountTransactionReceipt) GetExecutionStatus() TxnExecutionStatus {
	return tr.ExecutionStatus
}

// L1HandlerTransactionReceipt L1 Handler Transaction Receipt
type L1HandlerTransactionReceipt CommonTransactionReceipt

// Hash returns the transaction hash.
//
// No parameters.
// Returns a pointer to felt.Felt.
func (tr L1HandlerTransactionReceipt) Hash() *felt.Felt {
	return tr.TransactionHash
}

// GetExecutionStatus returns the execution status of the L1HandlerTransactionReceipt.
//
// This function does not take any parameters.
// It returns a TxnExecutionStatus.
func (tr L1HandlerTransactionReceipt) GetExecutionStatus() TxnExecutionStatus {
	return tr.ExecutionStatus
}

type PendingDeployTransactionReceipt struct {
	CommonTransactionReceipt
	// The address of the deployed contract
	ContractAddress *felt.Felt `json:"contract_address"`
}

// Hash returns the transaction hash associated with the PendingDeployTransactionReceipt.
//
// It does not take any parameters.
// It returns a pointer to a `felt.Felt` object.
func (tr PendingDeployTransactionReceipt) Hash() *felt.Felt {
	return tr.TransactionHash
}

// GetExecutionStatus returns the execution status of the pending deploy transaction receipt.
//
// No parameters.
// Returns the transaction execution status.
func (tr PendingDeployTransactionReceipt) GetExecutionStatus() TxnExecutionStatus {
	return tr.ExecutionStatus
}

type PendingCommonTransactionReceiptProperties struct {
	// TransactionHash The hash identifying the transaction
	TransactionHash *felt.Felt `json:"transaction_hash"`
	// ActualFee The fee that was charged by the sequencer
	ActualFee       *felt.Felt         `json:"actual_fee"`
	Type            TransactionType    `json:"type,omitempty"`
	MessagesSent    []MsgToL1          `json:"messages_sent"`
	ExecutionStatus TxnExecutionStatus `json:"execution_status"`
	FinalityStatus  TxnFinalityStatus  `json:"finality_status"`
	RevertReason    string             `json:"revert_reason"`
	// Events The events emitted as part of this transaction
	Events []Event `json:"events"`
}

// Hash returns the transaction hash of the PendingCommonTransactionReceiptProperties.
//
// It does not take any parameters.
// It returns a pointer to a felt.Felt object.
func (tr PendingCommonTransactionReceiptProperties) Hash() *felt.Felt {
	return tr.TransactionHash
}

// GetExecutionStatus returns the execution status of the PendingCommonTransactionReceiptProperties.
//
// No parameters are required.
// It returns a TxnExecutionStatus.
func (tr PendingCommonTransactionReceiptProperties) GetExecutionStatus() TxnExecutionStatus {
	return tr.ExecutionStatus
}

type TransactionReceipt interface {
	Hash() *felt.Felt
	GetExecutionStatus() TxnExecutionStatus
}

type MsgToL1 struct {
	// FromAddress The address of the L2 contract sending the message
	FromAddress *felt.Felt `json:"from_address"`
	// ToAddress The target L1 address the message is sent to
	ToAddress *felt.Felt `json:"to_address"`
	//Payload  The payload of the message
	Payload []*felt.Felt `json:"payload"`
}

type MsgFromL1 struct {
	// FromAddress The address of the L1 contract sending the message
	FromAddress string `json:"from_address"`
	// ToAddress The target L2 address the message is sent to
	ToAddress *felt.Felt `json:"to_address"`
	// EntryPointSelector The selector of the l1_handler in invoke in the target contract
	Selector *felt.Felt `json:"entry_point_selector"`
	//Payload  The payload of the message
	Payload []*felt.Felt `json:"payload"`
}

type UnknownTransactionReceipt struct{ TransactionReceipt }

// UnmarshalJSON unmarshals the JSON data into the UnknownTransactionReceipt struct.
//
// It takes a byte slice `data` as input and returns an error if the unmarshaling fails.
func (tr *UnknownTransactionReceipt) UnmarshalJSON(data []byte) error {
	var dec map[string]interface{}
	if err := json.Unmarshal(data, &dec); err != nil {
		return err
	}

	t, err := unmarshalTransactionReceipt(dec)
	if err != nil {
		return err
	}
	*tr = UnknownTransactionReceipt{t}
	return nil
}

// unmarshalTransactionReceipt unmarshals a transaction receipt from a generic interface{}.
//
// It takes in a generic interface{} 't' and returns a TransactionReceipt and an error.
func unmarshalTransactionReceipt(t interface{}) (TransactionReceipt, error) {
	switch casted := t.(type) {
	case map[string]interface{}:
		// NOTE(tvanas): Pathfinder 0.3.3 does not return
		// transaction receipt types. We handle this by
		// naively marshalling into an invoke type. Once it
		// is supported, this condition can be removed.
		typ, ok := casted["type"]
		if !ok {
			return nil, fmt.Errorf("unknown transaction type: %v", t)
		}

		// Pending doesn't have a block number
		if casted["block_hash"] == nil {
			switch TransactionType(typ.(string)) {
			case TransactionType_Deploy:
				var txn PendingDeployTransactionReceipt
				remarshal(casted, &txn)
				return txn, nil
			default:
				var txn PendingCommonTransactionReceiptProperties
				remarshal(casted, &txn)
				return txn, nil
			}
		}

		switch TransactionType(typ.(string)) {
		case TransactionType_Invoke:
			var txn InvokeTransactionReceipt
			remarshal(casted, &txn)
			return txn, nil
		case TransactionType_L1Handler:
			var txn L1HandlerTransactionReceipt
			remarshal(casted, &txn)
			return txn, nil
		case TransactionType_Declare:
			var txn DeclareTransactionReceipt
			remarshal(casted, &txn)
			return txn, nil
		case TransactionType_Deploy:
			var txn DeployTransactionReceipt
			remarshal(casted, &txn)
			return txn, nil
		case TransactionType_DeployAccount:
			var txn DeployAccountTransactionReceipt
			remarshal(casted, &txn)
			return txn, nil
		}
	}

	return nil, fmt.Errorf("unknown transaction type: %v", t)
}
