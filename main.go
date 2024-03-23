package main

import "fmt"

func main()  {
	// Initial use of a for
	// x := 0
	
	// for x < 5 {
	// 	fmt.Println("Value of x is:", x)
	// 	x++
	// }

	// More regular way to using it (like JS)
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Value of x is:", i)
	// }

	names := []string{"mario" , "luigi", "yoshi", "peach"}

	// Using the len function to get the lenght of the array
	// for i := 0; i < len(names); i++ {
	// fmt.Println(names[i])
	// }

	// Using the "range" that returns 2 values, index and value
	// for index, value := range names {
	// 	fmt.Printf("the positon at index %v is %v \n", index, value)
	// }

	// Ignoring one of the values using underscore (_)
	for _, value := range names {
		fmt.Printf("the value is %v \n", value)
		value = "new string" // this won't change the value, this is a local scoped variable
	}

	fmt.Println(names) // will print the original names slice
}