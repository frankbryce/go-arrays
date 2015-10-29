package constant

import "fmt"

//package stringutil

func Constant1() {
	const (
		c0 = 42 * iota
		c1
		c2
	)

	fmt.Println(c0)
	fmt.Println(c1)
	fmt.Println(c2)
}
