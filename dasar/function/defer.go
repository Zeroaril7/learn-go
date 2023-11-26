package function

// Defer dipakai ketika ingin sebuah func dieksekusi setelah sebuah func lain selesai dieksekusi meskipun terdapat error
func Defer() {
	println("In Defer File")
	defer logging()
	runApp()
}

func logging() {
	println("Stop app")
	println()
}

func runApp() {
	println("App Start")
}