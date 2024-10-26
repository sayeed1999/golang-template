package main

import (
	"fmt"

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

	fmt.Println("\n\n====== Two D Graph using points (x,y) =======\n\n")

	var input = [][]int{{1, 1, 0}, {1, 0, 1}, {0, 1, 0}}
	// 1 1 0
	// 1 0 1
	// 0 1 0

	// form adjacency list of points
	// adjacent items vertically/horizontally are connected

	graph := twodgraph.NewTwoDGraph()

	graph.FormGraphFromMatrix(input)

	count := graph.CountConnectedGraphs()
	fmt.Println(count)
}
