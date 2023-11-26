package pkg

import (
	"errors"
	"fmt"
)

type result struct {
	data interface{}
	err  error
}

func Err() {
	println("In Errors File")

	result := divide(10, 0)

	if result.err != nil {
		fmt.Println("Result:", result.err)
	} else {
		fmt.Println("Result:", result.data)
	}

	println()
}

func divide(value, divider int) result {
	if divider == 0 {
		return result{err: errors.New("Can't divide by zero")}
	} else {
		return result{data: value / divider}
	}
}
