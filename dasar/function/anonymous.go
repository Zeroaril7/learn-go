package function

func Anonymous() {

	id := 1

	getId := func(id int) string {
		switch id {
		case 1:
			return "Makanan"
		case 2:
			return "Minuman"
		default:
			return "Not Found"
		}
	}

	println("In Anonymous File")

	println(getId(id))
	id = 2
	println(getId(id))
	id = 3
	println(getId(id))

	println()
}