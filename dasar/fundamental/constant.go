package fundamental

// Constant Tidak bisa diubah setelah diinisisasi
const (
	levelAdmin = "Admin"
	levelStaf  = "Staf"
)

func Const() {
	println("In Constant File")

	// levelAdmin = "s" Will error
	println(levelAdmin)

	println(levelStaf)

	println()

}