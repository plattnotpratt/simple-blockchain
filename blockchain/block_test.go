package blockchain

import (
	"testing"
	"time"
)

func TestCalculateHash(t *testing.T) {
	dt := time.Now().String()
	b := Block{
		Index:        1,
		Timestamp:    dt,
		Transactions: []Transaction{},
		PrevHash:     "0",
		Nonce:        3000,
	}
	checkHash := b.CalculateHash()
	if checkHash != b.CalculateHash() {
		t.Errorf("Hash not consistent and would always be invalid: got %s... expected %s...", checkHash[:7], b.CalculateHash()[:7])
	}
	if len(checkHash) != 64 {
		t.Errorf("Hash is not the proper length: got %d and expected 64", len(checkHash))
	}
}

func TestCreateGenesisBlock(t *testing.T) {
	gb := CreateGenesisBlock()

	if gb.Index != 0 {
		t.Errorf("Index which was created should be 0 instead it was %d", gb.Index)
	}
	if gb.PrevHash != "0" {
		t.Errorf("Previous hash is set to something other than 0, ")
	}

	if len(gb.Hash) != 64 {
		t.Errorf("Hash is not the proper length: got %d and expected 64", len(gb.Hash))
	}
}
