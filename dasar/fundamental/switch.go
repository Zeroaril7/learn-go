package fundamental

func Switch() {
	typeOrder := "food"

	println("In Switch File")

	// Normal switch case
	switch typeOrder {
	case "food":
		println(typeOrder)
	case "beverage":
		println(typeOrder)
	default:
		println("Not Found Type Order")
	}

	// switch case with short statement
	switch isValid := false; isValid {
	case true:
		println(isValid)
	case false:
		println(isValid)
	}

	// switch case tanpa kondisi
	length := len(typeOrder)
	switch {
	case length > 5:
		println("panjang", typeOrder, "adalah", len(typeOrder))
	case length < 5:
		println("panjang", typeOrder, "adalah", len(typeOrder))
	default:
		println("panjang", typeOrder, "adalah", len(typeOrder))
	}

	println()
}