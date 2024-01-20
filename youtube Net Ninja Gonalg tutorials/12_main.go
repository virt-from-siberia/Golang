package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":   99.4,
		"pie":    6.4,
		"salad":  11,
		"coffie": 2.2,
	}

	fmt.Println(menu["pie"])

	// loop maps
	for k, v := range menu {
		fmt.Println(k, "- ", v)
	}

	//  int as key type
	phonebook := map[int]string{
		1312313: "Aleks",
		5675675: "Ivan",
		678686:  "Tanya",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[1312313])

	phonebook[1312313] = "Bob"
	fmt.Println(phonebook[1312313])

	phonebook[1312313] = "Bob"
	fmt.Println(phonebook[1312313])
}
