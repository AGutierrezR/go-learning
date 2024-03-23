package main

import "fmt"

func main()  {
	// age := 45

	// fmt.Println(age <= 50)
	// fmt.Println(age >= 50)
	// fmt.Println(age == 45)
	// fmt.Println(age != 50)

	// if age < 30 {
	// 	fmt.Println("Age is less than 30")
	// } else if age < 40 {
	// 	fmt.Println("age is less than 40")
	// } else {
	// 	fmt.Println("age is not less than 45")
	// }

	names := []string{"mario", "luigi", "yoshi", "peach", "bowser"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("Continuing at pos", index)
			continue // allows to skip to the next item
		}

		if index > 2 {
			fmt.Println("Breaking at pos", index)
			break // break the loop, exit out of it
		}

		fmt.Printf("The value at pos %v is %v \n", index, value)
	}
}