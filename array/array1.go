package array

import "fmt"

//package stringutil

func Array1() {
	var array1 [10]int

	for i := 0; i < len(array1); i += 1 {
		array1[i] = i * 2
	}

	fmt.Println(array1)
}
