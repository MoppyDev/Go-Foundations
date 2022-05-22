package main

import (
	"fmt"
)

// there is only for loops in go
// go has no do-while equivalent

// creating empty "array" or slice
// for each loop adding formatted string with count to the array
func main() {
	sliceVar := []string{}
	for i := 0; i < 10; i++ {
		result := fmt.Sprintf("Count #%d\n", i)
		sliceVar = append(sliceVar, result)
	}
	fmt.Println(sliceVar)
}
