package main

import "fmt"

type empolyee struct {
	firstName, lastName string
	salary              int
}

func main() {
	r := "☞"
	fmt.Printf("Print the lower casee base 16 by %%x : %x \n", r)
	fmt.Printf("Print the default format by %%v      : %v \n", r)
	fmt.Printf("Print type value by %%T              : %T \n", r)

	fmt.Println()

	var e empolyee
	fmt.Println("Print Employee struct :\n", e)
}
