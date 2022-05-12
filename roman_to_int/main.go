package main

var chars = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func Decode(roman string) int {
	if roman == "" {
		return 0
	}
	value := chars[roman[0]]
	for index := 1; index < len(roman); index++ {
		if last, now := chars[roman[index-1]], chars[roman[index]]; last < now {
			value += now - 2*last
		} else {
			value += now
		}
	}
	return value
}
