package main

import "fmt"

type rectangle struct {
	width, height int
}

func (r *rectangle) area() int {
	return r.width * r.height
}

func (r *rectangle) perimeter() int {
	return 2 * (r.width + r.height)
}

func main() {

	r := rectangle{width: 10, height: 5}
	fmt.Println("area:", r.area())
	fmt.Println("perimeter:", r.perimeter())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perimeter:", rp.perimeter())

}
