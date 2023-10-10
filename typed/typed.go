package typed

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math/big"
	"regexp"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/curve"
	"github.com/NethermindEth/starknet.go/utils"
)

type TypedData struct {
	Types       map[string]TypeDef
	PrimaryType string
	Domain      Domain
	Message     TypedMessage
}

type Domain struct {
	Name    string
	Version string
	ChainId string
}

type TypeDef struct {
	Encoding    *big.Int
	Definitions []Definition
}

type Definition struct {
	Name string
	Type string
}

type TypedMessage interface {
	FmtDefinitionEncoding(string) []*big.Int
}

// FmtDefinitionEncoding formats the definition (standard Starknet Domain) encoding.
//
// It takes a field string as input and returns a slice of big integers.
func (dm Domain) FmtDefinitionEncoding(field string) (fmtEnc []*big.Int) {
	processStrToBig := func(fieldVal string) {
		feltVal := strToFelt(fieldVal)
		bigInt := utils.FeltToBigInt(feltVal)
		fmtEnc = append(fmtEnc, bigInt)
	}

	switch field {
	case "name":
		processStrToBig(dm.Name)
	case "version":
		processStrToBig(dm.Version)
	case "chainId":
		processStrToBig(dm.ChainId)
	}
	return fmtEnc
}

// strToFelt converts a string (decimal, hexadecimal or UTF8 charset) to a *felt.Felt.
//
// It takes a string as a parameter and returns a pointer to a *felt.Felt.
func strToFelt(str string) *felt.Felt {
	var f = &felt.Zero
	asciiRegexp := regexp.MustCompile(`^([[:graph:]]|[[:space:]]){1,31}$`)

	if b, ok := new(big.Int).SetString(str, 0); ok {
		f.SetBytes(b.Bytes())
		return f
	}
	// TODO: revisit conversation on seperate 'ShortString' conversion
	if asciiRegexp.MatchString(str) {
		hexStr := hex.EncodeToString([]byte(str))
		if b, ok := new(big.Int).SetString(hexStr, 16); ok {
			f.SetBytes(b.Bytes())
			return f
		}
	}

	return f
}

// NewTypedData initializes a new TypedData object with the given types, primary type, and domain
// for interacting and signing in accordance with https://github.com/0xs34n/starknet.js/tree/develop/src/utils/typedData
//
// It takes the following parameters:
// - types: a map[string]TypeDef representing the types associated with their names.
// - pType: a string representing the primary type.
// - dom: a Domain representing the domain.
//
// It returns a TypedData object and an error. The TypedData object contains the following fields:
// - Types: a map[string]TypeDef representing the types associated with their names.
// - PrimaryType: a string representing the primary type.
// - Domain: a Domain representing the domain.
//
// The function returns the TypedData object and an error. If the primary type is invalid, it returns
// an error with the message "invalid primary type: {pType}". If there is an error encoding the type hash,
// it returns an error with the message "error encoding type hash: {enc.String()} {err}".
func NewTypedData(types map[string]TypeDef, pType string, dom Domain) (td TypedData, err error) {
	td = TypedData{
		Types:       types,
		PrimaryType: pType,
		Domain:      dom,
	}
	if _, ok := td.Types[pType]; !ok {
		return td, fmt.Errorf("invalid primary type: %s", pType)
	}

	for k, v := range td.Types {
		enc, err := td.GetTypeHash(k)
		if err != nil {
			return td, fmt.Errorf("error encoding type hash: %s %w", enc.String(), err)
		}
		v.Encoding = enc
		td.Types[k] = v
	}
	return td, nil
}

