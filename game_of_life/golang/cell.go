package main

func getNextState(liveNeighbours int) bool {

	if liveNeighbours > 3 {
		return false
	}
	if liveNeighbours > 1 {
		return true
	}

	return false
}
