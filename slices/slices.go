package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println(s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println(s, len(s) == 3, cap(s) == 3)

	s[0] = "hello"
	s[1] = "hi"
	s[2] = "good morning"
	fmt.Println(s)
	fmt.Println(s[2])
	fmt.Println(len(s))

	s = append(s, "good afternoon")
	s = append(s, "good bye")
	fmt.Println(s)

	t := make([]string, len(s))
	copy(t, s)
	fmt.Println(t)

	l := s[2:5]
	fmt.Println(l)

	c := []string{"hello", "hello", "hello"}
	fmt.Println(c)

	c2 := []string{"hello", "hello", "hello"}
	if slices.Equal(c, c2) {
		fmt.Println("c==c2")
	}

}
