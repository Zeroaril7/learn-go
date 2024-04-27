package gorountine

import "testing"

func SayYes() {
	println("Yes")
}

func Test_DeclareGorountine(t *testing.T) {
	// Digunakan saat ingin memanggil fungsi lain
	go SayYes()

	// Anonymous Func dan langsung execute menggunakan () di akhir bracket '{}'
	go func() {
		println("No")
	}()
}
