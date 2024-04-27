package ngide

import (
	"fmt"
	"strings"
)

func ExtractKey() {

	data := "\"purchase_order_number\":\"PO-2050-12-01-0001\",\"purchase_order_date\":\"2021-01-12\",\"invoice_date\":\"2021-10-29\",\"status\":\"JatuhTempo\",\"due_flag\":\"OVERDUE\",\"due_date\":\"2023-10-01\",\"days_to_due_date\":77,\"project_name\":\"BelanjaLangsungPO-2021-01-12-8383\",\"project_category\":\"MaterialKonstruksi\",\"total_project_value\":1045000.00,\"type_project_value\":\"CAPEX\",\"uid\":\"631a53e47255a77e0e6f923d\",\"umkm_name\":\"TEST\",\"bumnid\":\"TELKOM\",\"buyer_group\":\"TEST\",\"buyer_group_province\":\"DKIJAKARTA\",\"buyer_group_city\":\"JakartaSelatan\",\"cluster\":\"JasaKeuangan\",\"batch\":\"Batch1\",\"modified_date\":\"2023-12-2013\""

	keys := strings.FieldsFunc(data, func(r rune) bool { return strings.ContainsRune(" :,", r) })

	i := 1
	for i < len(keys) {
		keys = removeSlice(keys, i)
		i++
	}

	// Remove the element at index i from a.
	var result string
	for _, v := range keys {
		result += v + ", "
	}

	fmt.Println(result)
}

func removeSlice(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}
