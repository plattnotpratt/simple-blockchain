package blockchain

import (
	"math/rand"
	"strings"
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

func TestGenerateDifficultyLength(t *testing.T) {
	rn := rand.Intn(99999)
	diff := GenerateDifficulty(rn)

	if len(diff) != rn {
		t.Errorf("length of difficulty does not match. Expected %d, got %d", rn, len(diff))
	}
}

func TestGenerateDifficultyValue(t *testing.T) {
	rn := rand.Intn(10)
	diff := strings.Split(GenerateDifficulty(rn), "")
	for i := 0; i < len(diff); i++ {
		if diff[i] != "0" {
			t.Errorf("The character of the difficulty is incorrect. At index %d expected '0' and got %s", i, diff[i])
			return
		}
	}
}
