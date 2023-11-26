package function

func Return() {
	println("In Return File")

	result := saySomething("Hallo")

	println(result)

	println()
}

func saySomething(msg string) string {

	message := "Me: " + msg

	return message
}