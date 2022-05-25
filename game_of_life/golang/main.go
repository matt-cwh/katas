package main

import "fmt"

func main() {

	var current string

	inputArray := [][]bool{
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, true, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, true, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, true, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, true, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, true, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
		{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false},
	}

	s := printArray(inputArray)
	fmt.Println(s)
	grid := newGrid(inputArray)

	for current != "c" {
		fmt.Scanln(&current)
		inputArray = grid.nexGen()
		s = printArray(inputArray)
		fmt.Println(s)
	}

	fmt.Println("end")
}
