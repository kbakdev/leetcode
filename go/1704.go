package _go

func halvesAreAlike(s string) bool {
	var count int
	var vowels = map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	for i := 0; i < len(s)/2; i++ {
		if vowels[s[i]] {
			count++
		}
		if vowels[s[len(s)-1-i]] {
			count--
		}
	}
	return count == 0
}
