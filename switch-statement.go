package main

import (
	"fmt"
	"time"
)

func main() {
	dateFormat := "2006-01-02"

	christmas := time.Date(time.Now().Year(), time.December, 25, 0, 0, 0, 0, time.UTC)
	today := time.Now().Format(dateFormat)

	zeroedDate, _ := time.Parse(dateFormat, today)

	switch {
	case zeroedDate.Equal(christmas):
		fmt.Println("Merry Christmas!!")
	case zeroedDate.After(christmas):
		fmt.Println("Christmas has passed!")
	default:
		fmt.Println("Christmas is coming!")
	}
}
