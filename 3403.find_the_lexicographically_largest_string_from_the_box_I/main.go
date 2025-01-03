package main

// time complexity: O(n)
// because max() and min() have only two fixed parameters,
// the time complexity is O(1)
//
// space complexity: O(1)
//
// "the world is split into numFriends non-empty strings."
// from this, we can extract the following informations:
//  1. if numFriends == 1, simply return the word, as we are required
//     to split it into non-empty strings.
//  2. for any word, the maximum substring size that we can obtain is
//     (word.size() - (numFriends - 1))
func answerString(word string, numFriends int) string {
	if numFriends == 1 {
		return word
	}

	ansLen := len(word) - numFriends + 1
	answer := ""
	for i := 0; i < len(word); i++ {
		answer = max(answer, word[i:min(i+ansLen, len(word))])
	}
	return answer
}
