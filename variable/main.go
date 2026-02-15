package main

import (
	"fmt"
	"math"
)

func main() {
	// var name type = initialvalue

	var age int = 10 // this is a basic variable declaration with type and initialization

	fmt.Println("Age:", age)

	//	Type inference
	//
	// If a variable has an initial value,
	// Go will automatically be able to infer the type of that variable using that initial value.
	// Hence if a variable has an initial value, the type in the variable declaration can be removed.

	var name = "test"
	fmt.Println(name)

	// 	Short hand declaration
	// Go also provides another concise way to declare variables.
	//  This is known as short hand declaration and it uses := operator.

	a, b := 20, 30 //a and b declared
	fmt.Println("a is", a, "b is", b)
	// a, b := 40, 50 //error, no new variables

	c := math.Min(float64(a), float64(b))
	fmt.Println("Minimum of a and b is", c)
}
