package function

func Variadic() {
	println("In Variadic File")

	result := sumAll(1, 2, 3, 4, 5)

	println(result)

	numbers := []int{1, 2, 3}
	resultSlice := sumAll(numbers...)

	println(resultSlice)

	println()
}

func sumAll(numbers ...int) (result int) {
	for _, v := range numbers {
		result += v
	}
	return
}