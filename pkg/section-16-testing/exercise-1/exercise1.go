package dog

// Years converts human years to dog years
func Years(n int) int {
	return n * 7
}

// Years2 converts human years to dog years
func Years2(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		count += 7
	}
	return count
}
