package studycase

import "fmt"

func Fibo(length int) {
	firstFibo := 0
	secondFibo := 1

	for i := 1; i < length; i++ {
		firstFibo, secondFibo = secondFibo, firstFibo+secondFibo
	}

	fmt.Println(secondFibo)
}
