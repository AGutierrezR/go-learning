package main

import "fmt"

func main()  {
	age := 35
	name := "Andres"

	// Print (no new line)
	fmt.Print("Hello, ")
	fmt.Print("world \n")
	fmt.Print("new line \n")

	// Pringln (add a new line automatly)
	fmt.Println("Hello world!")
	fmt.Println("goodbye ninjas!")
	fmt.Println("my age is", age, "and my name is", name) // no need for concatenation

	// Printf (formatted string) %_ format specifier
	fmt.Printf("My age is %v and name is %v \n", age, name)
	fmt.Printf("My age is %q and name is %q \n", age, name) // %q is for strings, number will be '#'
	fmt.Printf("age is of type %T \n", age)
	fmt.Printf("your scored %f points! \n", 2.255)
	fmt.Printf("your scored %0.1f points! \n", 2.255)

	// Sprintf (save formatted strings)
	var str  = fmt.Sprintf("My age is %v and name is %v", age, name)
	println(str)
}