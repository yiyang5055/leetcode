package main

import (
	"fmt"
	"sort"
)

func sortedSquares(A []int) []int {
	var squares []int
	for _, i := range A {
		squares = append(squares, i*i)
	}

	sort.Ints(squares)
	return squares
}

func main() {
	//n := []int{-4, -1, 0, 3, 10}
	n := []int{-7, -3, 2, 3, 11}
	fmt.Println(sortedSquares(n))
}
