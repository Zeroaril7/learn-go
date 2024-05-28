package studycase

import (
	"fmt"
)

func Anagram(firstText, secondText string) {
	if len(firstText) != len(secondText) {
		fmt.Println(false)
		return
	}

	charCount := make(map[rune]int)

	for _, char := range firstText {
		charCount[char]++
	}

	for _, char := range secondText {
		if charCount[char] == 0 {
			fmt.Println(false)
			return
		}
		charCount[char]--
	}

	fmt.Println(true)
}
