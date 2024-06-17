package main

import (
	"fmt"
	"strings"
)

func main() {
	//  kosong , seharusnya anda sudah paham mengenai materi ini karena sangat sederhana :)
	var name string
	fmt.Print("Enter true or false: ")
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)

	// Convert the input to a boolean
	var booleanValue bool
	if name == "true" {
		booleanValue = true
	} else if name == "false" {
		booleanValue = false
	} else {
		fmt.Println("Invalid input. Please enter 'true' or 'false'.")
		return
	}

	// Make a choice based on the boolean value
	if booleanValue {
		fmt.Println("You entered true.")
	} else {
		fmt.Println("You entered false.")
	}
}
