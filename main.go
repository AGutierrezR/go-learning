package main

import "fmt"

var someName = "Yuji"
// otherName := "Gojo" not working

func main()  {
	// Strings	
	// var nameOne string = "Mario"
	// var nameTwo = "Luigi"
	// var nameThree string

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameOne = "Peach"
	// nameThree = "Bowser"

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameFour := "Yoshi"

	// fmt.Println(nameOne, nameTwo, nameThree, nameFour)
	// fmt.Println(someName)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 127 // int8 from -128 to 127
	var numTwo int8 = -128 // int8 from -128 to 127
	var numThree uint = 25 // unit are positive number

	fmt.Println(numOne, numTwo, numThree)

	var scoreOne float32 = 25.98 // floatings number 
	var scoreTwo float64 = 888888999.7 // large floating nubmers

	fmt.Println(scoreOne, scoreTwo)

}