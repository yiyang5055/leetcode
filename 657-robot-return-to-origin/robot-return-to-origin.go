func judgeCircle(moves string) bool {
	moveRoad := map[rune]int{
		'R': 1,
		'L': -1,
		'U': 2,
		'D': -2,
	}

	var sum int
	for _, move := range []rune(moves) {
		sum += moveRoad[move]
	}

	if sum == 0 {
		return true
	}
	return false    
}