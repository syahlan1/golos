package authService

import "strings"

func isValidUsername(username string, minLength int, validChars, validNums string) bool {
	if len(username) < minLength {
		return false
	}

	// Cek apakah username mengandung setidaknya satu huruf yang valid dan satu angka yang valid
	hasLetter := false
	hasNumber := false

	for _, char := range username {
		if strings.ContainsRune(validChars, char) {
			hasLetter = true
		}
		if strings.ContainsRune(validNums, char) {
			hasNumber = true
		}
		if hasLetter && hasNumber {
			return true
		}
	}

	return false
}
