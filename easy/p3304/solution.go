package p3304

/*
Alice and Bob are playing a game. Initially, Alice has a string word = "a".

You are given a positive integer k.

Now Bob will ask Alice to perform the following operation forever:

Generate a new string by changing each character in word to its next character in the
English alphabet, and append it to the original word.

For example, performing the operation on "c" generates "cd" and performing the operation on "zb" generates "zbac".

Return the value of the kth character in word, after enough operations have been done for
word to have at least k characters.

Note that the character 'z' can be changed to 'a' in the operation.

Example 1:
Input: k = 5
Output: "b"

Example 2:
Input: k = 10
Output: "c"

Constraints:
	- 1 <= k <= 500

*/

func kthCharacter(k int) byte {
	word := []byte{97}

	generate := func(word *[]byte) {
		if len(*word) >= k {
			return
		}

		var new []byte
		for _, char := range *word {
			newChar := char + 1

			if newChar > 122 {
				newChar = 97
			}

			new = append(new, newChar)
		}

		(*word) = append((*word), new...)
	}

	for len(word) < k {
		generate(&word)
	}

	return word[k-1]
}
