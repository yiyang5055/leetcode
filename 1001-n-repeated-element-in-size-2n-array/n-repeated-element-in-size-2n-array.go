func repeatedNTimes(A []int) int {
	var res int
	timesCount := make(map[int]int)

	N := len(A) / 2
	for _, n := range A {
		timesCount[n] = timesCount[n] + 1
		if timesCount[n] == N {
			res = n
		}
	}

	return res   
}