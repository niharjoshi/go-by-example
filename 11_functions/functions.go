package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func plus(a, b, c int) int {
	return a + b + c
}

func main() {

	res := add(1, 2)
	fmt.Println("add:", res)

	res = plus(1, 2, 3)
	fmt.Println("plus:", res)

}
