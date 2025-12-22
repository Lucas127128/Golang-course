package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.height * r.width
}

func (r rect) perim() int {
	return 2*r.height + 2*r.width
}
func main() {
	r := rect{width: 10, height: 5}

	fmt.Println(r.area())
	fmt.Println(r.perim())

	rp := &r
	fmt.Println(rp.area())
	fmt.Println(rp.perim())
}
