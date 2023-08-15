# pick-private

Generate private keys, public keys, addresses and scripts from numbers (decimal, binary or hex) and WIF.
Warning: you probably should **not** use this with actual BTC.

## Building

Clone repository, then run:

```
$ go build
```

Dependencies should be downloaded and the `pick-private` binary will be generated.

## Running

```
$ ./pick-private 1
Note: treating input key as decimal

[Raw private key]
Hex:
0000000000000000000000000000000000000000000000000000000000000001
Binary:
1
Decimal:
1

[Public key]
Uncompressed:
0479be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8
Hash:
91b24bf9f5288532960ac687abb035127b1d28a5

Compressed:
0279be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798
Hash:
751e76e8199196d454941c45d1b3a323f1433bd6

[Legacy uncompressed]
Address: 1EHNa6Q4Jz2uvNExL497mE43ikXhwF6kZm
Privkey: 5HpHagT65TZzG1PH3CSu63k8DbpvD8s5ip4nEB3kEsreAnchuDf
Script: 76a91491b24bf9f5288532960ac687abb035127b1d28a588ac

[Legacy compressed]
Address: 1BgGZ9tcN4rm9KBzDn7KprQz87SZ26SAMH
Privkey: KwDiBf89QgGbjEhKnhXJuH7LrciVrZi3qYjgd9M7rFU73sVHnoWn
Script: 76a914751e76e8199196d454941c45d1b3a323f1433bd688ac

[P2SH-Segwit]
Address: 3JvL6Ymt8MVWiCNHC7oWU6nLeHNJKLZGLN
Privkey: KwDiBf89QgGbjEhKnhXJuH7LrciVrZi3qYjgd9M7rFU73sVHnoWn
Script: a914bcfeb728b584253d5f3f70bcb780e9ef218a68f487

[SegWit]
Address: bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4
Privkey: KwDiBf89QgGbjEhKnhXJuH7LrciVrZi3qYjgd9M7rFU73sVHnoWn
Script: 0014751e76e8199196d454941c45d1b3a323f1433bd6
```

For testnet and other options:

```
$ ./pick-private -h
```

## Tests

To run the tests:

```
$ go test ./...
```
