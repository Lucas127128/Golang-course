package main

import (
	"fmt"
)

func main() {
	var a [5]int
	fmt.Println(a)

	a[4] = 67
	fmt.Println(a[4])
	fmt.Println("length:", len(a))

	b := [...]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	b = [...]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	b = [...]int{1, 3: 400, 4}
	fmt.Println(b)
	twoD := [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	fmt.Printf("twoD: %v\n", twoD)
}
