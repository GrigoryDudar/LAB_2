package main

import (
	"fmt"
)

func main() {
	var number, limit int

	fmt.Print("Enter a number to generate its multiplication table: ")
	fmt.Scanln(&number)

	fmt.Print("Enter the limit for the table: ")
	fmt.Scanln(&limit)

	fmt.Printf("Multiplication Table for %d:\n", number)
	for i := 1; i <= limit; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
	}
}
