package blocks_test

import (
	"testing"

	"github.com/joshuaswickirl/notblocks/blocks"
)

func TestBocks(t *testing.T) {
	t.Run("scenario 1", func(t *testing.T) {
		s := "abbabbaaa"
		minNumLetters := blocks.FindMinNumberOfAdditionalLettersForBlocksToBeEqualLengths(s)
		if minNumLetters != 6 {
			t.Errorf("got %d, wanted 6", minNumLetters)
		}
	})

	t.Run("scenario 2", func(t *testing.T) {
		s := "abbabbaaabbbaabbbaaaa"
		minNumLetters := blocks.FindMinNumberOfAdditionalLettersForBlocksToBeEqualLengths(s)
		if minNumLetters != 16 {
			t.Errorf("got %d, wanted 16", minNumLetters)
		}
	})

}
