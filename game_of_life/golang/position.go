package main

type position int

const (
	Top position = iota
	TopRight
	Right
	BottomRight
	Bottom
	BottomLeft
	Left
	TopLeft
	PositionLimit //For looping only
)
