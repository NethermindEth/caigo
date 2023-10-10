package curve

import (
	"crypto/hmac"
	"hash"
	"math/big"
)

// DivMod calculates the quotient and remainder of a division operation between two big integers (0 <= x < p such that (m * x) % p == n).
// (ref: https://github.com/starkware-libs/cairo-lang/blob/master/src/starkware/crypto/starkware/crypto/signature/math_utils.py)
//
// The function takes three parameters:
// - n: a pointer to a big integer representing the dividend
// - m: a pointer to a big integer representing the divisor
// - p: a pointer to a big integer representing the modulus
//
// The function returns a pointer to a big integer representing the remainder of the division operation.
func DivMod(n, m, p *big.Int) *big.Int {
	q := new(big.Int)
	gx := new(big.Int)
	gy := new(big.Int)
	q.GCD(gx, gy, m, p)

	r := new(big.Int).Mul(n, gx)
	r = r.Mod(r, p)
	return r
}

// int2octets returns a byte array representation of a big integer in octets format.
// https://tools.ietf.org/html/rfc6979#section-2.3.3
//
// It takes a pointer to a big.Int and the length of the resulting byte array as input.
// The function pads the byte array with zeros if it is shorter than the specified length.
// If the byte array is longer than the specified length, it drops the most significant bytes.
// The function returns the resulting byte array.
func int2octets(v *big.Int, rolen int) []byte {
	out := v.Bytes()

	// pad with zeros if it's too short
	if len(out) < rolen {
		out2 := make([]byte, rolen)
		copy(out2[rolen-len(out):], out)
		return out2
	}

	// drop most significant bytes if it's too long
	if len(out) > rolen {
		out2 := make([]byte, rolen)
		copy(out2, out[len(out)-rolen:])
		return out2
	}

	return out
}

// bits2octets generates octets from bits.
// https://tools.ietf.org/html/rfc6979#section-2.3.4
//
// It takes in two big integers, 'in' and 'q', along with the integer values
// 'qlen' and 'rolen'. It converts the 'in' big integer to an octet and 
// subtracts 'q' from it. If the result is negative, the 'in' big integer is 
// converted to an octet and returned. Otherwise, the result of the subtraction 
// is converted to an octet and returned.
//
// It returns a byte slice representing the octets.
func bits2octets(in, q *big.Int, qlen, rolen int) []byte {
	z1 := bits2int(in, qlen)
	z2 := new(big.Int).Sub(z1, q)
	if z2.Sign() < 0 {
		return int2octets(z1, rolen)
	}
	return int2octets(z2, rolen)
}


// bits2int converts a big.Int representing a bit string to an integer.
// https://tools.ietf.org/html/rfc6979#section-2.3.2
//
// It takes two parameters: 
//    - in: a pointer to a big.Int representing the input bit string.
//    - qlen: an integer representing the desired length of the output integer.
// 
// It returns a pointer to a big.Int representing the converted integer.
func bits2int(in *big.Int, qlen int) *big.Int {
	blen := len(in.Bytes()) * 8

	if blen > qlen {

		return new(big.Int).Rsh(in, uint(blen-qlen))
	}
	return in
}


// mac calculates the message authentication code (MAC) using the provided hash algorithm,
// key, message, and buffer.
// mac returns an HMAC of the given key and message.
//
// The `alg` parameter is a function that returns a hash.Hash implementation. It is used to
// specify the hash algorithm to be used for calculating the MAC.
//
// The `k` parameter is a byte slice representing the key to be used for the MAC calculation.
//
// The `m` parameter is a byte slice representing the message for which the MAC is calculated.
//
// The `buf` parameter is a byte slice used as a buffer for storing the result of the MAC
// calculation.
//
// The function returns a byte slice containing the calculated MAC.
func mac(alg func() hash.Hash, k, m, buf []byte) []byte {
	h := hmac.New(alg, k)
	h.Write(m)
	return h.Sum(buf[:0])
}

// MaskBits masks the specified (excess) bits in a byte slice.
//
// The function takes three parameters: mask, wordSize, and slice.
// - mask is an integer representing the number of bits to mask.
// - wordSize is an integer representing the number of bits in each element of the slice.
// - slice is a byte slice on which the masking operation is performed.
//
// The function returns a new byte slice that contains the masked bits.
func MaskBits(mask, wordSize int, slice []byte) (ret []byte) {
	excess := len(slice)*wordSize - mask
	for _, by := range slice {
		if excess > 0 {
			if excess > wordSize {
				excess = excess - wordSize
				continue
			}
			by <<= excess
			by >>= excess
			excess = 0
		}
		ret = append(ret, by)
	}
	return ret
}

// FmtKecBytes formats a big integer into a Keccak hash.
//
// It takes a *big.Int as input and an integer rolen to specify the length of the
// output byte slice. The function returns a byte slice that contains the bytes
// representation of the input big integer.
func FmtKecBytes(in *big.Int, rolen int) (buf []byte) {
	buf = append(buf, in.Bytes()...)

	// pad with zeros if too short
	if len(buf) < rolen {
		padded := make([]byte, rolen)
		copy(padded[rolen-len(buf):], buf)

		return padded
	}

	return buf
}
