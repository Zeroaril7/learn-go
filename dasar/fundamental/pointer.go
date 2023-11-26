package fundamental

import "fmt"

/*

Semua variabel di Go-Lang memiliki default Pass by value

Pass by value = Mengirim ke variabel/parameter/function/method lain dg cara duplikasi value
Reference = = Mengirim ke variabel/parameter/function lain dg cara mereferensikan value

Misal variabel x dan y
Pass By value : x = 1 y = x, (Jumlah nilai 1 ada 2, yaitu di variabel x dan y)
Reference : x = 1 y = x, (Jumlah nilai 1 ada 1, yaitu di variabel x saja)

Method Asterik (*) = Asterik digunakan untuk mengakses value dari reference

operator new = Mendeklarasikan variable pointer dg data kosong

*/

type Email struct {
	Name, Provider string
}

func Pointer() {
	// Pass by value
	firstEmail := Email{Name: "test", Provider: "@gmail.com"}
	secondEmail := firstEmail

	// Reference
	thirdEmail := Email{Name: "third", Provider: "@gmail.com"}
	fourthEmail := &thirdEmail

	fourthEmail.Name = "fourth" // Mengubah Name pada thirdEmail dan fourthEmail

	secondEmail.Provider = "@yahoo.com" // Mengubah provider secondEmail

	println("In Pointer File")

	fmt.Println(firstEmail)
	fmt.Println(secondEmail)
	fmt.Println(thirdEmail)
	fmt.Println(*fourthEmail)

	// Operator new
	newEmail := new(Email)
	newEmail.Name = "Zera"
	newEmail.Provider = "@test.com"

	fmt.Println(newEmail)
	newMail := newEmail
	newMail.Name = "Ze"

	fmt.Println(newEmail)
	fmt.Println(newMail)

	// Function Pointer
	setName(newMail)

	fmt.Println(newMail)

	// Method Pointer
	newMail.GetEmail()

	fmt.Println(newMail)
	println()
}

// Function Pointer
func setName(email *Email) {
	email.Name = "New Name"
}

// Method Pointer
func (m *Email) GetEmail() Email {
	m.Name = "Tes"
	m.Provider = "@test.com"
	return *m
}
