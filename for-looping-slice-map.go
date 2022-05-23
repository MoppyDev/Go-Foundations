package main

import (
	"fmt"
)

func main() {
	// looping through a slice to list the index and value of each item
	fmt.Println("Looping through slice data to list index and value")
	slice := []int{2, 4, 8, 16, 32}

	// if you don't want to use i replace it with _
	for i, v := range slice {
		fmt.Println(i, "->", v)
	}

	// looping through a map to list the key and values
	fmt.Println("\nLooping through a map to list key and values")
	hashValues := map[string]string{
		"Name":                 "MoppyDev",
		"Occupation":           "Cloud Engineer",
		"Programming Language": "Go",
	}

	for k, v := range hashValues {
		fmt.Println("\nKey:", k, "\nValue:", v)
	}
}
