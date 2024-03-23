package main

import (
	"fmt"
	"sort"
)

func main()  {

	// greeting := "Hello there friends!"

	// fmt.Println(strings.Contains(greeting, "Hello")) // string.Contains is case sensitive 
	// fmt.Println(strings.ReplaceAll(greeting, "friends", "ninjas"))
	// fmt.Println(strings.ToUpper(greeting))
	// fmt.Println(strings.Index(greeting, "ll"))
	// fmt.Println(strings.Split(greeting, " "))
	
	// The original value is unchanged
	// fmt.Println("Original value is unchanged", greeting)

	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}
	
	sort.Ints(ages) // this will change the original slice
	fmt.Println(ages)

	index := sort.SearchInts(ages, 90) // if the item doesn't exist will return one number above the length
	fmt.Println(index)

	names := []string{"yoshi", "mario", "peach", "bowser", "luigi"}

	sort.Strings(names)
	fmt.Println((names))

	fmt.Println(sort.SearchStrings(names, "bowser"))
}