package main

import "fmt"

func zeroVal(ival int) {
	ival = 0
}

func zeroPtr(iptr *int) {
	*iptr = 0
}

func main() {

	i := 1

	fmt.Println("initial:", i)

	zeroVal(i)
	fmt.Println("zeroval:", i)

	zeroPtr(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("address:", &i)

}
