package blocks

import "testing"

func TestFindBlocks(t *testing.T) {
	tests := []struct {
		input  string
		block  int
		output string
	}{
		{"ab", 0, "a"},
		{"aab", 0, "aa"},
		{"aaaaaab", 0, "aaaaaa"},
		{"ab", 1, "b"},
		{"aab", 1, "b"},
		{"aabb", 1, "bb"},
	}

	for _, tc := range tests {
		block := findBlock(tc.input, tc.block)
		if block != tc.output {
			t.Errorf("got %q, want %q", block, tc.output)
		}
	}
}
