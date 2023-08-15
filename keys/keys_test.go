package keys_test

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/ottosch/pick-private/keys"
)

type network struct {
	name    string
	testnet bool
}

var networks = []network{
	{"mainnet", false},
	{"testnet", true},
}

type testData struct {
	input                     *big.Int
	pubkey                    string
	pubkeyUncompressed        string
	pubKeyHash                string
	pubKeyHashUncompressed    string
	wif                       string
	wifUncompressed           string
	addressLegacy             string
	addressLegacyUncompressed string
	scriptLegacy              string
	scriptLegacyUncompressed  string
	addressSegWitCompat       string
	scriptSegwitCompat        string
	addressSegWit             string
	scriptSegwit              string
}

var testsMainnet = []testData{
	{
		input:                     big.NewInt(1),
		pubkey:                    "0279be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798",
		pubkeyUncompressed:        "0479be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8",
		pubKeyHash:                "751e76e8199196d454941c45d1b3a323f1433bd6",
		pubKeyHashUncompressed:    "91b24bf9f5288532960ac687abb035127b1d28a5",
		wif:                       "KwDiBf89QgGbjEhKnhXJuH7LrciVrZi3qYjgd9M7rFU73sVHnoWn",
		wifUncompressed:           "5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreAnchuDf",
		addressLegacy:             "1BgGZ9tcN4rm9KBzDn7KprQz87SZ26SAMH",
		addressLegacyUncompressed: "1EHNa6Q4Jz2uvNExL497mE43ikXhwF6kZm",
		scriptLegacy:              "76a914751e76e8199196d454941c45d1b3a323f1433bd688ac",
		scriptLegacyUncompressed:  "76a91491b24bf9f5288532960ac687abb035127b1d28a588ac",
		addressSegWitCompat:       "3JvL6Ymt8MVWiCNHC7oWU6nLeHNJKLZGLN",
		scriptSegwitCompat:        "a914bcfeb728b584253d5f3f70bcb780e9ef218a68f487",
		addressSegWit:             "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
		scriptSegwit:              "0014751e76e8199196d454941c45d1b3a323f1433bd6",
	},
	{
		input:                     big.NewInt(12345),
		pubkey:                    "03f01d6b9018ab421dd410404cb869072065522bf85734008f105cf385a023a80f",
		pubkeyUncompressed:        "04f01d6b9018ab421dd410404cb869072065522bf85734008f105cf385a023a80f0eba29d0f0c5408ed681984dc525982abefccd9f7ff01dd26da4999cf3f6a295",
		pubKeyHash:                "1520f087720e1811802ded9bc38018da99111f90",
		pubKeyHashUncompressed:    "a42d4d68affbb92a4a733df0d5bf9375456921e5",
		wif:                       "KwDiBf89QgGbjEhKnhXJuH7LrciVrZi3qYjgd9M7rFVw2pgpVHKU",
		wifUncompressed:           "5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEss4BPiFsjb",
		addressLegacy:             "12vieiAHxBe4qCUrwvfb2kRkDuc8kQ2VZ2",
		addressLegacyUncompressed: "1Fy668EHkFwsrBQJfZsXYVgsGzKDaZhUEj",
		scriptLegacy:              "76a9141520f087720e1811802ded9bc38018da99111f9088ac",
		scriptLegacyUncompressed:  "76a914a42d4d68affbb92a4a733df0d5bf9375456921e588ac",
		addressSegWitCompat:       "3BEqJ8hzdNhtknPpkNQcB7VS86Vqm7qy5r",
		scriptSegwitCompat:        "a91468bd923bc087818d8f16418a4418741bef69b47b87",
		addressSegWit:             "bc1qz5s0ppmjpcvprqpdakdu8qqcm2v3z8us334at6",
		scriptSegwit:              "00141520f087720e1811802ded9bc38018da99111f90",
	},
	{
		input:                     hexToBigInt("9ae65d9154ac2490d7fb3f5e63d37d174a2e8d8a1744f9114f6486f315c08f06"),
		pubkey:                    "0200ec785d8cec2c2dac4272b5b86b97f6e6931131f03ddf301d20cd3f17279323",
		pubkeyUncompressed:        "0400ec785d8cec2c2dac4272b5b86b97f6e6931131f03ddf301d20cd3f17279323e5839e7730867dc506c4141f8ec6a5051e9e4b0b612d05493816a3247b28f050",
		pubKeyHash:                "56bb356338ca1e13b4abea563727fec7394a692d",
		pubKeyHashUncompressed:    "2ed46f46d5aa2147c19e96c02e7d2e5fde08e90a",
		wif:                       "L2QpHCj82EZdtYpYLwFpLbDjSaehhkZ4BX3x1mb67RzqGsV5biZG",
		wifUncompressed:           "5JzWHp4eqmJpKvt6k2sUwjeKgKxDPpvwBzZPS6a6gyQyhrmbJgF",
		addressLegacy:             "18ubNyZVhVSeZm8naxTuv8y9hFrgR2boWS",
		addressLegacyUncompressed: "15Gca3ZY6HCL1PZKX1JcCdtYfCD1HxVnFw",
		scriptLegacy:              "76a91456bb356338ca1e13b4abea563727fec7394a692d88ac",
		scriptLegacyUncompressed:  "76a9142ed46f46d5aa2147c19e96c02e7d2e5fde08e90a88ac",
		addressSegWitCompat:       "3PrWumDE8vR8JdiNPAMhPmez1Cgxxt4UVh",
		scriptSegwitCompat:        "a914f31eb5282c316e7dab6fd782021e555f577a577b87",
		addressSegWit:             "bc1q26an2ceceg0p8d9taftrwfl7cuu556fd6rr4jy",
		scriptSegwit:              "001456bb356338ca1e13b4abea563727fec7394a692d",
	},
}

