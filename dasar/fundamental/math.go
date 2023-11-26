package fundamental

const (
	addition       = "+"
	subtraction    = "-"
	multiplication = "*"
	division       = "/"
	enrichment     = "%"
)

// Ada juga Augmented Assignments
// a = a + 1 / a += 1

func Math() {
	x := 2
	y := 3

	println("In Math File")

	println("x", addition, "y =", x+y)

	x += y
	println(x)

	println()
}