package main

import "fmt"

func printSlice(slice []string) {
	for i, value := range slice {
		fmt.Printf("%d. %s\n", i+1, value)
	}
}

func main() {
	names := []string{"putri", "irvan", "arif", "bagas", "shohib", "fascal", "hafid", "karima"}

	printSlice(names)
}
