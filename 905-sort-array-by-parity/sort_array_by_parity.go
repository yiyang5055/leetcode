package main

import "fmt"

/*
Given an array A of non-negative integers,
return an array consisting of all the even elements of A,
followed by all the odd elements of A.

You may return any answer array that satisfies this condition.

Example 1:
	Input: [3,1,2,4]
	Output: [2,4,3,1]
	The outputs [4,2,3,1], [2,4,1,3], and [4,2,1,3] would also be accepted.

Note:
	1 <= A.length <= 5000
	0 <= A[i] <= 5000
*/

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

func main() {
	A := []int{3, 1, 2, 4}
	fmt.Println(sortArrayByParity(A))
	A = []int{0, 1, 2}
	fmt.Println(sortArrayByParity(A))
	A = []int{1, 0, 3}
	fmt.Println(sortArrayByParity(A))
}
