package main

import "fmt"

func main() {
	x := 1
	for x < 2 {
		fmt.Print("The value of x is ")
		fmt.Println(x)
		x++
		fmt.Print("The value of x is ")
		fmt.Println(x)
	}

	//for i := 0; i <= 5; i++ {
	//	fmt.Println(i)
	//}

	for i := range 3 {
		fmt.Println(i)
	}

	//for {
	//	fmt.Println("loop forever")
	//}

	fmt.Println("")
	for a := range 7 {
		if a < 5 {
			continue
		}
		fmt.Println(a)
	}
}
