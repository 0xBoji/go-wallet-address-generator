# Blockchain Wallet Generator

A Go library for generating wallets for different networks (EVM, Solana).

## Installation

```bash
go get github.com/0xboji/go-wallet-address-generator
```

## Usage

```go
package main

import (
    "go-wallet-address-generator/pkg/wallet"
)

func main() {
    // Create a new wallet generator
    generator := wallet.NewGenerator()

    // Generate an EVM wallet
    evmWallet, err := generator.CreateEVM()
    if err != nil {
        log.Fatal(err)
    }

    // Generate a Solana wallet
    solanaWallet, err := generator.CreateSolana()
    if err != nil {
        log.Fatal(err)
    }
}
```

## Supported Networks

- EVM (Ethereum and compatible chains)
- Solana

## Wallet Structure

Each wallet contains:
- `Address`: The wallet's public address
- `PrivateKey`: The wallet's private key
- `PublicKey`: The wallet's public key (where applicable)
- `NetworkType`: The blockchain network type

## Example Output

```json
{
  "address": "0x...",
  "privateKey": "...",
  "publicKey": "...",
  "networkType": "evm"
}
```

## License

MIT 