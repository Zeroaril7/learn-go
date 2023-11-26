package function

func Param() {
	println("In Param File")

	sum(2, 3)
	multiply(2, 3)

	println()
}

func sum(x, y int) {
	println("x+y =", x+y)
}

func multiply(x int, y int) {
	println("x*y =", x*y)
}