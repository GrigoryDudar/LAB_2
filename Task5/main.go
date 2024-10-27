package main

import (
    "fmt"
)

func factorial(number int) int {
    result := 1
    for i := 1; i <= number; i++ {
        result *= i
    }
    return result
}

func main() {
    var num int
    fmt.Print("Enter a number: ")
    fmt.Scan(&num)
    fmt.Printf("The factorial of %d is %d\n", num, factorial(num))
}
