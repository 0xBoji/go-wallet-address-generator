package main

import (
	"encoding/json"
	"fmt"
	"log"
	"server-wallet-v2/pkg/wallet"
)

func main() {
	// Create a new wallet generator
	generator := wallet.NewGenerator()

	// Generate an EVM wallet
	evmWallet, err := generator.CreateEVM()
	if err != nil {
		log.Fatalf("Failed to create EVM wallet: %v", err)
	}
	printWallet("EVM Wallet", evmWallet)

	// Generate a Solana wallet
	solanaWallet, err := generator.CreateSolana()
	if err != nil {
		log.Fatalf("Failed to create Solana wallet: %v", err)
	}
	printWallet("Solana Wallet", solanaWallet)
}

func printWallet(title string, w *wallet.Wallet) {
	fmt.Printf("\n%s:\n", title)
	jsonData, _ := json.MarshalIndent(w, "", "  ")
	fmt.Println(string(jsonData))
}
