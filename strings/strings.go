package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "hello"
	fmt.Println(len(s))

	for i := range len(s) {
		fmt.Printf("%x", s[i])
	}
	fmt.Println("")

	fmt.Println(utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
		fmt.Println("")
	}

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineResult(runeValue)
	}
}
func examineResult(r rune) {
	if r == 'h' {
		fmt.Println("found h")
	} else if r == 'e' {
		fmt.Println("Found e")
	}
}
