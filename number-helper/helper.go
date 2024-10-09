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

func CreateFrequencyMap(nums []int) map[int]int {
	hash := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, found := hash[nums[i]]; !found {
			hash[nums[i]] = 1
		} else {
			hash[nums[i]]++
		}
	}

	return hash
}
