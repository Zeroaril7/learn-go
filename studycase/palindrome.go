package studycase

import "fmt"

func Palindrome(text string) {
	reverseText := ""
	for _, v := range text {
		reverseText = string(v) + reverseText
	}

	fmt.Println(reverseText == text)
}
