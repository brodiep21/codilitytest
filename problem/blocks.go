package blocks

func FindMinNumberOfAdditionalLettersForBlocksToBeEqualLengths(s string) int {
	// Loop over and find the length of the longest block
	currentBlock := 0
	lengthOfLongestBlock := len(findBlock(s, 0))
	for {
		block := findBlock(s, currentBlock)
		if len(block) > lengthOfLongestBlock {
			lengthOfLongestBlock = len(block)
		}
		currentBlock++
	}

	// Find how many letters needs to be added to other blocks

	// Add those numbers together
	return 6
}

func findBlock(s string, block int) (string, bool) {

	currentBlock := 0
	startingIndex := 0
	endingIndex := 0

	for {
		endOfBlock, endOfString := findEndOfBlock(s, startingIndex)
		if block == currentBlock {
			endingIndex = endOfBlock
			break
		}
		startingIndex = endOfBlock
		if endOfString {

		}
		currentBlock++
	}
	return s[startingIndex:endingIndex]
}

const endOfString = true

func findEndOfBlock(s string, startingIndex int) (int, bool) {
	endingIndex := startingIndex + 1
	for {
		if endingIndex >= len(s) {
			return endingIndex, endOfString
		}
		if s[startingIndex] != s[endingIndex] {
			break
		}
		endingIndex++
	}
	return endingIndex
}
