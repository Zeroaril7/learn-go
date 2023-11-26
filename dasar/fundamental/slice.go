package fundamental

import "fmt"

/*
	Type data dalam slice
	• Pointer = Penunjuk data pertama di array pada slice
	• Length = panjang dari slice yang tak boleh lebih dari kapasitas
	• Capacity = kapasitas slice

	Membuat slice dari array
	• array[low:high] = membuat slice dari array mulai dari index low sampai high
	• array[low:] =  membuat slice dari array mulai dari index low sampai akhir
	• array[:high] = membuat slice dari array mulai dari index 0 sampai high
	• array[:] =  membuat slice dari array mulai dari index 0 sampai akhir

	Fungsi dalam slice
	• len(slice)					= Cek panjang slice
	• cap(slice)					= Cek kapasitas slice
	• append(slice, data)			= Menambah data ke slice, jika cap penuh membuat array baru
	• make([]TypeData, len, cap)	= Membuat slice baru
	• copy(destination, source)		= Menyalin slice dari source ke destination
*/

func Slice() {
	bulan := [...]string{
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"Mei",
		"Jun",
		"Jul",
		"Agu",
		"Sep",
		"Okt",
		"Nov",
		"Des",
	}

	println("In Slice File")

	slices1 := bulan[4:6]
	fmt.Println(slices1)
	
	slices2 := bulan[:6]
	fmt.Println(slices2)
	
	slices3 := bulan[4:]
	fmt.Println(slices3)

	slices4 := bulan[:]
	fmt.Println(slices4)

	value := make([]int, 3, 5)
	
	n := 5
	fmt.Println(value)
	fmt.Println(len(value))
	fmt.Println(cap(value))
	for i:=1; i <= n; i++ {
		value = append(value, i)
	}
	fmt.Println(value)
	fmt.Println(len(value))
	fmt.Println(cap(value))

	fromSlice := bulan[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(toSlice)

	println()

	
}