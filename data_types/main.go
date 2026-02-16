package main

import "fmt"

func main() {

	// ================================
	// Boolean
	// ================================
	a := true
	b := false
	fmt.Println("a:", a, "b:", b)

	// ================================
	// Integer (int)
	// ================================
	var age int = 25
	fmt.Printf("age = %d, type = %T\n", age, age)

	// ================================
	// Unsigned Integer (uint)
	// ================================
	var count uint = 10
	fmt.Printf("count = %d, type = %T\n", count, count)

	// ================================
	// Byte (alias of uint8)
	// ================================
	var b1 byte = 'A'
	fmt.Printf("byte value = %v, type = %T\n", b1, b1)

	// ================================
	// Rune (alias of int32)
	// ================================
	var r rune = '✓'
	fmt.Printf("rune value = %v, char = %c, type = %T\n", r, r, r)

	// ================================
	// String
	// ================================
	first := "Naveen"
	last := "Ramanathan"
	name := first + " " + last
	fmt.Println("My name is", name)
}



