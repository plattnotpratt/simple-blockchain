package blockchain

import (
	"fmt"
	"strings"
	"time"
)

const difficulty int = 4

type Blockchain struct {
	Blocks []Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]Block{CreateGenesisBlock()}}
}

func generateDifficulty(difficulty int) string {
	if difficulty <= 0 {
		return ""
	}
	return strings.Repeat("0", difficulty)
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

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := Block{
		Index:     len(bc.Blocks),
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevBlock.Hash,
	}
	newBlock.mineBlock(difficulty)
	bc.Blocks = append(bc.Blocks, newBlock)
}
