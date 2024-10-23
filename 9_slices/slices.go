package main

import (
	"fmt"
	"slices"
)

func main() {

	var s []string
	fmt.Println("uninitialized:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("empty:", s, "length:", len(s), "capacity:", cap(s))

	s[0], s[1], s[2] = "a", "b", "c"
	fmt.Println("set:", s, "get:", s[0])

	s = append(s, "d", "e", "f")
	fmt.Println("appended:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copied:", c)

	l := s[1:2]
	fmt.Println("slice:", l)

	l = s[2:]
	fmt.Println("slice:", l)

	t := []string{"x", "y", "z"}
	fmt.Println("declared:", t)

	u := []string{"x", "y", "z"}
	if slices.Equal(t, u) {
		fmt.Println("equal:", t, u)
	}

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

}