// GetMessageHash calculates the hash of a typed message for a given account using the StarkCurve.
// (ref: https://github.com/0xs34n/starknet.js/blob/767021a203ac0b9cdb282eb6d63b33bfd7614858/src/utils/typedData/index.ts#L166)
//
// Parameters:
// - account: A pointer to a big.Int representing the account.
// - msg: A TypedMessage object representing the message.
// - sc: A StarkCurve object representing the curve.
//
// Returns:
// - hash: A pointer to a big.Int representing the calculated hash.
// - err: An error object indicating any error that occurred during the calculation.
func (td TypedData) GetMessageHash(account *big.Int, msg TypedMessage, sc curve.StarkCurve) (hash *big.Int, err error) {
	elements := []*big.Int{utils.UTF8StrToBig("Starknet Message")}

	domEnc, err := td.GetTypedMessageHash("StarknetDomain", td.Domain, sc)
	if err != nil {
		return hash, fmt.Errorf("could not hash domain: %w", err)
	}
	elements = append(elements, domEnc)
	elements = append(elements, account)

	msgEnc, err := td.GetTypedMessageHash(td.PrimaryType, msg, sc)
	if err != nil {
		return hash, fmt.Errorf("could not hash message: %w", err)
	}

	elements = append(elements, msgEnc)
	hash, err = sc.ComputeHashOnElements(elements)
	return hash, err
}

// GetTypedMessageHash calculates the hash of a typed message using the provided StarkCurve.
//
// Parameters:
//    - inType: the type of the message
//    - msg: the typed message
//    - sc: the StarkCurve used for hashing
//
// Returns:
//    - hash: the calculated hash
//    - err: any error that occurred during the calculation
func (td TypedData) GetTypedMessageHash(inType string, msg TypedMessage, sc curve.StarkCurve) (hash *big.Int, err error) {
	prim := td.Types[inType]
	elements := []*big.Int{prim.Encoding}

	for _, def := range prim.Definitions {
		if def.Type == "felt" {
			fmtDefinitions := msg.FmtDefinitionEncoding(def.Name)
			elements = append(elements, fmtDefinitions...)
			continue
		}

		innerElements := []*big.Int{}
		encType := td.Types[def.Type]
		innerElements = append(innerElements, encType.Encoding)
		fmtDefinitions := msg.FmtDefinitionEncoding(def.Name)
		innerElements = append(innerElements, fmtDefinitions...)
		innerElements = append(innerElements, big.NewInt(int64(len(innerElements))))

		innerHash, err := sc.HashElements(innerElements)
		if err != nil {
			return hash, fmt.Errorf("error hashing internal elements: %v %w", innerElements, err)
		}
		elements = append(elements, innerHash)
	}

	hash, err = sc.ComputeHashOnElements(elements)
	return hash, err
}

// GetTypeHash returns the hash of the given type.
//
// It takes inType as a parameter, which is the type to encode.
// It returns ret, which is the hash of the given type, and err, which is an error if there was an issue encoding the type.
func (td TypedData) GetTypeHash(inType string) (ret *big.Int, err error) {
	enc, err := td.EncodeType(inType)
	if err != nil {
		return ret, err
	}
	sel := utils.GetSelectorFromName(enc)
	return sel, nil
}

// EncodeType encodes the given inType using the TypedData struct.
//
// It takes in a string representing the inType and returns the encoded string and an error if any.
func (td TypedData) EncodeType(inType string) (enc string, err error) {
	var typeDefs TypeDef
	var ok bool
	if typeDefs, ok = td.Types[inType]; !ok {
		return enc, fmt.Errorf("can't parse type %s from types %v", inType, td.Types)
	}
	var buf bytes.Buffer
	customTypes := make(map[string]TypeDef)
	buf.WriteString(inType)
	buf.WriteString("(")
	for i, def := range typeDefs.Definitions {
		if def.Type != "felt" {
			var customTypeDef TypeDef
			if customTypeDef, ok = td.Types[def.Type]; !ok {
				return enc, fmt.Errorf("can't parse type %s from types %v", def.Type, td.Types)
			}
			customTypes[def.Type] = customTypeDef
		}
		buf.WriteString(fmt.Sprintf("%s:%s", def.Name, def.Type))
		if i != (len(typeDefs.Definitions) - 1) {
			buf.WriteString(",")
		}
	}
	buf.WriteString(")")

	for customTypeName, customType := range customTypes {
		buf.WriteString(fmt.Sprintf("%s(", customTypeName))
		for i, def := range customType.Definitions {
			buf.WriteString(fmt.Sprintf("%s:%s", def.Name, def.Type))
			if i != (len(customType.Definitions) - 1) {
				buf.WriteString(",")
			}
		}
		buf.WriteString(")")
	}
	return buf.String(), nil
}
