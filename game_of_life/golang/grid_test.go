package main

import "testing"

func getSampleSlice() [][]bool {
	return [][]bool{
		{true, false, false, true, true, true, false, false, true, true},
		{true, false, false, true, true, true, false, false, true, true},
		{true, false, false, true, true, true, false, false, true, true},
		{true, false, false, true, true, true, false, false, true, true},
		{true, false, false, true, true, true, false, false, true, true},
		{true, true, false, true, true, true, false, false, true, true},
		{true, false, false, true, true, true, false, false, true, true},
		{true, false, false, true, true, true, false, false, true, true},
		{true, false, false, true, true, true, false, false, true, true},
		{true, false, false, true, true, true, false, false, true, true},
	}
}

func assertCorrectNeighbour(t testing.TB, got *coordinate, want *coordinate) {
	t.Helper()

	if want == nil && got != nil {
		t.Errorf("expect nil but got %v %v", got.x, got.y)
		return
	}

	if want != nil && got == nil {
		t.Errorf("expect %v %v but got nil", want.x, want.y)
		return
	}

	if got != nil && want != nil && got.x != want.x && got.y != want.y {
		t.Errorf("expect %v %v but got %v %v", want.x, want.y, got.x, got.y)
	}
}

func TestGetCellLeftNeighbours(t *testing.T) {
	grid := newGrid(getSampleSlice())
	leftNeighbour := grid.getNeighbour(Left, newCoordinate(0, 0))
	assertCorrectNeighbour(t, leftNeighbour, nil)

	leftNeighbour = grid.getNeighbour(Left, newCoordinate(0, 3))
	assertCorrectNeighbour(t, leftNeighbour, nil)

	leftNeighbour = grid.getNeighbour(Left, newCoordinate(1, 0))
	assertCorrectNeighbour(t, leftNeighbour, newCoordinate(0, 0))
}

func TestGetCellTopLeftNeighbours(t *testing.T) {
	grid := newGrid(getSampleSlice())
	topLeftNeighbours := grid.getNeighbour(TopLeft, newCoordinate(0, 0))
	assertCorrectNeighbour(t, topLeftNeighbours, nil)

	topLeftNeighbours = grid.getNeighbour(TopLeft, newCoordinate(1, 9))
	assertCorrectNeighbour(t, topLeftNeighbours, nil)

	topLeftNeighbours = grid.getNeighbour(TopLeft, newCoordinate(7, 3))
	assertCorrectNeighbour(t, topLeftNeighbours, newCoordinate(6, 4))

	topLeftNeighbours = grid.getNeighbour(TopLeft, newCoordinate(9, 3))
	assertCorrectNeighbour(t, topLeftNeighbours, newCoordinate(8, 4))
}

func TestGetCellTopNeighbours(t *testing.T) {
	grid := newGrid(getSampleSlice())
	topNeighbour := grid.getNeighbour(Top, newCoordinate(0, 0))
	assertCorrectNeighbour(t, topNeighbour, newCoordinate(0, 1))

	topNeighbour = grid.getNeighbour(Top, newCoordinate(0, 9))
	assertCorrectNeighbour(t, topNeighbour, nil)

	topNeighbour = grid.getNeighbour(Top, newCoordinate(7, 3))
	assertCorrectNeighbour(t, topNeighbour, newCoordinate(7, 4))
}

func TestGetCellTopRightNeighbours(t *testing.T) {
	grid := newGrid(getSampleSlice())
	topRightNeighbour := grid.getNeighbour(TopRight, newCoordinate(0, 0))
	assertCorrectNeighbour(t, topRightNeighbour, newCoordinate(1, 1))

	topRightNeighbour = grid.getNeighbour(TopRight, newCoordinate(1, 3))
	assertCorrectNeighbour(t, topRightNeighbour, newCoordinate(2, 4))

	topRightNeighbour = grid.getNeighbour(TopRight, newCoordinate(9, 0))
	assertCorrectNeighbour(t, topRightNeighbour, nil)

	topRightNeighbour = grid.getNeighbour(TopRight, newCoordinate(0, 9))
	assertCorrectNeighbour(t, topRightNeighbour, nil)
}

func TestGetCellRightNeighbours(t *testing.T) {
	grid := newGrid(getSampleSlice())
	rightNeighbour := grid.getNeighbour(Right, newCoordinate(0, 0))
	assertCorrectNeighbour(t, rightNeighbour, newCoordinate(1, 0))

	rightNeighbour = grid.getNeighbour(Right, newCoordinate(1, 0))
	assertCorrectNeighbour(t, rightNeighbour, newCoordinate(2, 0))

	rightNeighbour = grid.getNeighbour(Right, newCoordinate(9, 0))
	assertCorrectNeighbour(t, rightNeighbour, nil)
}

