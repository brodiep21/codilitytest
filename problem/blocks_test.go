package blocks_test

import (
	"testing"

	"github.com/joshuaswickirl/notblocks/blocks"
)

func TestBocks(t *testing.T) {
	s := "abbabbaaa"
	minNumLetters := blocks.FindMinNumberOfAdditionalLettersForBlocksToBeEqualLengths(s)
	if minNumLetters != 6 {
		t.Errorf("got %d, wanted 6", minNumLetters)
	}
}
