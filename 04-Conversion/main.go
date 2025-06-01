package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Enter the Rating of Pizza :")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error inputing:", err)
	}
	fmt.Println("Thanks for rating ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println("Error in conv:", err)
	}

	fmt.Println("numRating :", numRating)
	fmt.Printf("numRating type : %T", numRating)

}
