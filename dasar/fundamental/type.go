package fundamental

func Type() {

	println("In Type File")

	type NISN string

	var siswa NISN

	siswa = "00299128"

	newSiswa := "0012233"
	convertToSiswa := NISN(newSiswa)

	println(siswa)

	println(convertToSiswa)

	println()
}