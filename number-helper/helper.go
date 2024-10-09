package numberhelper

// Note: This will work for array with unique values only.
func CreateValueToIndexLookupHash(nums []int) map[int]int {
	// golang's built-in 'map' implements hash table
	hash := map[int]int{}

	// create lookup hash
	for i := 0; i < len(nums); i++ {
		hash[nums[i]] = i
	}
	return hash
}
