package main

import (
	"strings"
)

type romanNumeral struct {
	value  int
	symbol string
}

var romanNumerals = []romanNumeral{
	{value: 1000, symbol: "M"},
	{value: 900, symbol: "CM"},
	{value: 500, symbol: "D"},
	{value: 400, symbol: "CD"},
	{value: 100, symbol: "C"},
	{value: 90, symbol: "XC"},
	{value: 50, symbol: "L"},
	{value: 40, symbol: "XL"},
	{value: 10, symbol: "X"},
	{value: 9, symbol: "IX"},
	{value: 5, symbol: "V"},
	{value: 4, symbol: "IV"},
	{value: 1, symbol: "I"},
}

// ConvertToRoman takes arabic numeral and returns roman string
func ConvertToRoman(arabic int) string {
	roman := strings.Builder{}
	for arabic > 0 {
		for _, n := range romanNumerals {
			for arabic >= n.value {
				arabic -= n.value
				roman.WriteString(n.symbol)
			}
		}
	}

	return roman.String()
}

type decrementSet struct {
	prevChar string
	currChar string
}

type decrementSets []decrementSet

var dSets = decrementSets{
	{prevChar: "V", currChar: "I"},
	{prevChar: "X", currChar: "I"},
	{prevChar: "L", currChar: "X"},
	{prevChar: "C", currChar: "X"},
	{prevChar: "D", currChar: "C"},
	{prevChar: "M", currChar: "C"},
}

func (d decrementSets) contains(s decrementSet) bool {
	for _, item := range d {
		if item == s {
			return true
		}
	}
	return false
}

func needToDecrement(prevChar string, currChar string) bool {

	return dSets.contains(decrementSet{prevChar: prevChar, currChar: currChar})
}

// ConvertToArabic takes roman symbol and return roman numeral
func ConvertToArabic(roman string) int {
	var result int
	var prevChar string
	for i := len(roman) - 1; i >= 0; i-- {
		character := string(roman[i])
		for _, n := range romanNumerals {
			if needToDecrement(prevChar, character) && n.symbol == character {
				result -= n.value
			} else if n.symbol == character {
				result += n.value
			}
		}

		if i > 0 {
			prevChar = character
		}
	}
	return result
}
