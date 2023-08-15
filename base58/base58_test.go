package base58_test

import (
	"math/big"
	"testing"

	"github.com/ottosch/pick-private/base58"
)

type encodeTestData struct {
	input  string
	output string
}

type decodeTestData struct {
	input  string
	output int64
}

var encodeTests = []encodeTestData{
	{"11", "J"},
	{"aaaaaa", "zKrZ"},
	{"C1C44F4876C8C0FB72152462968432FC003B5A0EB9B8AD4C2CCD3451BA9457C8", "E3PJaAPGyx9upBRGwjRtAgbrayPvnCHLJxsnhitkXMVZ"},
	{"800000000000000000000000000000000000000000000000000000000000000001014671fc3f", "KwDiBf89QgGbjEhKnhXJuH7LrciVrZi3qYjgd9M7rFU73sVHnoWn"},
	{"809ae65d9154ac2490d7fb3f5e63d37d174a2e8d8a1744f9114f6486f315c08f06010ac19d23", "L2QpHCj82EZdtYpYLwFpLbDjSaehhkZ4BX3x1mb67RzqGsV5biZG"},
	{"809ae65d9154ac2490d7fb3f5e63d37d174a2e8d8a1744f9114f6486f315c08f06f1ffc4f8", "5JzWHp4eqmJpKvt6k2sUwjeKgKxDPpvwBzZPS6a6gyQyhrmbJgF"},
	{"ef000000000000000000000000000000000000000000000000000000000000000201e7102cb5", "cMahea7zqjxrtgAbB7LSGbcQUr1uX1ojuat9jZodMN87K7XCyj5v"},
}

var decodeTests = []decodeTestData{
	{"KwDiBf89QgGbjEhKnhXJuH7LrciVrZi3qYjgd9M7rFU73sVHnoWn", 1},
	{"cMahea7zqjxrtgAbB7LSGbcQUr1uX1ojuat9jZodMN87K7XCyj5v", 2},
	{"5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kt87rU1oi9ho", 0xdeadbeef},
	{"91avARGdfge8E4tZfYLoxeJ5sGBdNJQH4kvjJoQFacbgx3cTMqe", 3},
}

var decodeTestsInvalidCharacter = []string{
	"OWDiBf89QgGbjEhKnhXJuH7LrciVrZi3qYjgd9M7rFU73sVHnoWn",
	"lMahea7zqjxrtgAbB7LSGbcQUr1uX1ojuat9jZodMN87K7XCyj5V",
	"IHpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kt87rU1oi9ao",
}

var decodeTestsInvalidChecksum = []string{
	"KWDiBf89QgGbjEhKnhXJuH7LrciVrZi3qYjgd9M7rFU73sVHnoWn",
	"cMahea7zqjxrtgAbB7LSGbcQUr1uX1ojuat9jZodMN87K7XCyj5V",
	"5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kt87rU1oi9ao",
}

var decodeTestsInvalidShortInput = []string{
	"noWn",
	"j5V",
	"ao",
}

func TestEncodeOk(t *testing.T) {
	for _, test := range encodeTests {
		result := base58.Encode(test.input)
		if result != test.output {
			t.Errorf("Encode for %s FAILED. Expected %s, got %s\n", test.input, test.output, result)
		} else {
			t.Logf("Encode passed: %s, %s\n", test.input, test.output)
		}
	}
}

func TestDecodeOk(t *testing.T) {
	for _, test := range decodeTests {
		result, err := base58.Decode(test.input)

		switch {
		case err != nil:
			t.Errorf("Decode for %s FAILED: %v\n", test.input, err)
		case result.Cmp(big.NewInt(test.output)) != 0:
			t.Errorf("Decode for %s FAILED. Expected %d, got %s\n", test.input, test.output, result)
		default:
			t.Logf("Decode passed: %s, %d\n", test.input, test.output)
		}
	}
}

func TestDecodeInvalidChecksum(t *testing.T) {
	for _, test := range decodeTestsInvalidChecksum {
		_, err := base58.Decode(test)
		if err == nil {
			t.Errorf("Decode for %s passed, should've failed due to checksum: FAIL\n", test)
		} else {
			t.Logf("Decode for %s failed: %v\n", test, err)
		}
	}
}

func TestDecodeInvalidCharacter(t *testing.T) {
	for _, test := range decodeTestsInvalidCharacter {
		_, err := base58.Decode(test)
		if err == nil {
			t.Errorf("Decode for %s passed, should've failed due to invalid character: FAIL\n", test)
		} else {
			t.Logf("Decode for %s failed: %v\n", test, err)
		}
	}
}

func TestDecodeInvalidShortInput(t *testing.T) {
	for _, test := range decodeTestsInvalidShortInput {
		_, err := base58.Decode(test)
		if err == nil {
			t.Errorf("Decode for %s passed, should've failed due to input too short: FAIL\n", test)
		} else {
			t.Logf("Decode for %s failed: %v\n", test, err)
		}
	}
}
