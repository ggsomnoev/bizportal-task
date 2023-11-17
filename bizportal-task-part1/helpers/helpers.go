package helpers

// Helper function to calculate the absolute value
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// Helper function to calculate the sign of a number
func Sign(n int) int {
	if n < 0 {
		return -1
	} else if n > 0 {
		return 1
	}
	return 0
}
