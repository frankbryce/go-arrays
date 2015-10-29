package main

import (
	"crypto"
	"fmt"

	"github.com/frankbryce/go-explore/array"
	"github.com/frankbryce/go-explore/constant"
	"github.com/frankbryce/go-explore/slice"
)

func main() {
	arrays()
	slices()
	constants()
	cryptos()
}

func arrays() {
	fmt.Printf("Array1: \n")
	array.Array1()
}

func slices() {
	fmt.Printf("\nSlice1: \n")
	slice.Slice1()
}

func constants() {
	fmt.Printf("\nConstant1: \n")
	constant.Constant1()

	fmt.Printf("Constant2: \n")
	constant.Constant2()
}

func cryptos() {
	fmt.Printf("\nCrypto1: \n")
	fmt.Println(crypto.MD5)
}
