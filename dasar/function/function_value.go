package function

func AsValue() {

	say := sayHi

	println("In Function Value file")

	println(say("Mamat"))

	println()
}

func sayHi(name string) string {
	return "Hi " + name
}