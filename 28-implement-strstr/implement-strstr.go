func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}

	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		k := i
		has := false
		for j := 0; j < len(needle); j++ {
			if needle[j] != haystack[k] {
				has = false
				break
			}
			k += 1
			has = true
		}
		if has {
			return i
		}
	}
	return -1
}
