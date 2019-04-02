package main

import "fmt"

/*
In a array A of size 2N, there are N+1 unique elements,
and exactly one of these elements is repeated N times.

Return the element repeated N times.

Example 1:
	Input: [1,2,3,3]
	Output: 3

Example 2:
	Input: [2,1,2,5,3,2]
	Output: 2

Example 3:
	Input: [5,1,5,2,5,3,5,4]
	Output: 5

Note:
	4 <= A.length <= 10000
	0 <= A[i] < 10000
	A.length is even
*/

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

func main() {
	//A := []int{1, 2, 3, 3}
	A := []int{2, 1, 2, 5, 3, 2}
	//A := []int{5, 1, 5, 2, 5, 3, 5, 4}

	fmt.Println(repeatedNTimes(A))
}
