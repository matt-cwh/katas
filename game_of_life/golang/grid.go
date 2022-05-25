package main

type coordinate struct {
	x int
	y int
}

func newCoordinate(x, y int) *coordinate {
	c := coordinate{x: x, y: y}
	return &c
}

type grid struct {
	width       int
	height      int
	positionMap map[position]coordinate
	data        [][]bool
}

func newGrid(data [][]bool) *grid {
	m := make(map[position]coordinate)

	m[Top] = *newCoordinate(0, 1)
	m[TopRight] = *newCoordinate(1, 1)
	m[Right] = *newCoordinate(1, 0)
	m[BottomRight] = *newCoordinate(1, -1)
	m[Bottom] = *newCoordinate(0, -1)
	m[BottomLeft] = *newCoordinate(-1, -1)
	m[Left] = *newCoordinate(-1, 0)
	m[TopLeft] = *newCoordinate(-1, 1)

	g := grid{width: len(data), height: len(data[0]), positionMap: m, data: data}
	return &g
}

func (grid *grid) getNeighbour(p position, c *coordinate) *coordinate {
	result := newCoordinate(c.x+grid.positionMap[p].x, c.y+grid.positionMap[p].y)

	if result.x < 0 || result.x >= grid.width || result.y < 0 || result.y >= grid.height {
		return nil
	}

	return result
}

func (grid *grid) getAllNeighbours(c *coordinate) []*coordinate {
	var neighbours []*coordinate
	for p := position(0); p < PositionLimit; p++ {
		neighbour := grid.getNeighbour(p, c)
		if neighbour != nil {
			neighbours = append(neighbours, neighbour)
		}
	}

	return neighbours
}

func (grid *grid) getAllLiveNeighbours(c *coordinate) int {

	allNeighbours := grid.getAllNeighbours(c)
	cnt := 0

	for i := 0; i < len(allNeighbours); i++ {
		n := *allNeighbours[i]
		if grid.data[n.x][n.y] == true {
			cnt++
		}
	}

	return cnt
}

func (grid *grid) nexGen() [][]bool {
	nextGenData := [][]bool{}

	for i := 0; i < len(grid.data); i++ {
		nextLine := []bool{}
		for j := 0; j < len(grid.data[i]); j++ {
			noLiveNeighbours := grid.getAllLiveNeighbours(newCoordinate(i, j))
			nextLine = append(nextLine, getNextState(noLiveNeighbours))
		}
		nextGenData = append(nextGenData, nextLine)
	}

	grid.data = nextGenData

	return grid.data
}
