package fundamental

import "fmt"

type Student interface {
	GetID() string
	GetName() string
}

type Person struct {
	ID   string
	Name string
}

func (person Person) GetID() string {
	return person.ID
}

func (person Person) GetName() string {
	return person.Name
}

func IdCard(student Student) {
	fmt.Println("Student ID:", student.GetID())
	fmt.Println("Student Name:", student.GetName())
}

// Interface berisi kumpulan Method

func Interface() {

	println("In Interface File")

	msg := Msg
	student := Person{ID: "1", Name: "Budi"}
	if student.ID != "" && student.Name != "" {
		fmt.Println("Student Found:", msg(true))
		fmt.Println("Student", msg("Appeared"))
	}
	IdCard(student)
	println()
}

// Interface kosong merupakan type data yg dapat menampung segala type data
func Msg(msg interface{}) interface{} {
	return msg
}
