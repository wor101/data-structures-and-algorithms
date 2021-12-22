package main

import "fmt"

func missingLetter(input string) string {
	var missing string
	var inputLetters = make(map[string]bool)
	var alphabet []string

	for _, rune := range input {
		inputLetters[string(rune)] = true
	}

	for i := 'a'; i <= 'z'; i++ {
		alphabet = append(alphabet, string(i))
	}

	for _, letter := range alphabet {
		if inputLetters[letter] {
			continue
		} else {
			missing = letter
			break
		}
	}
	return missing
}

func main() {
	fmt.Println(missingLetter("the quick brown box jumps over a lazy dog"))

}
