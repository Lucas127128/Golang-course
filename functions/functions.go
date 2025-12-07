package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func val() (int, int) {
	return 6, 7
}

func sum(nums ...int) {
	fmt.Println(nums)
	total := 0
	for _, number := range nums {
		total += number
	}
	fmt.Println(total)
	fmt.Println(len(nums))
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return fact(n - 1)
}
func main() {
	//normal function
	fmt.Println(plus(2, 3))

	//multiple return values
	a, b := val()
	fmt.Println(a, b)
	c, _ := val()
	_, d := val()
	fmt.Println(c, d)

	//Variadic Functions(don't need to define number of params)
	sum(3, 4, 5, 6)
	nums := []int{1, 2, 3}
	sum(nums...)

	//closures(remember variable from outer scope)
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	nextInts := intSeq()
	fmt.Println(nextInts())

	//recursion
	fmt.Println(fact(12))
	var fib func(n int) int
	//must be explicitly declared before define the function
	fib = func(n int) int {
		if n > 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}
	fib(7)
}
