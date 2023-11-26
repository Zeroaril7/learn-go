package main

import (
	"study-golang/dasar/function"
	"study-golang/dasar/fundamental"
	"study-golang/dasar/pkg"
)

func main() {

	// Package
	pkg.GetDatabase()
	pkg.Err()
	pkg.CustomError()

	// fundamental
	fundamental.HelloWorld()
	fundamental.Number()
	fundamental.IsValid()
	fundamental.String()
	fundamental.DecklarationVariable()
	fundamental.Const()
	fundamental.ConvertString()
	fundamental.Type()
	fundamental.Unary()
	fundamental.Comparison()
	fundamental.Array()
	fundamental.Slice()
	fundamental.Map()
	fundamental.If()
	fundamental.Switch()
	fundamental.For()
	fundamental.BreakContinue()
	fundamental.Struct()
	fundamental.Interface()
	fundamental.TypeAssertion()
	fundamental.Pointer()

	// Function
	function.Fun()
	function.Param()
	function.Return()
	function.ReturnMultiple()
	function.NamedReturn()
	function.Variadic()
	function.AsValue()
	function.AsParameter()
	function.Anonymous()
	function.Recursive()
	function.Closure()
	function.Defer()
	function.PanicRecover()
}
