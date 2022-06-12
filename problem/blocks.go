package blocks

func FindMinNumberOfAdditionalLettersForBlocksToBeEqualLengths(s string) int {
	// Loop over and find the length of the longest block
	lengthOfLongestBlock, numberOfBlocks := findLengthOfLongestBlock(s)

	// Find how many letters needs to be added to other blocks
	totalNumberOfAdditionalLetters := 0
	for i := 0; i <= numberOfBlocks; i++ {
		startingIndex, endingIndex := findBlock(s, i)
		lengthOfBlock := endingIndex - startingIndex
		totalNumberOfAdditionalLetters += lengthOfLongestBlock - lengthOfBlock

	}

	return totalNumberOfAdditionalLetters
}

func findLengthOfLongestBlock(s string) (int, int) {
	currentBlock := 0
	lengthOfLongestBlock := 0
	for {
		startingIndex, endingIndex := findBlock(s, currentBlock)
		lengthOfBlock := endingIndex - startingIndex
		if lengthOfBlock > lengthOfLongestBlock {
			lengthOfLongestBlock = lengthOfBlock
		}
		if endingIndex >= len(s) {
			break
		}
		currentBlock++
	}
	return lengthOfLongestBlock, currentBlock
}

func findBlock(s string, block int) (int, int) {

	currentBlock := 0
	startingIndex := 0
	endingIndex := 0

	for {
		endOfBlock := findEndOfBlock(s, startingIndex)
		if block == currentBlock {
			endingIndex = endOfBlock
			break
		}
		startingIndex = endOfBlock
		currentBlock++
	}
	return startingIndex, endingIndex
}

func findEndOfBlock(s string, startingIndex int) int {
	endingIndex := startingIndex + 1
	for {
		if endingIndex >= len(s) {
			return endingIndex
		}
		if s[startingIndex] != s[endingIndex] {
			break
		}
		endingIndex++
	}
	return endingIndex
}
