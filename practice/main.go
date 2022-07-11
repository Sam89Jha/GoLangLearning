package main

import (
	"fmt"
)

func main() {
	doSomething()

	sum := addValues(5, 8)
	fmt.Println("Sum is", sum)

	multiSum := AddAllValues(1, 2, 3, 4, 5, 6)
	fmt.Println("Sum is", multiSum)
}

func doSomething() {
	fmt.Println("Doing Siomething")
}

func addValues(value1, value2 int) int {
	return value1 + value2
}

func AddAllValues(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}
