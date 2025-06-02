package main

import "fmt"

func main() {
	fmt.Println("Hello Pointers...")

	// var ptr *int
	// fmt.Println("Value of ptr :", ptr)

	mynumber := 23
	var ptrMyNumber = &mynumber
	fmt.Println("Reference of ptrMyNUmber :", ptrMyNumber)
	fmt.Println("value of ptrMyNUmber :", *ptrMyNumber)
	fmt.Println("myNumber :", mynumber)

	*ptrMyNumber = *ptrMyNumber * 2
	fmt.Println("mynumber :", mynumber)
}
