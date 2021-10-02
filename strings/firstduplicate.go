package main

func FirstDup(s string)rune{
	m := make(map[rune]bool)
	for _, char := range s{
		if m[char]{
			return char
		}
		m[char] = true
	}
	return ' '
}
