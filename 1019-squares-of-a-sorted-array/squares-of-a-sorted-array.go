func sortedSquares(A []int) []int {
	var squares []int
	for _, i := range A {
		squares = append(squares, i*i)
	}

	sort.Ints(squares)
	return squares
}