package main

import "fmt"

func main()  {

	// Arrays (are fixed size)
	// var ages [3]int = [3]int{2, 27, 35}
	var ages = [3]int{2, 27, 35}
	names := [4]string{"Yoshi", "Mario", "Peach", "Bowser"}
	names[1] = "Luigi"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))
	
	// Slices (can be manipulated)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	// Slice ranges
	rangeOne := names[1:3] // include the first number but no the second
	rangeTwo := names[2:] // from the first number until the dend
	rangeThree := names[:3] // include the first number but no the second

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "Koopa")
	fmt.Println(rangeOne)
}