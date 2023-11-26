package fundamental

import "fmt"

// Type Assertion digunakan untuk konversi antar tipe data, sering kali digunakan ketika bertemu dg data interface kosong/any
func TypeAssertion() {

	println("In Type Assertion File")

	var msg any

	// type data still any
	msg = "Hai"

	// type data convert into string
	msgString := msg.(interface{})

	fmt.Println(msgString)

	switch msg.(type) {
	case string:
		println("msg is string")
	case int:
		println("msg is int")
	default:
		println("Not string and int")
	}

	println()
}
