package fundamental

import (
	"fmt"
	"reflect"
	"strconv"
)

func ConvertString() {
	println("In Convert File")

	strData := "123"

	fmt.Printf(strData + " type of strData is %v\n", reflect.TypeOf(strData))

	newStrData, _ := strconv.Atoi(strData)

	fmt.Print(newStrData)

	fmt.Printf(" type of newStrData is %v\n", reflect.TypeOf(newStrData))

	println()
}