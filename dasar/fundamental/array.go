package fundamental

import "fmt"

// Array termasuk tipe data yang menampung kumpulan data yang sama dengan tipe data yang sama(string, int, dll)
// Array selalu dimulai dari nol
func Array(){
	println("In Array File")

	// Deklarasi Array
	var names [3]string
	
	umur := [3]int{
		30, 34, 36,
	}

	hobi := [...]string{
		"masak",
		"mancing",
		"basket",
	}
	
	// Pengisian nilai dalam array
	names[0] = "Mamat"
	names[1] = "Budi"
	names[2] = "Wawan"

	fmt.Println(names)
	fmt.Println(umur)
	fmt.Println(hobi)
	

	// function bawaan array
	fmt.Println(len(hobi))
	fmt.Println(names[0])
	names[0] = "Yanto"
	fmt.Println(names[0])


	println()
}