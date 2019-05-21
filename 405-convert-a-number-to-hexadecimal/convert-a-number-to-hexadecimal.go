func toHex(num int) string {
	s := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	if num == 0 {
		return "0"
	}

	res := ""
	shift := 0
	for num != 0 && shift < 8 {
        res = string(s[(num&15)]) + res
		num = num >> 4
        shift++
	}
	return res
}
