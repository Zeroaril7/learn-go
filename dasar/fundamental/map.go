package fundamental

import "fmt"

/*
	Map = tipe data yang mengumpulkan data atau value berdasarkan key yang mana key harus unik

	map[key]value
	• key = berisi tipe data bebas yang digunakan untuk mengisi, mengubah, menghapus, dan mengakses value
	• value = berisi tipe data bebas yang digunakan untuk menyimpan nilai

	Fungsi bawaan map
	• len(map)						= Mendapat panjang/jumlah data map
	• map[key]						= Mengakses value
	• map[key] = value				= Mengubah data/value map di key
	• make(map[TypeKey]TypeValue)	= deklarasi map baru
	• delete(map, key)				= menghapus data di map dg key
*/

func Map() {

	println("In Map File")

	// Deklarasi
	food := map[int]string{
		0: "Potato",
		1: "Burger",
	}

	drink := make(map[int]string)

	drink[0] = "Tea"
	drink[1] = "Coffe"
	drink[2] = "Orange Jus"

	fmt.Println(food)
	fmt.Println(drink)
	fmt.Println("In key drink[0]", drink[0])

	drink[0] = "Lemon Jus"

	delete(drink, 1)
	fmt.Println(drink)

	println()
}