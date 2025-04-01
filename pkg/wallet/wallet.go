package wallet

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gagliardetto/solana-go"
)

// Wallet represents a blockchain wallet with its address and private key
type Wallet struct {
	Address     string
	PrivateKey  string
	PublicKey   string
	NetworkType string
}

// Generator provides methods to generate different types of blockchain wallets
type Generator struct{}

// NewGenerator creates a new wallet generator
func NewGenerator() *Generator {
	return &Generator{}
}

// CreateEVM generates a new EVM (Ethereum) wallet
func (g *Generator) CreateEVM() (*Wallet, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, fmt.Errorf("failed to generate private key: %w", err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyHex := hex.EncodeToString(privateKeyBytes)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to get public key")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return &Wallet{
		Address:     address,
		PrivateKey:  privateKeyHex,
		NetworkType: "evm",
	}, nil
}

// CreateSolana generates a new Solana wallet
func (g *Generator) CreateSolana() (*Wallet, error) {
	account := solana.NewWallet()

	return &Wallet{
		Address:     account.PublicKey().String(),
		PrivateKey:  account.PrivateKey.String(),
		NetworkType: "solana",
	}, nil
}