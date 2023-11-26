package fundamental

func For() {

	println("In For File")

	i := 0

	// For Loop
	for i < 5 {
		i++
	}
	println("i =", i)

	// For Loop with statement
	for count := 0; count < 5; count++ {
		count++
		i = count
	}
	println("count", i)

	// For Each
	names := []string{"Mamat", "Budi", "Wawan"}

	for i, v := range names {
		println("index", i)
		println("value", v)
	}

	println()
}