package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("wrong number")
	}
	return arg + 3, nil
}

var ErrOutOfTea = fmt.Errorf("out of tea")
var ErrNoPower = fmt.Errorf("no power")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("%w", ErrNoPower)
	}
	return nil
}

func main() {

	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println("f failed", e)
		} else {
			fmt.Println("f worked", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("time to buy tea")
			} else if errors.Is(err, ErrNoPower) {
				fmt.Println("wait for power")
			} else {
				fmt.Printf("unknown error: %s", err)
			}
			continue
		}
		fmt.Println("tea is ready")
	}

}
