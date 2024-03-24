package main

import (
	"fmt"
)

func main() {

	menu := map[string]float64{
		"soup": 4.99,
		"pie": 6.99,
		"salad": 6.99,
		"toffe pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["soup"])
	
	// Looping maps
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// Ints as keys type
	phonebook := map[int]string {
		267584967: "mario",
		984759373: "luigi",
		845775485: "peach",
	}
		
	// Geting a value by key
	fmt.Println(phonebook[267584967])

	// looping
	for k, v := range phonebook {
		fmt.Println(k, "-", v)
	}

	// Modifying a value by key
	phonebook[984759373] = "bowser"

	fmt.Println(phonebook)
	
}