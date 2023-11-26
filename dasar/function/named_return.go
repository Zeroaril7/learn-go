package function

import "fmt"

func NamedReturn() {
	println("In Named Return File")

	adress, country := getAdress("Jln Yanto Bagus", "Kawala", "Kin", "0000")

	println(adress, country)

	println()
}

func getAdress(street, district, subDistrict, zipcode string) (adress, country string) {

	adress = fmt.Sprintf("%s, %s, %s, %s,", street, subDistrict, district, zipcode)
	country = "Konoha"

	return
}