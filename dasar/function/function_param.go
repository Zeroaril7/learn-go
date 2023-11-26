package function

type Filter func(string) string

func AsParameter() {
	sayIt := sayHello

	println("In Function Parameter File")

	println(sayIt("Bos", filter))

	println(sayIt("employe", filter))

	println()
}

func sayHello(name string, filter Filter) string {
	return "Hello " + filter(name)
}

func filter(name string) string {
	switch name {
	case "Bos":
		return "Bos"
	default:
		return "Bro"
	}
}