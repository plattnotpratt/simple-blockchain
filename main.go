package main

import (
	"fmt"

	"github.com/plattnotpratt/simple-blockchain/blockchain"
)

func main() {
	chain := blockchain.NewBlockchain()

	chain.AddBlock("Alice sent 5 BTC to Bob")
	chain.AddBlock("Bob sent 2 BTC to Charlie")

	for _, block := range chain.Blocks {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Previous Hash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("Nonce: %d\n", block.Nonce)
	}
	fmt.Println("\nBlockchain valid?", chain.IsValid())

}
