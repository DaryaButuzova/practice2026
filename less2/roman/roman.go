package roman

var romanMap = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func RomanToArabic(s string) int {
	total := 0

	for i := 0; i < len(s); i++ {
		val := romanMap[rune(s[i])]

		if i+1 < len(s) && val < romanMap[rune(s[i+1])] {
			total -= val
		} else {
			total += val
		}
	}

	return total
}
