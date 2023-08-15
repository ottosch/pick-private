package crypto

import (
	"crypto/sha256"

	"golang.org/x/crypto/ripemd160"
)

func Hash160(data []byte) []byte {
	sha256 := sha256.Sum256(data)
	ripe := ripemd160.New()
	ripe.Write(sha256[:])
	return ripe.Sum(nil)
}

func Hash256(data []byte) []byte {
	digest := sha256.Sum256(data)
	digest = sha256.Sum256(digest[:])
	return digest[:]
}
