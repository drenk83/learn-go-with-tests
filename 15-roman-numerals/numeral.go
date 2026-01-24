package romannumerals

import "strings"

type RomanNumeral struct {
	Arabic int
	Roman  string
}

var allRomanNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(num int) string {
	var res strings.Builder

	for _, v := range allRomanNumerals {
		if count := num / v.Arabic; count >= 1 {
			res.WriteString(strings.Repeat(v.Roman, count))
			num -= count * v.Arabic
		}
	}

	return res.String()
}