var testsTestnet = []testData{
	{
		input:                     big.NewInt(1),
		pubkey:                    "0279be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798",
		pubkeyUncompressed:        "0479be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8",
		pubKeyHash:                "751e76e8199196d454941c45d1b3a323f1433bd6",
		pubKeyHashUncompressed:    "91b24bf9f5288532960ac687abb035127b1d28a5",
		wif:                       "cMahea7zqjxrtgAbB7LSGbcQUr1uX1ojuat9jZodMN87JcbXMTcA",
		wifUncompressed:           "91avARGdfge8E4tZfYLoxeJ5sGBdNJQH4kvjJoQFacbgwmaKkrx",
		addressLegacy:             "mrCDrCybB6J1vRfbwM5hemdJz73FwDBC8r",
		addressLegacyUncompressed: "mtoKs9V381UAhUia3d7Vb9GNak8Qvmcsme",
		scriptLegacy:              "76a914751e76e8199196d454941c45d1b3a323f1433bd688ac",
		scriptLegacyUncompressed:  "76a91491b24bf9f5288532960ac687abb035127b1d28a588ac",
		addressSegWitCompat:       "2NAUYAHhujozruyzpsFRP63mbrdaU5wnEpN",
		scriptSegwitCompat:        "a914bcfeb728b584253d5f3f70bcb780e9ef218a68f487",
		addressSegWit:             "tb1qw508d6qejxtdg4y5r3zarvary0c5xw7kxpjzsx",
		scriptSegwit:              "0014751e76e8199196d454941c45d1b3a323f1433bd6",
	},
	{
		input:                     big.NewInt(12345),
		pubkey:                    "03f01d6b9018ab421dd410404cb869072065522bf85734008f105cf385a023a80f",
		pubkeyUncompressed:        "04f01d6b9018ab421dd410404cb869072065522bf85734008f105cf385a023a80f0eba29d0f0c5408ed681984dc525982abefccd9f7ff01dd26da4999cf3f6a295",
		pubKeyHash:                "1520f087720e1811802ded9bc38018da99111f90",
		pubKeyHashUncompressed:    "a42d4d68affbb92a4a733df0d5bf9375456921e5",
		wif:                       "cMahea7zqjxrtgAbB7LSGbcQUr1uX1ojuat9jZodMN9wHZo77nNv",
		wifUncompressed:           "91avARGdfge8E4tZfYLoxeJ5sGBdNJQH4kvjJoQFacc6xVKZXV1",
		addressLegacy:             "mhSfwmFGmD5KcJxUfVdxrfe55uCqkptc6a",
		addressLegacyUncompressed: "mvV3PBKGZHP8dHsvP8quNQuC8yuvPnpoSE",
		scriptLegacy:              "76a9141520f087720e1811802ded9bc38018da99111f9088ac",
		scriptLegacyUncompressed:  "76a914a42d4d68affbb92a4a733df0d5bf9375456921e588ac",
		addressSegWitCompat:       "2N2o3Mse2EqDExa2NRW2Uo4UhLSi1Xf1Rwj",
		scriptSegwitCompat:        "a91468bd923bc087818d8f16418a4418741bef69b47b87",
		addressSegWit:             "tb1qz5s0ppmjpcvprqpdakdu8qqcm2v3z8usmhwwsf",
		scriptSegwit:              "00141520f087720e1811802ded9bc38018da99111f90",
	},
	{
		input:                     hexToBigInt("9ae65d9154ac2490d7fb3f5e63d37d174a2e8d8a1744f9114f6486f315c08f06"),
		pubkey:                    "0200ec785d8cec2c2dac4272b5b86b97f6e6931131f03ddf301d20cd3f17279323",
		pubkeyUncompressed:        "0400ec785d8cec2c2dac4272b5b86b97f6e6931131f03ddf301d20cd3f17279323e5839e7730867dc506c4141f8ec6a5051e9e4b0b612d05493816a3247b28f050",
		pubKeyHash:                "56bb356338ca1e13b4abea563727fec7394a692d",
		pubKeyHashUncompressed:    "2ed46f46d5aa2147c19e96c02e7d2e5fde08e90a",
		wif:                       "cSmok7iyTJFu3zHojM4whuio4ox7NCekFZCR8C3bcYeqXcfeD8TZ",
		wifUncompressed:           "92m8sYtCRzNxHzPPNNmPpLCHKzJvYzU8XwRLWivc2iA2UspcGoW",
		addressLegacy:             "moRYg2eUWWsuLscQJXSHk4BUZFTPKAnXpx",
		addressLegacyUncompressed: "mjnZs6eWuJdanW2wEaGz2Z6sXBoiEXY6QH",
		scriptLegacy:              "76a91456bb356338ca1e13b4abea563727fec7394a692d88ac",
		scriptLegacyUncompressed:  "76a9142ed46f46d5aa2147c19e96c02e7d2e5fde08e90a88ac",
		addressSegWitCompat:       "2NFQiyW9FkNvUWRLv4Hya1ieFDYu8hj5edj",
		scriptSegwitCompat:        "a914f31eb5282c316e7dab6fd782021e555f577a577b87",
		addressSegWit:             "tb1q26an2ceceg0p8d9taftrwfl7cuu556fds9cxfh",
		scriptSegwit:              "001456bb356338ca1e13b4abea563727fec7394a692d",
	},
}

