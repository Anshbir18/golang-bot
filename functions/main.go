package main

import "fmt"

func circle_radius(r float32)int{
	return int(2 * 3.14159 * float64(r))
}

func main(){

	radius := float32(5.0)
	fmt.Println("Circumference of circle with radius", radius, "is", circle_radius(radius))
}