package countVowels

import (
	"fmt"
	"strings"
)

var vowels = []string{"a", "e", "i", "o", "u"}

func CountVowels(sentence string) {
	total := 0
	for _, v := range vowels {
		total += strings.Count(strings.ToLower(sentence), v)
	}

	fmt.Println("The word", sentence, "contains", total, "vowels")
}
