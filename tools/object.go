package tools

//IntInSlice search a int in a slice of int
func IntInSlice(n int, list []int) bool {
	for _, v := range list {
		if v == n {
			return true
		}
	}
	return false
}
