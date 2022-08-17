package main

import "fmt"

type Car struct {
	Make      string
	Model     string
	Year      int
	Color     string
	DoorCount int
}

func main() {
	// Using a type struct doesn't require all fields.
	car := Car{
		Make:      "Honda",
		Model:     "Accord",
		DoorCount: 4,
	}

	// You can still call fields you didn't fill out
	fmt.Println(car.Make, car.Model, car.DoorCount, car.Color)
}
