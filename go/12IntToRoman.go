package main

var IntToRoman = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}
var IntToRomanArr = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

func intToRoman(num int) string {
	answer := ""
	for _, val := range IntToRomanArr {
		hello := num / val
		for hello > 0 {
			answer += IntToRoman[val]
			hello--
		}
		num %= val
	}
	return answer
}
