package fundamental

import "fmt"

func BreakContinue() {

	println("In Break Continue File")

	var index int
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		index = i
	}

	println(index)

	genap := make([]int, 0, 5)
	ganjil := make([]int, 0, 5)

	number := 0
	for number < 10 {
		if number%2 == 0 {
			genap = append(genap, number)
			number++
			continue
		}
		ganjil = append(ganjil, number)
		number++
	}

	fmt.Println("Numbers =", number)
	fmt.Println("Genap =", genap)
	fmt.Println("Ganjil =", ganjil)

	println()
}