package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for k := range 3 {
		fmt.Println("range", k)
	}

	for {
		fmt.Println("loop")
		break
	}

	for l := range 6 {
		fmt.Println(l % 2)
	}

}