func hexToBigInt(hex string) *big.Int {
	number, _ := new(big.Int).SetString(hex, 16)
	return number
}

func TestPublicKey(t *testing.T) {
	for _, network := range networks {
		var testCases []testData
		if network.testnet {
			testCases = testsTestnet
		} else {
			testCases = testsMainnet
		}
		for _, test := range testCases {
			privateKey := keys.FromBigInt(test.input, network.testnet)
			pubkey := hex.EncodeToString(privateKey.PublicKey())
			pubkeyUncompressed := hex.EncodeToString(privateKey.PublicKeyUncompressed())

			if pubkey != test.pubkey {
				t.Errorf("PublicKey for [%s] %d FAILED. Expected %s, got %s\n", network.name, test.input, test.pubkey, pubkey)
			} else {
				t.Logf("PublicKey passed: [%s] %d, %s\n", network.name, test.input, test.pubkey)
			}

			if pubkeyUncompressed != test.pubkeyUncompressed {
				t.Errorf("PublicKeyUncompressed for [%s] %d FAILED. Expected %s, got %s\n", network.name, test.input, test.pubkeyUncompressed, pubkeyUncompressed)
			} else {
				t.Logf("PublicKeyUncompressed passed: [%s] %d, %s\n", network.name, test.input, test.pubkeyUncompressed)
			}
		}
	}
}

