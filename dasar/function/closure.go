package function

// Kemampuan function berinteraksi dg data disekitarnya
func Closure() {

	println("In Closure File")
	// data
	counter := 0

	// function
	increament := func() {
		counter ++
		println(counter)
	}

	// function
	increament()
	increament()
	increament()

	println()
}