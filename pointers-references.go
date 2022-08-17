package main

import "fmt"

func main() {
	stringVar := "Initial String"
	fmt.Println("stringVar variable is equal to", stringVar)
	// &Variable - & in front equals the reference value in memory of that variable
	fmt.Println("Reference location of stringVar variable is equal to", &stringVar)
	updateStringPointer(&stringVar)
	fmt.Println("After func call, stringVar variable is equal to", stringVar)
}

func updateStringPointer(s *string) {
	fmt.Println("s is equal to", s)
	updateVal := "Post String"
	// Sending the reference into this function you can then change the pointer of that reference to the new value
	*s = updateVal
}
