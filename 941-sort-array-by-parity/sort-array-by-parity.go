func sortArrayByParity(A []int) []int {
	for i, j := 0, len(A)-1; i < j; {
		if A[i]%2 == 1 && A[j]%2 == 0 {
			A[i], A[j] = A[j], A[i]
			i++
			j--
		} else if A[j]%2 == 1 {
			j--
		} else if A[i]%2 == 0 {
			i++
		}
	}

	return A 
}