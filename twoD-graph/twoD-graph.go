package twodgraph

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

type TwoDGraph struct {
	Hash    map[Point]map[Point]struct{}
	visited map[Point]bool
}

func NewTwoDGraph() *TwoDGraph {
	return &TwoDGraph{
		Hash:    make(map[Point]map[Point]struct{}),
		visited: make(map[Point]bool),
	}
}

func (g *TwoDGraph) FormGraphFromMatrix(matrix [][]int) {
	// reset old entries!
	g.Hash = make(map[Point]map[Point]struct{})

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				continue
			}

			p1 := NewPoint(i, j)
			g.AddPointIfNotExists(*p1)

			if i+1 < len(matrix) && matrix[i][j] > 0 && matrix[i+1][j] > 0 {
				p2 := NewPoint(i+1, j)
				g.ConnectPoints(*p1, *p2, true)
			}

			if i-1 >= 0 && matrix[i][j] > 0 && matrix[i-1][j] > 0 {
				p2 := NewPoint(i-1, j)
				g.ConnectPoints(*p1, *p2, true)
			}

			if j+1 < len(matrix[i]) && matrix[i][j] > 0 && matrix[i][j+1] > 0 {
				p2 := NewPoint(i, j+1)
				g.ConnectPoints(*p1, *p2, true)
			}

			if j-1 >= 0 && matrix[i][j] > 0 && matrix[i][j-1] > 0 {
				p2 := NewPoint(i, j-1)
				g.ConnectPoints(*p1, *p2, true)
			}
		}
	}
}

func (g *TwoDGraph) AddPointIfNotExists(p1 Point) bool {
	if _, found := g.Hash[p1]; !found {
		g.Hash[p1] = make(map[Point]struct{})
		return true
	}
	return false
}

func (g *TwoDGraph) ConnectPoints(p1, p2 Point, bidirectional bool) {
	g.AddPointIfNotExists(p1)
	g.AddPointIfNotExists(p2)
	g.Hash[p1][p2] = struct{}{} //acts as a set to ensure uniqueness
	if bidirectional {
		g.Hash[p2][p1] = struct{}{} //connect opposite direction too!
	}
}

func (g *TwoDGraph) CountConnectedGraphs() int {
	// reset visited
	g.visited = map[Point]bool{}
	count := 0

	for k, _ := range g.Hash {
		if g.visited[k] {
			continue
		}
		count++
		g.visit(k)
	}

	return count
}

// private
func (g *TwoDGraph) visit(pt Point) {
	g.visited[pt] = true

	for k, _ := range g.Hash[pt] {
		if g.visited[k] {
			continue
		}
		g.visit(k)
	}
}
