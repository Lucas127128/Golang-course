package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println(("three"))
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend!")
	default:
		fmt.Println("It's weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() >= 12:
		fmt.Println("It's afternoon")
	default:
		fmt.Println("It's morning")
	}
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("I'm an int")
		case bool:
			fmt.Println("I'm an boolean")
		default:
			fmt.Printf("Don't know the type of %s\n", t)
		}
	}
	whatAmI(t.Hour())
}
