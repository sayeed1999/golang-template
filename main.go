package main

import (
	"fmt"

	numberhelper "github.com/sayeed1999/golang-template/number-helper"
	stringhelper "github.com/sayeed1999/golang-template/string-helper"
)

func main() {
	str := "sayeed"
	s := stringhelper.SortString(str)
	fmt.Println(s)

	str = "sayeed"
	x := stringhelper.GetLowercaseCharFrequencyMap(str)
	fmt.Println(x)

	str = "SAYEED"
	x = stringhelper.GetUppercaseCharFrequencyMap(str)
	fmt.Println(x)

	nums := []int{1, 3, 2, 4}
	lookupHash := numberhelper.CreateValueToIndexLookupHash(nums)
	fmt.Println(lookupHash)
}