func TestPublicKeyHash(t *testing.T) {
	for _, network := range networks {
		var testCases []testData
		if network.testnet {
			testCases = testsTestnet
		} else {
			testCases = testsMainnet
		}
		for _, test := range testCases {
			privateKey := keys.FromBigInt(test.input, network.testnet)
			pubkeyHash := hex.EncodeToString(privateKey.ToPublicKeyHash())
			pubkeyHashUncompressed := hex.EncodeToString(privateKey.ToPublicKeyHashUncompressed())

			if pubkeyHash != test.pubKeyHash {
				t.Errorf("ToPublicKeyHash for [%s] %d FAILED. Expected %s, got %s\n", network.name, test.input, test.pubKeyHash, pubkeyHash)
			} else {
				t.Logf("ToPublicKeyHash passed: [%s] %d, %s\n", network.name, test.input, test.pubKeyHash)
			}

			if pubkeyHashUncompressed != test.pubKeyHashUncompressed {
				t.Errorf("ToPublicKeyHashUncompressed for [%s] %d FAILED. Expected %s, got %s\n", network.name, test.input, test.pubKeyHashUncompressed, pubkeyHashUncompressed)
			} else {
				t.Logf("ToPublicKeyHashUncompressed passed: [%s] %d, %s\n", network.name, test.input, test.pubKeyHashUncompressed)
			}
		}
	}
}

func TestWIF(t *testing.T) {
	for _, network := range networks {
		var testCases []testData
		if network.testnet {
			testCases = testsTestnet
		} else {
			testCases = testsMainnet
		}
		for _, test := range testCases {
			privateKey := keys.FromBigInt(test.input, network.testnet)
			wif := privateKey.ToWIF()
			wifUncompressed := privateKey.ToWIFUncompressed()

			if wif != test.wif {
				t.Errorf("ToWIF for [%s] %d FAILED. Expected %s, got %s\n", network.name, test.input, test.wif, wif)
			} else {
				t.Logf("ToWIF passed: [%s] %d, %s\n", network.name, test.input, test.wif)
			}

			if wifUncompressed != test.wifUncompressed {
				t.Errorf("ToWIFUncompressed for [%s] %d FAILED. Expected %s, got %s\n", network.name, test.input, test.wifUncompressed, wifUncompressed)
			} else {
				t.Logf("ToWIFUncompressed passed: [%s] %d, %s\n", network.name, test.input, test.wifUncompressed)
			}
		}
	}
}

