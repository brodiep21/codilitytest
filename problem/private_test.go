package blocks

import "testing"

func TestFindBlocks(t *testing.T) {
	tests := []struct {
		input         string
		block         int
		startingIndex int
		endingIndex   int
	}{
		{"ab", 0, 0, 1},
		{"aab", 0, 0, 2},
		{"aaaaaab", 0, 0, 6},
		{"ab", 1, 1, 2},
		{"aab", 1, 2, 3},
		{"aabb", 1, 2, 4},
	}

	for _, tc := range tests {
		startingIndex, endingIndex := findBlock(tc.input, tc.block)
		if startingIndex != tc.startingIndex && endingIndex != tc.endingIndex {
			t.Errorf("got %d, %d, want %d, %d", startingIndex, endingIndex, tc.startingIndex, tc.endingIndex)
		}
	}
}
