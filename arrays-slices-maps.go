package main

import (
	"fmt"
)

func main() {

	// [...] allows the number of items in the array to be determined by what is in the curly braces
	arrayVal := [...]int{0, 1, 2, 3}
	fmt.Println("Value:", arrayVal)
	fmt.Println("Sub-Value:", arrayVal[2])
	fmt.Printf("Type: %T\n", arrayVal)

	sliceLiteralVal := []string{"position 0", "position 1", "etc"}
	fmt.Println("\nValue:", sliceLiteralVal)
	fmt.Println("Sub-Value:", sliceLiteralVal[2])
	fmt.Printf("Type: %T\n", sliceLiteralVal)

	sliceMakeVal := make([]int, 3)
	fmt.Println("\nValue:", sliceMakeVal)
	fmt.Println("Sub-Value:", sliceMakeVal[2])
	fmt.Printf("Type: %T\n", sliceMakeVal)

	// utilizing an interface allows different type values in what feels like an array
	interfaceVal := []interface{}{1, 2, "apple", true}
	fmt.Println("\nValue:", interfaceVal)
	fmt.Println("Sub-Value:", interfaceVal[2])
	fmt.Printf("Type: %T\n", interfaceVal)

	mapVal := make(map[string]float32)
	mapVal["Student1"] = 90.98
	mapVal["Student2"] = 55.23
	fmt.Println("\nMap values:", mapVal)
	fmt.Println("Subvalue:", mapVal["Student2"])

}
