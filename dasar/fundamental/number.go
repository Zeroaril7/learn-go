package fundamental

var (
	// Jenis Int
	sInt  int8  // -128 - 127
	mInt  int16 // -32768 - 32767
	lInt  int32 // -2147483648 - 2147483647
	xlInt int64 // -9223372036854775808 - 9223372036854775807

	// Jenis uint (unsigned int)
	sUint  uint8  // 0 -255
	mUint  uint16 // 0 - 65535
	lUint  uint32 // 0 - 4294967295
	xlUint uint64 // 0 - 18446744073709551615

	// Jenis floating point
	// type complex digunakan saat ingin ada angka real dan imaginary partsnya
	lFloat     float32
	xlFloat    float64
	xlComplex  complex64  //	= float32 real dan imaginary partsnya
	xxlComplex complex128 //	= float64 real dan imaginary partsnya

	// Alias
	aByte byte // = uint8
	aRune rune // = int32
	aInt  int  // = Minimal int32
	aUint uint // = Minimal uint32
)

func Number() {

	println("In Number File")

	sInt = -8
	println(sInt)

	sUint = 8
	println(sUint)

	lFloat = 3.5
	println(lFloat)

	aByte = 1
	println(aByte)

	println()
}