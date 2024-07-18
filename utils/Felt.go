package utils

import (
	"fmt"
	"math/big"

	"github.com/NethermindEth/juno/core/felt"
)

// Uint64ToFelt generates a new *felt.Felt from a given uint64 number.
//
// Parameters:
// - num: the uint64 number to convert to a *felt.Felt
// Returns:
// - *felt.Felt: a *felt.Felt
func Uint64ToFelt(num uint64) *felt.Felt {
	return new(felt.Felt).SetUint64(num)
}

// HexToFelt converts a hexadecimal string to a *felt.Felt object.
//
// Parameters:
// - hex: the input hexadecimal string to be converted.
// Returns:
// - *felt.Felt: a *felt.Felt object
// - error: if conversion fails
func HexToFelt(hex string) (*felt.Felt, error) {
	return new(felt.Felt).SetString(hex)
}

// HexArrToFelt converts an array of hexadecimal strings to an array of felt objects.
//
// The function iterates over each element in the hexArr array and calls the HexToFelt function to convert each hexadecimal value to a felt object.
// If any error occurs during the conversion, the function will return nil and the corresponding error.
// Otherwise, it appends the converted felt object to the feltArr array.
// Finally, the function returns the feltArr array containing all the converted felt objects.
//
// Parameters:
// - hexArr: an array of strings representing hexadecimal values
// Returns:
// - []*felt.Felt: an array of *felt.Felt objects, or nil if there was
// - error: an error if any
func HexArrToFelt(hexArr []string) ([]*felt.Felt, error) {

	feltArr := make([]*felt.Felt, len(hexArr))
	for i, e := range hexArr {
		felt, err := HexToFelt(e)
		if err != nil {
			return nil, err
		}
		feltArr[i] = felt
	}
	return feltArr, nil

}

// FeltToBigInt converts a Felt value to a *big.Int.
//
// Parameters:
// - f: the Felt value to convert
// Returns:
// - *big.Int: the converted value
func FeltToBigInt(f *felt.Felt) *big.Int {
	tmp := f.Bytes()
	return new(big.Int).SetBytes(tmp[:])
}

// BigIntToFelt converts a big integer to a felt.Felt.
//
// Parameters:
// - big: the big integer to convert
// Returns:
// - *felt.Felt: the converted value
func BigIntToFelt(big *big.Int) *felt.Felt {
	return new(felt.Felt).SetBytes(big.Bytes())
}

// FeltArrToBigIntArr converts an array of Felt objects to an array of big.Int objects.
//
// Parameters:
// - f: the array of Felt objects to convert
// Returns:
// - []*big.Int: the array of big.Int objects
func FeltArrToBigIntArr(f []*felt.Felt) []*big.Int {
	var bigArr []*big.Int
	for _, felt := range f {
		bigArr = append(bigArr, FeltToBigInt(felt))
	}
	return bigArr
}

// U256ToFelt converts big int an array of Felt objects to represent U256 datat type
//
// Parameters:
// - num: big.Int
// Returns:
// - []*felt.Felt: the array of felt.Felt objects
// - error: an error, if any
func U256ToFelt(num *big.Int) ([]*felt.Felt, error) {
	bytes := num.Bytes()
	if len(bytes) > 32 {
		return nil, fmt.Errorf("not a valid U256")
	}

	all := make([]byte, 32)
	copy(all[32-len(bytes):], bytes[:])

	least := new(felt.Felt).SetBytes(all[16:])
	significant := new(felt.Felt).SetBytes(all[0:16])
	return []*felt.Felt{least, significant}, nil
}

// FeltArrToU256 array of Felt objects that represents U256 data type to big.Int
//
// Parameters:
// - arr: []*felt.Felt
// Returns:
// - *big.Int: big.Int representation of U256
// - error: an error, if any
func FeltArrToU256(arr []*felt.Felt) (*big.Int, error) {
	if len(arr) != 2 {
		return nil, fmt.Errorf("not a valid felt array for U256 conversion")
	}

	significant := arr[1].Bytes()
	least := arr[0].Bytes()
	res := make([]byte, 32)
	copy(res[0:16], significant[16:])
	copy(res[16:], least[16:])
	return new(big.Int).SetBytes(res), nil
}
