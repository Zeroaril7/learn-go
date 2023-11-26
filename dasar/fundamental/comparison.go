package fundamental

/*
	Operator Perbandingan
	>	(lebih dari)
	<	(kurang dari)
	>= 	(lebih dari sama dengan)
	<=	(kurang dari sama dengan)
	==	(sama dengan)
	!=	(tidak sama dengan)

	Operator Boolean
	&&	And
	||	Or
	!	Not
*/

func Comparison(){
	x := 2
	y := 3

	println("In Comparison File")

	println(x > y)

	if x > 0 && x < y {
		println(x)
	}

	println()
}