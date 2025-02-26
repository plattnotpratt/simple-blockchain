package blockchain

import (
	"testing"
	"time"
)

func testCalculateHash(t *testing.T) {
	dt := time.Now().String()
	b := Block{
		Index:        1,
		Timestamp:    dt,
		Transactions: []Transaction{},
		PrevHash:     "0",
		Nonce:        3000,
	}
	checkHash := b.CalculateHash()
	b.Index = 2
	if checkHash != b.CalculateHash() {
		t.Errorf("Hash not consistent and would always be invalid: got %s expected %s", checkHash[7:], b.CalculateHash()[7:])
	}

}
