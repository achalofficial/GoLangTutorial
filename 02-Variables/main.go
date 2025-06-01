package main

import "fmt"

// := symbole not Allowed outside function

// Constant declaration not allowed
const PI float32 = 3.14

// Keeping the function's name or variable's name first letter as capital allow it to be exported, if needed.

func main() {
	var username string = "Achal"
	fmt.Println("User Name: ", username)
	fmt.Printf("Type of %s is %T \n", username, username)

	var isLoggedIn bool = false
	fmt.Println("Logged in :", isLoggedIn)
	fmt.Printf("Type of isLoggedIn is %T \n", username)

	var smallVal uint8 = 255
	fmt.Println("Logged in :", smallVal)
	fmt.Printf("Type of smallVal %T \n", smallVal)

	var smallFloat float64 = 255.343432343242424
	fmt.Println("Logged in :", smallFloat)
	fmt.Printf("Type of smallFloat %T \n", smallFloat)

	// Default value
	var anotherVariable int
	fmt.Println("Logged in :", anotherVariable)
	fmt.Printf("Type of anotherVariable %T \n", anotherVariable)

	// implicit type
	var website = "www.google.com"
	fmt.Println("Website :", website)

	// website = 3 ::::: Not allowed since it took string as default type based on value

	// no var style
	numberOfUser := 300
	fmt.Println("number of users: ", numberOfUser)
}
