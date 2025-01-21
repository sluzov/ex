package main

import (
	"fmt"
	// Импортируйте ваш пакет по имени из go.mod
	"log"
	"sluzov/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("C:/1Sluzov/Test1/data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0

	for _, number := range numbers {
		sum += number
	}

	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
