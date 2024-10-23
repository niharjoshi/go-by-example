package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]int)

	m["a"] = 1
	m["b"] = 2
	fmt.Println("map:", m)

	u := m["a"]
	fmt.Println("get:", u)

	v := m["c"]
	fmt.Println("get:", v)

	fmt.Println("length:", len(m))

	delete(m, "b")
	fmt.Println("delete:", m)

	clear(m)
	fmt.Println("clear:", m)

	_, r := m["a"]
	fmt.Println("present:", r)

	t := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("declare:", t)

	w := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(t, w) {
		fmt.Println("equal:", t, w)
	}

}
