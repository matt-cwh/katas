package main

import (
	"testing"
)

func TestPrintArray(t *testing.T) {
	inputArray := [][]bool{
		{true, false, false},
		{true, true, false},
		{true, true, false},
		{true, false, false},
	}
	s := printArray(inputArray)
	expect := "*     \n* *   \n* *   \n*     "
	if s != expect {
		t.Errorf("expect %v but got %v", expect, s)
	}
}
