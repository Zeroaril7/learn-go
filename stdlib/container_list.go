package stdlib

import (
	"container/list"
	"fmt"
)

// Struktur data untuk elemen antrian
type AntrianElement struct {
	Data string
}

// Struktur data antrian
type Antrian struct {
	antrian *list.List
}

// Membuat antrian baru
func NewAntrian() *Antrian {
	return &Antrian{antrian: list.New()}
}

// Menambahkan elemen ke antrian
func (a *Antrian) Enqueue(element AntrianElement) {
	a.antrian.PushBack(element)
}

// Mengambil elemen pertama dari antrian
func (a *Antrian) Dequeue() AntrianElement {
	if a.antrian.Len() == 0 {
		return AntrianElement{} // Elemen kosong jika antrian kosong
	}

	front := a.antrian.Front()
	a.antrian.Remove(front)
	return front.Value.(AntrianElement)
}

// Mengecek apakah antrian kosong
func (a *Antrian) IsEmpty() bool {
	return a.antrian.Len() == 0
}

// Menampilkan isi antrian
func (a *Antrian) Print() {
	for element := a.antrian.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(AntrianElement).Data)
	}
}

func SistemAntrian() {
	// Membuat antrian baru
	antrian := NewAntrian()

	// Menambahkan elemen ke antrian
	antrian.Enqueue(AntrianElement{Data: "Orang 1"})
	antrian.Enqueue(AntrianElement{Data: "Orang 2"})
	antrian.Enqueue(AntrianElement{Data: "Orang 3"})

	// Menampilkan isi antrian
	fmt.Println("Isi antrian:")
	antrian.Print()

	// Mengambil elemen dari antrian
	orangPertama := antrian.Dequeue()
	orangKedua := antrian.Dequeue()

	fmt.Println("\nOrang pertama yang dilayani:", orangPertama.Data)
	fmt.Println("Orang kedua yang dilayani:", orangKedua.Data)

	// Menampilkan isi antrian setelah dequeue
	fmt.Println("\nIsi antrian setelah dequeue:")
	antrian.Print()
}
