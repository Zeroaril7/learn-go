package function

import "fmt"

func PanicRecover() {

	println("In Panic File")

	run(false)
	run(true)

	startApp(false)
	startApp(true)
}

func startApp(err bool) {
	defer endApp()

	println("Error is", err)
	if err {
		panic("Something error")
	}

	println("App starting")
}

func endApp() {
	println("App Stoping")
	println()
}

func run(err bool) {
	defer stop()

	println("Error is", err)
	if err {
		panic("Something error")
	}

	println("App starting")
}

func stop() {
	message := recover()
	fmt.Println("App Stoping", message)
	println()
}