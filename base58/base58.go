package base58

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/ottosch/pick-private/crypto"
)

const alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

// Encode encodes argument string to base58check
func Encode(hexString string) string {
	decimalData := new(big.Int)
	decimalData, _ = decimalData.SetString(hexString, 16)

	divisor, zero := big.NewInt(58), big.NewInt(0)

	var encoded string
	for decimalData.Cmp(zero) > 0 {
		mod := new(big.Int)
		decimalData.DivMod(decimalData, divisor, mod)
		encoded = string(alphabet[mod.Int64()]) + encoded
	}

	var leadingZeros int
	for _, char := range hexString {
		if char == '0' {
			leadingZeros++
		} else {
			break
		}
	}

	return strings.Repeat("1", leadingZeros/2) + encoded
}

// Decode decodes argument WIF (compressed or uncompressed) to big.Int private key.
func Decode(wif string) (*big.Int, error) {
	compressed := wif[0] == 'K' || wif[0] == 'L' || wif[0] == 'c'

	multiplier := big.NewInt(58)
	total := new(big.Int)
	for _, c := range wif {
		if !strings.ContainsRune(alphabet, c) {
			err := fmt.Errorf("invalid character in WIF: %c\n", c)
			return new(big.Int), err
		}

		total.Mul(total, multiplier)
		digit := int64(strings.IndexRune(alphabet, c))
		total.Add(total, big.NewInt(digit))
	}

	if err := verifyChecksum(total.Bytes()); err != nil {
		return new(big.Int), err
	}

	bytes := total.Bytes()[1 : len(total.Bytes())-4]
	if compressed {
		bytes = bytes[:len(bytes)-1]
	}

	privBigInt, _ := new(big.Int).SetString(hex.EncodeToString(bytes), 16)
	return privBigInt, nil
}

func verifyChecksum(private []byte) error {
	if len(private) <= 4 {
		return errors.New("invalid WIF: too short")
	}

	extendedPrivate := private[:len(private)-4]

	inputChecksum := private[len(private)-4:]
	expectedChecksum := crypto.Hash256(extendedPrivate)[:4]

	for i := range inputChecksum {
		if inputChecksum[i] != expectedChecksum[i] {
			return fmt.Errorf("invalid WIF checksum, expected %x, got %x\n", expectedChecksum, inputChecksum)
		}
	}

	return nil
}
