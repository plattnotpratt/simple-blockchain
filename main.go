package main

import (
	"fmt"

	"github.com/plattnotpratt/simple-blockchain/blockchain"
)

func main() {
	genesisBlock := blockchain.CreateGenesisBlock()
	fmt.Println("Genesis Block:")
	fmt.Printf("Index: %d\n", genesisBlock.Index)
	fmt.Printf("Timestamp: %s\n", genesisBlock.Timestamp)
	fmt.Printf("Data: %s\n", genesisBlock.Data)
	fmt.Printf("Previous Hash: %s\n", genesisBlock.PrevHash)
	fmt.Printf("Hash: %s\n", genesisBlock.Hash)
}
