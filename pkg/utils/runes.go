package utils

const ASCII_UPPER_LIMIT = 127

func ExtractRune(text string) rune {
	result := []rune(text)
	return result[0]
}

func IsAscii(r rune) bool {
	return r < ASCII_UPPER_LIMIT
}
