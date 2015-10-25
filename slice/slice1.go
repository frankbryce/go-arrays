package slice

import "fmt"

var slice1 []int
var slice2 []int

func Slice1() {
	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice1 = array[2:7]
	slice2 = slice1[0:2]
	printSlices()

	array[3] = 23
	printSlices()

	slice2 = slice1[3:5]
	printSlices()

	slice1[4] = 46
	printSlices()
}

func printSlices() {
	fmt.Println(slice1)
	fmt.Println(slice2)
}
