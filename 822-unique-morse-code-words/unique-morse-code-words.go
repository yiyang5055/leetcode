func uniqueMorseRepresentations(words []string) int {
	morseCode := []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
	res := make(map[string]string)

	for _, word := range words {
		morse := ""
		for _, s := range []rune(word) {
			morse = morse + morseCode[s-'a']
		}

		res[morse] = ""
	}

	return len(res)    
}