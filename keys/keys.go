package keys

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/decred/dcrd/dcrec/secp256k1"
	"github.com/ottosch/pick-private/base58"
	"github.com/ottosch/pick-private/bech32"
	"github.com/ottosch/pick-private/crypto"
)

type PrivateKey struct {
	privKey *big.Int
	pubkey  []byte
	testnet bool
}

// FromBigInt creates a PrivateKey from a big.Int.
func FromBigInt(number *big.Int, testnet bool) PrivateKey {
	secPrivKey, _ := secp256k1.PrivKeyFromBytes(number.Bytes())
	x, y := secPrivKey.Public()
	pubkey := make([]byte, 64)
	x.FillBytes(pubkey[:32])
	y.FillBytes(pubkey[32:])
	return PrivateKey{number, pubkey, testnet}
}

// PrivateKey returns the internal *big.Int private key.
func (priv *PrivateKey) PrivateKey() *big.Int {
	return priv.privKey
}

// PublicKey returns the compressed public key.
func (priv *PrivateKey) PublicKey() []byte {
	return priv.publicKey(true)
}

// PublicKeyUncompressed returns the uncompressed public key.
func (priv *PrivateKey) PublicKeyUncompressed() []byte {
	return priv.publicKey(false)
}

func (priv *PrivateKey) publicKey(compressed bool) []byte {
	xBytes := priv.pubkey[:32]
	yBytes := priv.pubkey[32:]

	var pubkey []byte
	if compressed {
		if priv.pubkey[63]%2 == 0 {
			pubkey = []byte{0x02}
		} else {
			pubkey = []byte{0x03}
		}

		pubkey = append(pubkey, xBytes...)
	} else {
		pubkey = []byte{0x04}
		pubkey = append(pubkey, xBytes...)
		pubkey = append(pubkey, yBytes...)
	}

	return pubkey
}

// ToWIF returns the private key in WIF (compressed public key)
func (priv *PrivateKey) ToWIF() string {
	return priv.toWif(true)
}

// ToWIFUncompressed returns the private key in WIF (uncompressed public key)
func (priv *PrivateKey) ToWIFUncompressed() string {
	return priv.toWif(false)
}

func (priv *PrivateKey) toWif(compressed bool) string {
	hexPriv := priv.privKey.Bytes()
	hexPadded := fmt.Sprintf("%064x", hexPriv)

	var extended string
	switch {
	case priv.testnet && compressed:
		extended = fmt.Sprintf("EF%s01", hexPadded)
	case priv.testnet && !compressed:
		extended = fmt.Sprintf("EF%s", hexPadded)
	case !priv.testnet && compressed:
		extended = fmt.Sprintf("80%s01", hexPadded)
	case !priv.testnet && !compressed:
		extended = fmt.Sprintf("80%s", hexPadded)
	}

	hexExtended, _ := hex.DecodeString(extended)
	checksum := crypto.Hash256(hexExtended)[:4]
	checksumString := hex.EncodeToString(checksum[:])
	return base58.Encode(extended + checksumString)
}

// ToLegacy returns the legacy address (compressed public key)
func (priv *PrivateKey) ToAddressLegacy() string {
	return priv.p2pkh(true)
}

// ToLegacyUncompressed returns the legacy address (uncompressed public key)
func (priv *PrivateKey) ToAddressLegacyUncompressed() string {
	return priv.p2pkh(false)
}

// ToScriptLegacy returns the P2PKH scriptPubKey (compressed public key)
func (priv *PrivateKey) ToScriptLegacy() string {
	return priv.legacyScript(true)
}

// ToScriptLegacyUncompressed returns the P2PKH scriptPubKey (uncompressed public key)
func (priv *PrivateKey) ToScriptLegacyUncompressed() string {
	return priv.legacyScript(false)
}

func (priv *PrivateKey) legacyScript(compressed bool) string {
	pkh := hex.EncodeToString(priv.pkh(compressed))
	return fmt.Sprintf("76a914%s88ac", pkh)
}

// ToScriptSegwitCompat returns the P2SH-P2WPKH scriptPubKey
func (priv *PrivateKey) ToScriptSegwitCompat() string {
	redeem := []byte{0x00, 0x14}
	redeem = append(redeem, crypto.Hash160(priv.PublicKey())...)
	hash160Redeem := hex.EncodeToString(crypto.Hash160(redeem))
	return fmt.Sprintf("a914%s87", hash160Redeem)
}

// ToScriptSegwit returns the P2WPKH scriptPubKey
func (priv *PrivateKey) ToScriptSegwit() string {
	pkh := hex.EncodeToString(priv.pkh(true))
	return fmt.Sprintf("0014%s", pkh)
}

// ToPublicKeyHash returns the (compressed) public key hash
func (priv *PrivateKey) ToPublicKeyHash() []byte {
	return priv.pkh(true)
}

// ToPublicKeyHashUncompressed returns the (uncompressed) public key hash
func (priv *PrivateKey) ToPublicKeyHashUncompressed() []byte {
	return priv.pkh(false)
}

func (priv *PrivateKey) pkh(compressed bool) []byte {
	return crypto.Hash160(priv.publicKey(compressed))
}

func (priv *PrivateKey) p2pkh(compressed bool) string {
	hash160 := priv.pkh(compressed)
	data := []byte{0x00}
	if priv.testnet {
		data[0] = 0x6F
	}

	extendedHash160 := append(data, hash160...)
	checksum := crypto.Hash256(extendedHash160)[0:4]
	final := append(extendedHash160, checksum...)

	return base58.Encode(hex.EncodeToString(final))
}

// ToSegWitCompat the P2SH-SegWit address
func (priv *PrivateKey) ToAddressSegWitCompat() string {
	pubkey := priv.PublicKey()

	redeem := []byte{0x00, 0x14}
	redeem = append(redeem, crypto.Hash160(pubkey)...)

	hash160Redeem := crypto.Hash160(redeem)
	data := []byte{0x05}
	if priv.testnet {
		data[0] = 0xC4
	}

	extendedHash160 := append(data, hash160Redeem...)
	checksum := crypto.Hash256(extendedHash160)[0:4]
	final := append(extendedHash160, checksum...)

	return base58.Encode(hex.EncodeToString(final))
}

// ToSegWit the P2SH-SegWit address
func (priv *PrivateKey) ToAddressSegWit() string {
	hash160PubKey := crypto.Hash160(priv.PublicKey())

	program := make([]int, len(hash160PubKey))
	for i, b := range hash160PubKey {
		program[i] = int(b)
	}

	hrp := "bc"
	if priv.testnet {
		hrp = "tb"
	}

	addr, _ := bech32.SegwitAddrEncode(hrp, 0, program)
	return addr
}