func TestAddress(t *testing.T) {
	for _, network := range networks {
		var testCases []testData
		if network.testnet {
			testCases = testsTestnet
		} else {
			testCases = testsMainnet
		}
		for _, test := range testCases {
			privateKey := keys.FromBigInt(test.input, network.testnet)
			addressLegacy := privateKey.ToAddressLegacy()
			addressLegacyUncompressed := privateKey.ToAddressLegacyUncompressed()
			addressSegWitCompat := privateKey.ToAddressSegWitCompat()
			addressSegWit := privateKey.ToAddressSegWit()

			if addressLegacy != test.addressLegacy {
				t.Errorf("ToAddressLegacy for [%s] %d FAILED. Expected %s, got %s\n", network.name, test.input, test.addressLegacy, addressLegacy)
			} else {
				t.Logf("ToAddressLegacy passed: [%s] %d, %s\n", network.name, test.input, test.addressLegacy)
			}

			if addressLegacyUncompressed != test.addressLegacyUncompressed {
				t.Errorf("ToAddressLegacyUncompressed for [%s] %d FAILED. Expected %s, got %s\n", network.name, test.input, test.addressLegacyUncompressed, addressLegacyUncompressed)
			} else {
				t.Logf("ToAddressLegacyUncompressed passed: [%s] %d, %s\n", network.name, test.input, test.addressLegacyUncompressed)
			}

			if addressSegWitCompat != test.addressSegWitCompat {
				t.Errorf("ToAddressSegWitCompat for [%s] %d FAILED. Expected %s, got %s\n", network.name, test.input, test.addressSegWitCompat, addressSegWitCompat)
			} else {
				t.Logf("ToAddressSegWitCompat passed: [%s] %d, %s\n", network.name, test.input, test.addressSegWitCompat)
			}

			if addressSegWit != test.addressSegWit {
				t.Errorf("ToAddressSegWit for [%s] %d FAILED. Expected %s, got %s\n", network.name, test.input, test.addressSegWit, addressSegWit)
			} else {
				t.Logf("ToAddressSegWit passed: [%s] %d, %s\n", network.name, test.input, test.addressSegWit)
			}
		}
	}
}

func TestScript(t *testing.T) {
	for _, network := range networks {
		var testCases []testData
		if network.testnet {
			testCases = testsTestnet
		} else {
			testCases = testsMainnet
		}
		for _, test := range testCases {
			privateKey := keys.FromBigInt(test.input, network.testnet)
			scriptLegacy := privateKey.ToScriptLegacy()
			scriptLegacyUncompressed := privateKey.ToScriptLegacyUncompressed()
			scriptSegwitCompat := privateKey.ToScriptSegwitCompat()
			scriptSegwit := privateKey.ToScriptSegwit()

			if scriptLegacy != test.scriptLegacy {
				t.Errorf("ToScriptLegacy for [%s] %d FAILED. Expected %s, got %s\n", network.name, test.input, test.scriptLegacy, scriptLegacy)
			} else {
				t.Logf("ToScriptLegacy passed: [%s] %d, %s\n", network.name, test.input, test.scriptLegacy)
			}

			if scriptLegacyUncompressed != test.scriptLegacyUncompressed {
				t.Errorf("ToScriptLegacyUncompressed for [%s] %d FAILED. Expected %s, got %s\n", network.name, test.input, test.scriptLegacyUncompressed, scriptLegacyUncompressed)
			} else {
				t.Logf("ToScriptLegacyUncompressed passed: [%s] %d, %s\n", network.name, test.input, test.scriptLegacyUncompressed)
			}

			if scriptSegwitCompat != test.scriptSegwitCompat {
				t.Errorf("ToScriptSegwitCompat for [%s] %d FAILED. Expected %s, got %s\n", network.name, test.input, test.scriptSegwitCompat, scriptSegwitCompat)
			} else {
				t.Logf("ToScriptSegwitCompat passed: [%s] %d, %s\n", network.name, test.input, test.scriptSegwitCompat)
			}

			if scriptSegwit != test.scriptSegwit {
				t.Errorf("ToScriptSegwit for [%s] %d FAILED. Expected %s, got %s\n", network.name, test.input, test.scriptSegwit, scriptSegwit)
			} else {
				t.Logf("ToScriptSegwit passed: [%s] %d, %s\n", network.name, test.input, test.scriptSegwit)
			}
		}
	}
}
