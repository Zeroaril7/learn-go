package function

func ReturnMultiple() {
	println("In Return Multiple")

	v, l := getLength("Mamat")

	println(v)
	println(l)

	println()

}

func getLength(v string) (string, int) {
	return v, len(v)
}