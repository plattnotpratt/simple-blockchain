package blockchain

import (
	"fmt"
	"strings"
	"time"
)

const difficulty int = 4

type Blockchain struct {
	Blocks      []Block
	PendingTxns []Transaction
}

func generateDifficulty(difficulty int) string {
	if difficulty <= 0 {
		return ""
	}
	return strings.Repeat("0", difficulty)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]Block{CreateGenesisBlock()}, []Transaction{}}
}

func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		prevBlock := bc.Blocks[i-1]
		if currentBlock.Hash != currentBlock.CalculateHash() {
			return false
		}
		if currentBlock.PrevHash != prevBlock.Hash {
			return false
		}
	}
	return true
}

func (bc *Blockchain) AddTransaction(sender, receiver string, amount float64) {
	txn := Transaction{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
	}
	bc.PendingTxns = append(bc.PendingTxns, txn)
}

func (bc *Blockchain) MinePendingTransactions() {
	if len(bc.PendingTxns) == 0 {
		return
	}
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := Block{
		Index:        len(bc.Blocks),
		Timestamp:    time.Now().String(),
		Transactions: bc.PendingTxns,
		PrevHash:     prevBlock.Hash,
	}
	newBlock.mineBlock(difficulty)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (b *Block) mineBlock(difficulty int) {
	difficultyString := generateDifficulty(difficulty)
	for {
		b.Hash = b.CalculateHash()
		if b.Hash[:difficulty] == difficultyString {
			fmt.Println("Block mined:", b.Hash)
			return
		}
		//fmt.Print(b.Nonce, ", ")
		b.Nonce++
	}
}

func (bc *Blockchain) GetBalance(user string) float64 {
	balance := 0.0

	for _, block := range bc.Blocks {
		for _, txn := range block.Transactions {
			if txn.Sender == user {
				balance -= txn.Amount
			}
			if txn.Receiver == user {
				balance += txn.Amount
			}
		}
	}
	return balance
}

// func (bc *Blockchain) AddBlock(data string) {
// 	prevBlock := bc.Blocks[len(bc.Blocks)-1]
// 	newBlock := Block{
// 		Index:     len(bc.Blocks),
// 		Timestamp: time.Now().String(),
// 		Transactions:      data,
// 		PrevHash:  prevBlock.Hash,
// 	}
// 	newBlock.mineBlock(difficulty)
// 	bc.Blocks = append(bc.Blocks, newBlock)
// }
