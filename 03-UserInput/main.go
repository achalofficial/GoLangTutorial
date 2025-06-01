package main

import (
	"bufio"
	"fmt"
	"os"
)

// Learning about comma-ok or comma-error syntax

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating:")

	// comma-ok Or Err-Ok syntax
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of rating, %T", input)

}
