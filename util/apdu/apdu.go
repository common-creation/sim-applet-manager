package apdu

func SwapPairs(input string) string {
	if len(input)%2 != 0 {
		return input
	}

	result := make([]rune, len(input))
	for i := 0; i < len(input); i += 2 {
		if i+1 < len(input) {
			result[i], result[i+1] = rune(input[i+1]), rune(input[i])
		}
	}

	return string(result)
}
