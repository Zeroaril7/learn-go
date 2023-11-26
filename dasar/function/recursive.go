package function

func Recursive() {
	println("In Recursive File")
	println(factorial(4))
	println()
}

func factorial(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorial(value-1)
	}
}