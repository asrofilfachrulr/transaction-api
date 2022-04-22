package helpers

// dont care about order
func RemoveItemFromSlice(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
