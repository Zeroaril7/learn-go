package studycase

import "fmt"

func FizzBuzz(iteration int) {

	for i := 1; i <= iteration; i++ {

		condition := 0
		output := ""
		if i%3 == 0 {
			condition += 1
			output += "Fizz"
		}
		if i%5 == 0 {
			condition += 2
			output += "Buzz"
		}
		if condition != 0 {
			fmt.Printf("Number %d: %s\n", i, output)
		}

	}
}
