package pkg

import (
	"errors"
	"fmt"
)

type CustomErr struct {
	err error
	msg string
}

func CustomError() {

	println("In Custom Error File")
	custom := CustomErr{err: errors.New("404"), msg: "Not Found"}
	fmt.Println(custom.err.Error(), custom.msg)

	custom.msg = "Tak ditemukan"
	fmt.Println(custom.err.Error(), custom.msg)

	println()
}

func (e *CustomErr) Error() CustomErr {
	return CustomErr{err: e.err, msg: e.msg}
}