func TestGetCellBottomRightNeighbours(t *testing.T) {
	grid := newGrid(getSampleSlice())
	bottomRightNeighbour := grid.getNeighbour(BottomRight, newCoordinate(0, 0))
	assertCorrectNeighbour(t, bottomRightNeighbour, nil)

	bottomRightNeighbour = grid.getNeighbour(BottomRight, newCoordinate(1, 3))
	assertCorrectNeighbour(t, bottomRightNeighbour, newCoordinate(2, 2))

	bottomRightNeighbour = grid.getNeighbour(BottomRight, newCoordinate(9, 1))
	assertCorrectNeighbour(t, bottomRightNeighbour, nil)

	bottomRightNeighbour = grid.getNeighbour(BottomRight, newCoordinate(9, 0))
	assertCorrectNeighbour(t, bottomRightNeighbour, nil)

	bottomRightNeighbour = grid.getNeighbour(BottomRight, newCoordinate(8, 9))
	assertCorrectNeighbour(t, bottomRightNeighbour, newCoordinate(9, 7))
}

func TestGetCellBottomNeighbours(t *testing.T) {
	grid := newGrid(getSampleSlice())
	bottomNeighbour := grid.getNeighbour(Bottom, newCoordinate(0, 0))
	assertCorrectNeighbour(t, bottomNeighbour, nil)

	bottomNeighbour = grid.getNeighbour(Bottom, newCoordinate(2, 0))
	assertCorrectNeighbour(t, bottomNeighbour, nil)

	bottomNeighbour = grid.getNeighbour(Bottom, newCoordinate(0, 1))
	assertCorrectNeighbour(t, bottomNeighbour, newCoordinate(0, 0))

	bottomNeighbour = grid.getNeighbour(Bottom, newCoordinate(3, 5))
	assertCorrectNeighbour(t, bottomNeighbour, newCoordinate(3, 4))
}

func TestGetCellBottomLeftNeighbours(t *testing.T) {
	grid := newGrid(getSampleSlice())
	bottomLeftNeighbour := grid.getNeighbour(BottomLeft, newCoordinate(0, 0))
	assertCorrectNeighbour(t, bottomLeftNeighbour, nil)

	bottomLeftNeighbour = grid.getNeighbour(BottomLeft, newCoordinate(0, 1))
	assertCorrectNeighbour(t, bottomLeftNeighbour, nil)

	bottomLeftNeighbour = grid.getNeighbour(BottomLeft, newCoordinate(1, 1))
	assertCorrectNeighbour(t, bottomLeftNeighbour, newCoordinate(0, 0))

	bottomLeftNeighbour = grid.getNeighbour(BottomLeft, newCoordinate(9, 7))
	assertCorrectNeighbour(t, bottomLeftNeighbour, newCoordinate(8, 6))
}

func TestGetCellAllNeighbours(t *testing.T) {
	grid := newGrid(getSampleSlice())
	neighbours := grid.getAllNeighbours(newCoordinate(0, 0))
	if neighbours == nil || len(neighbours) != 3 {
		t.Errorf("expected 3 neighbours but got %d", len(neighbours))
	}

	neighbours = grid.getAllNeighbours(newCoordinate(9, 9))
	if neighbours == nil || len(neighbours) != 3 {
		t.Errorf("expected 3 neighbours but got %d", len(neighbours))
	}

	neighbours = grid.getAllNeighbours(newCoordinate(5, 5))
	if neighbours == nil || len(neighbours) != 8 {
		t.Errorf("expected 8 neighbours but got %d", len(neighbours))
	}

	neighbours = grid.getAllNeighbours(newCoordinate(0, 5))
	if neighbours == nil || len(neighbours) != 5 {
		t.Errorf("expected 5 neighbours but got %d", len(neighbours))
	}

	neighbours = grid.getAllNeighbours(newCoordinate(9, 0))
	if neighbours == nil || len(neighbours) != 3 {
		t.Errorf("expected 3 neighbours but got %d", len(neighbours))
	}
}

func TestGetCellAllLiveNeighbours(t *testing.T) {
	grid := newGrid(getSampleSlice())

	noOfLiveNeighbors := grid.getAllLiveNeighbours(newCoordinate(0, 0))
	if noOfLiveNeighbors != 1 {
		t.Errorf("expected 1 neighbours but got %d", noOfLiveNeighbors)
	}

	noOfLiveNeighbors = grid.getAllLiveNeighbours(newCoordinate(4, 1))
	if noOfLiveNeighbors != 4 {
		t.Errorf("expected 4 neighbours but got %d", noOfLiveNeighbors)
	}

	noOfLiveNeighbors = grid.getAllLiveNeighbours(newCoordinate(4, 4))
	if noOfLiveNeighbors != 8 {
		t.Errorf("expected 8 neighbours but got %d", noOfLiveNeighbors)
	}
}
