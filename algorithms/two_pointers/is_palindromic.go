package main

// isPalindrome checks whether a given string is a palindrome.
// A palindrome is a string that reads the same forward and backward.
func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

// Example Inputs and Outputs:
// isPalindrome("racecar") → true
// isPalindrome("abab")    → false
// isPalindrome("madam")   → true
