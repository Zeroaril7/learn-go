package gorountine

import "testing"

func SayYes() {
	println("Yes")
}

func TestDeclareGorountine(t *testing.T) {
	// Digunakan saat ingin memanggil fungsi lain
	go SayYes()

	// Anonymous Func dan langsung execute menggunakan () di akhir bracket '{}'
	go func() {
		println("No")
	}()
}
