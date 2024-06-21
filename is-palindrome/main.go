package main

func isPalindrome(word string) bool {
	const minTrueLength = 2

	runes := []rune(word)

	length := len(runes)
	if length < minTrueLength {
		return true
	}

	for i := range runes {
		if runes[i] != runes[length-1-i] {
			return false
		}

		if i == length/2 {
			break
		}
	}

	return true
}

func main() {}
