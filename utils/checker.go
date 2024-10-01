package utils

// Cleanstring removes all unprintable characters from a string
func Cleanstring(s string) string {
	var cleanString []rune

	for _, r := range s {
		if r >= 32 && r <= 126 { // Check if the character is printable
			cleanString = append(cleanString, r)
		}
	}

	return string(cleanString)
}

func Isempty(text []string) bool {
	for _, l := range text {
		if l != "" {
			return false
		}
	}
	return true
}
