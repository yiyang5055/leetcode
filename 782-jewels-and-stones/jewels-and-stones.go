func numJewelsInStones(J string, S string) int {
	var jewels int
	for _, s := range []rune(S) {
		if strings.ContainsRune(J, s) {
			jewels++
		}
	}
	return jewels  
}