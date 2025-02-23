package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"time"
)

type Transaction struct {
	Sender   string
	Receiver string
	Amount   float64
}

type Block struct {
	Index        int
	Timestamp    string
	Transactions []Transaction
	PrevHash     string
	Hash         string
	Nonce        int
}

func (b *Block) CalculateHash() string {
	txnByte, _ := json.Marshal(b.Transactions)
	record := strconv.Itoa(b.Index) + b.Timestamp + string(txnByte) + b.PrevHash + strconv.Itoa(b.Nonce)
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func CreateGenesisBlock() Block {
	genesis := Block{
		Index:        0,
		Timestamp:    time.Now().String(),
		Transactions: []Transaction{},
		PrevHash:     "0",
	}
	genesis.Hash = genesis.CalculateHash()
	return genesis
}
