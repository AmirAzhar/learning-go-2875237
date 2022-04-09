package main

import (
	"fmt"
)

const aConst int = 64

func main() {
	// Implicit typing
	var aString string = "This is Go!"
	fmt.Println(aString)
	fmt.Printf("aString's type is %T\n", aString)
	fmt.Println()

	var anInt int = 42
	fmt.Printf("anInt's type is %T\n", anInt)
	fmt.Println()

	var defaultInt int
	fmt.Println(defaultInt)
	fmt.Println()

	// Explicit Typing
	var anotherString = "This is another string"
	fmt.Println(anotherString)
	fmt.Printf("anotherString's type is %T\n", anotherString)
	fmt.Println()

	var anotherInt = "53"
	fmt.Println(anotherInt)
	fmt.Printf("anotherInt's type is %T\n", anotherInt)
	fmt.Println()

	// Declaration only used inside functions
	myString := "This is also a string"
	fmt.Println(myString)
	fmt.Printf("myString's type is %T\n", myString)
	fmt.Println()

	fmt.Println(aConst)
	fmt.Printf("aConst's type is %T\n", aConst)

}
