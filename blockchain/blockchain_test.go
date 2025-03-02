package blockchain

import (
	"math/rand"
	"testing"
)

func GenerateDifficultyTest(t *testing.T) {
	rn := rand.Intn(99999)
	diff := GenerateDifficulty(rn)

	if len(diff) != rn {
		t.Errorf("length of difficulty does not match. Expected %d, got %d", rn, len(diff))
	}
}
