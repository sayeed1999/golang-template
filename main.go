package main

import (
	"fmt"

	bidirectionalgraph "github.com/sayeed1999/golang-template/bidirectional-graph"
	twodgraph "github.com/sayeed1999/golang-template/twoD-graph"
)

func main() {
	// str := "sayeed"
	// s := stringhelper.SortString(str)
	// fmt.Println(s)

	// str = "sayeed"
	// x := stringhelper.GetLowercaseCharFrequencyMap(str)
	// fmt.Println(x)

	// str = "SAYEED"
	// x = stringhelper.GetUppercaseCharFrequencyMap(str)
	// fmt.Println(x)

	// nums := []int{1, 3, 2, 4}
	// hash := numberhelper.CreateValueToIndexLookupHash(nums)
	// fmt.Println(hash)

	// nums = []int{1, 1, 2, 2, 1, 3}
	// hash = numberhelper.CreateFrequencyMap(nums)
	// fmt.Println(hash)

	// parenthesis := "({})"
	// isValid := stack.IsValidParenthesis(parenthesis)
	// fmt.Println(isValid)

	// parenthesis = "({])"
	// isValid = stack.IsValidParenthesis(parenthesis)
	// fmt.Println(isValid)

	// fmt.Println("\n\n======= Built in generic binary search in go's slice package =======")
	// sortedArray := []int{1, 2, 3, 4, 5, 6}
	// index, found := slices.BinarySearch(sortedArray, 4)
	// fmt.Println(index, found, "found element through go's built-in binary search")
	// index, found = slices.BinarySearch(sortedArray, 8)
	// fmt.Println(index, found, "ignore the index its not constant for missing elements")

	fmt.Println("\n====== Two D Graph using points (x,y) =======\n")

	var matrix = [][]int{{1, 1, 0}, {1, 0, 1}, {0, 1, 0}}
	// 1 1 0
	// 1 0 1
	// 0 1 0

	// form adjacency list of points
	// adjacent items vertically/horizontally are connected

	graph2D := twodgraph.NewTwoDGraph()

	graph2D.FormGraphFromMatrix(matrix)

	count := graph2D.CountConnectedGraphs()
	fmt.Println(count)

	fmt.Println("\n====== Single Dimensional Undirected/Bidirectional Graph =======\n")

	input := [][]int{{2, 1}, {3, 1}}
	graph := bidirectionalgraph.NewOneDGraph()
	graph.PopulateEdgesFromMatrix(input)
	isCircular := graph.IsCircular()
	fmt.Printf("The single dimensional graph is circular: %v\n", isCircular)
}
