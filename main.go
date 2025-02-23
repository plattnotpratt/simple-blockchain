package main

import (
	"fmt"

	"github.com/plattnotpratt/simple-blockchain/blockchain"
)

func main() {
	chain := blockchain.NewBlockchain()

	chain.AddTransaction("Alice", "Bob", 10)
	chain.AddTransaction("Bob", "Charlie", 5)

	chain.MinePendingTransactions()

	for _, block := range chain.Blocks {
		fmt.Printf("\nBlock #%d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Previous Hash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		for _, txn := range block.Transactions {
			fmt.Printf("- %s sent %f BTC to %s\n", txn.Sender, txn.Amount, txn.Receiver)
		}
	}
	fmt.Println("\nBlockchain valid?", chain.IsValid())

	// Check balances
	fmt.Println("\nBalances:")
	fmt.Printf("Alice: %.2f BTC\n", chain.GetBalance("Alice"))
	fmt.Printf("Bob: %.2f BTC\n", chain.GetBalance("Bob"))
	fmt.Printf("Charlie: %.2f BTC\n", chain.GetBalance("Charlie"))

}
