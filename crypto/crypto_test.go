package crypto_test

import (
	"encoding/hex"
	"testing"

	"github.com/ottosch/pick-private/crypto"
)

type testData struct {
	input  string
	output string
}

var hash160Tests = []testData{
	{"0279be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798", "751e76e8199196d454941c45d1b3a323f1433bd6"},
	{"02c6047f9441ed7d6d3045406e95c07cd85c778e4b8cef3ca7abac09b95c709ee5", "06afd46bcdfd22ef94ac122aa11f241244a37ecc"},
	{"02f9308a019258c31049344f85f89d5229b531c845836f99b08601f113bce036f9", "7dd65592d0ab2fe0d0257d571abf032cd9db93dc"},
	{"02e493dbf1c10d80f3581e4904930b1404cc6c13900ee0758474fa94abe8c4cd13", "c42e7ef92fdb603af844d064faad95db9bcdfd3d"},
}

var hash256Tests = []testData{
	{"0041a8c88523c0323a6c577d329aa4373c42d1556c", "0639ceeaa769780a1855bdbcadb5b74ccb0a34858b61ef0f8cd3b06b82eea122"},
	{"00ff9f0151e96cf8e1773765907f5a066cb9971d4f", "c6154de71457a23f88380b2c60840411da318bcc59888480fe9720416a423e1c"},
	{"800000000000000000000000000000000000000000000000000000000000000014", "abcb919d67cc751c62a7cc095088342f4a12ebeb7e11c993a88b4d272de23b83"},
	{"05910473597fa5a7289e29a98bccb87be363606709", "0b94a623d50876c16cf5b84265473666b78891fae0d3e052c45af37076be6714"},
}

func TestHash160(t *testing.T) {
	for _, test := range hash160Tests {
		input, _ := hex.DecodeString(test.input)
		result := hex.EncodeToString((crypto.Hash160(input)))
		if result != test.output {
			t.Errorf("Hash160 for %s FAILED. Expected %s, got %s\n", test.input, test.output, result)
		} else {
			t.Logf("Hash160 passed: %s, %s\n", test.input, test.output)
		}
	}
}

func TestHash256(t *testing.T) {
	for _, test := range hash256Tests {
		input, _ := hex.DecodeString(test.input)
		result := hex.EncodeToString((crypto.Hash256(input)))
		if result != test.output {
			t.Errorf("Hash256 for %s FAILED. Expected %s, got %s\n", test.input, test.output, result)
		} else {
			t.Logf("Hash256 passed: %s, %s\n", test.input, test.output)
		}
	}
}
