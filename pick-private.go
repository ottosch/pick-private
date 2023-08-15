package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strings"

	"github.com/ottosch/pick-private/base58"
	"github.com/ottosch/pick-private/keys"
)

var (
	regexDecimal = regexp.MustCompile(`^\d+$`)
	regexBinary  = regexp.MustCompile(`^[01]+$`)
	regexHex     = regexp.MustCompile(`^[a-fA-F0-9]+$`)
	regexWif     = regexp.MustCompile(`^[KL59c][1-9a-km-zA-HJ-NP-Z]+$`)

	testnet    bool
	keyDecimal bool
	keyBinary  bool
	keyHex     bool
	keyWif     bool

	keyType  string
	inputKey string

	privateKey keys.PrivateKey
)

func main() {
	configCliArgs()
	parseCliArgs()
	parsePrivateKey()
	printOutput()
}

func configCliArgs() {
	flag.CommandLine.SetOutput(os.Stdout)
	flag.Usage = func() {
		fmt.Printf("Usage: %s [options] private key\n", os.Args[0])
		fmt.Println("\nOptions:")
		flag.PrintDefaults()

		fmt.Println("\nExamples:")
		fmt.Printf("  %s 1\n", os.Args[0])
		fmt.Printf("  %s -testnet deadbeef\n", os.Args[0])
		fmt.Printf("  %s 110001\n", os.Args[0])
		fmt.Printf("  %s -testnet -type hex 2222\n", os.Args[0])
		fmt.Printf("  %s KxR42n9vD54RcZgCvuaDgfbXfRGiJcpSfJMicjmaJzr7V17x5gXP2\n", os.Args[0])

	}
	flag.BoolVar(&testnet, "testnet", false, "generate testnet instead of mainnet")
	flag.StringVar(&keyType, "type", "", "force input into a specific type. Possible values: decimal [d], binary [b], hex [h] or wif [w]")
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
}

func parseCliArgs() {
	keyType = strings.ToLower(keyType)
	switch {
	case keyType == "decimal" || keyType == "d":
		keyDecimal = true
	case keyType == "binary" || keyType == "b":
		keyBinary = true
	case keyType == "hex" || keyType == "h":
		keyHex = true
	case keyType == "wif" || keyType == "w":
		keyWif = true
	case keyType == "":
		break
	default:
		fmt.Fprintf(os.Stderr, "unrecognized key type: %s\n", keyType)
		os.Exit(1)
	}

	inputKey = flag.Arg(0)
}

func parsePrivateKey() {
	if !keyDecimal && !keyBinary && !keyHex && !keyWif {
		switch {
		case regexBinary.MatchString(inputKey) && len(inputKey) >= 3:
			keyBinary = true
		case regexDecimal.MatchString(inputKey):
			keyDecimal = true
		case regexHex.MatchString(inputKey):
			keyHex = true
		case regexWif.MatchString(inputKey):
			keyWif = true
		}
	}

	var bigIntKey *big.Int
	note := "Note: treating input key as "
	switch {
	case keyDecimal:
		bigIntKey, _ = new(big.Int).SetString(inputKey, 10)
		note += "decimal"
	case keyBinary:
		bigIntKey, _ = new(big.Int).SetString(inputKey, 2)
		note += "binary"
	case keyHex:
		bigIntKey, _ = new(big.Int).SetString(inputKey, 16)
		note += "hex"
	case keyWif || regexWif.MatchString(inputKey):
		var err error
		if bigIntKey, err = base58.Decode(inputKey); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		note += "WIF"
	default:
		fmt.Fprintf(os.Stderr, "invalid private key: %s\n", inputKey)
		os.Exit(1)
	}

	fmt.Println(note)
	fmt.Println()
	privateKey = keys.FromBigInt(bigIntKey, testnet)
}

func printOutput() {
	fmt.Println("[Raw private key]")
	fmt.Println("Hex:")
	fmt.Println(fmt.Sprintf("%064x", privateKey.PrivateKey()))
	fmt.Println("Binary:")
	fmt.Println(fmt.Sprintf("%b", privateKey.PrivateKey()))
	fmt.Println("Decimal:")
	fmt.Println(privateKey.PrivateKey())
	fmt.Println()

	fmt.Println("[Public key]")
	fmt.Println("Uncompressed:")
	fmt.Println(hex.EncodeToString(privateKey.PublicKeyUncompressed()))
	fmt.Println("Hash:")
	fmt.Println(hex.EncodeToString(privateKey.ToPublicKeyHashUncompressed()))
	fmt.Println()

	fmt.Println("Compressed:")
	fmt.Println(hex.EncodeToString(privateKey.PublicKey()))
	fmt.Println("Hash:")
	fmt.Println(hex.EncodeToString(privateKey.ToPublicKeyHash()))
	fmt.Println()

	fmt.Println("[Legacy uncompressed]")
	fmt.Printf("Address: %s\n", privateKey.ToAddressLegacyUncompressed())
	fmt.Printf("Privkey: %s\n", privateKey.ToWIFUncompressed())
	fmt.Printf(" Script: %s\n", privateKey.ToScriptLegacyUncompressed())
	fmt.Println()

	fmt.Println("[Legacy compressed]")
	fmt.Printf("Address: %s\n", privateKey.ToAddressLegacy())
	fmt.Printf("Privkey: %s\n", privateKey.ToWIF())
	fmt.Printf(" Script: %s\n", privateKey.ToScriptLegacy())
	fmt.Println()

	fmt.Println("[P2SH-Segwit]")
	fmt.Printf("Address: %s\n", privateKey.ToAddressSegWitCompat())
	fmt.Printf("Privkey: %s\n", privateKey.ToWIF())
	fmt.Printf(" Script: %s\n", privateKey.ToScriptSegwitCompat())
	fmt.Println()

	fmt.Println("[SegWit]")
	fmt.Printf("Address: %s\n", privateKey.ToAddressSegWit())
	fmt.Printf("Privkey: %s\n", privateKey.ToWIF())
	fmt.Printf(" Script: %s\n", privateKey.ToScriptSegwit())
	fmt.Println()
}
