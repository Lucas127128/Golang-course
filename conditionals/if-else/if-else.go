package main

import (
	"fmt"
)

func main() {
	if 9%4 == 0 {
		fmt.Println("Nine is divisible by 4")
	} else {
		fmt.Println("Nine is not divisible by 4")
	}
	if 8%2 == 0 || 5%2 == 0 {
		fmt.Println("Either eight or five is even")
	}
	if num := 11; num > 10 {
		fmt.Println(num, "is greater than 10")
	} else if num < 10 {
		fmt.Println(num, "is less than 10")
	} else {
		fmt.Println(num, "is equal to 10")
	}
}
