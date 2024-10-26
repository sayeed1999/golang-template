package bidirectionalgraph

import "fmt"

type OneDGraph struct {
	Hash    map[int]map[int]bool
	visited map[int]bool
}

func NewOneDGraph() *OneDGraph {
	return &OneDGraph{
		Hash:    make(map[int]map[int]bool),
		visited: make(map[int]bool),
	}
}

// Note:- The matrix must have column length of 2, one src, one dest
func (g *OneDGraph) PopulateEdgesFromMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		g.PushEdges(matrix[i][0], matrix[i][1])
		g.PushEdges(matrix[i][1], matrix[i][0])
	}
}

func (g *OneDGraph) PushEdges(src, dest int) {
	if _, found := g.Hash[src]; !found {
		g.Hash[src] = make(map[int]bool)
	}
	g.Hash[src][dest] = true
}

func (g *OneDGraph) IsCircular() bool {
	// reset visited
	g.visited = map[int]bool{}

	for k, _ := range g.Hash {
		if g.visited[k] {
			continue
		}
		isCircular := g.traverseByDFS(k, -1) // need to pass a parent to track cycle for an undirected graph
		if isCircular {
			return true
		}
		// if not circular, let it traverse, otherwise stop!
	}

	return false
}

func (g *OneDGraph) traverseByDFS(node, parent int) bool {
	g.visited[node] = true

	for k, _ := range g.Hash[node] {
		if !g.visited[k] {
			isCircular := g.traverseByDFS(k, node)
			if isCircular {
				return true
			}
		} else if k != parent { // in a undirected graph, we dont consider 1=2 cyclic, otherwise all undirected graphs would be cyclic!
			fmt.Println(node, parent, k)
			return true
		}
	}

	return false
}
