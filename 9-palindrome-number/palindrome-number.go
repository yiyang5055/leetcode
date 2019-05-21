func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	n := x
	res := 0
	for n > 0 {
		res = res*10 + n%10
		n = n / 10
	}
	return x == res
}
