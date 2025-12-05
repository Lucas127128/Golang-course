package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)
	m["key1"] = 7
	m["key2"] = 13
	fmt.Println(m)

	v := m["key1"]
	v3 := m["key3"]
	fmt.Println(v3)
	fmt.Println(v)
	fmt.Println(len(m))

	delete(m, "key2")
	fmt.Println(len(m))

	clear(m)
	fmt.Println(m)

	n := map[string]int{"foo": 2, "bar": 2}
	fmt.Println(n)

	n2 := map[string]int{"foo": 2, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n==n2")
	}
}
