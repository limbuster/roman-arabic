package main

import "testing"

var testCases = []struct {
	Arabic int
	Roman  string
}{
	{Arabic: 1, Roman: "I"},
	{Arabic: 2, Roman: "II"},
	{Arabic: 3, Roman: "III"},
	{Arabic: 4, Roman: "IV"},
	{Arabic: 5, Roman: "V"},
	{Arabic: 6, Roman: "VI"},
	{Arabic: 7, Roman: "VII"},
	{Arabic: 8, Roman: "VIII"},
	{Arabic: 9, Roman: "IX"},
	{Arabic: 10, Roman: "X"},
	{Arabic: 14, Roman: "XIV"},
	{Arabic: 18, Roman: "XVIII"},
	{Arabic: 20, Roman: "XX"},
	{Arabic: 39, Roman: "XXXIX"},
	{Arabic: 40, Roman: "XL"},
	{Arabic: 47, Roman: "XLVII"},
	{Arabic: 49, Roman: "XLIX"},
	{Arabic: 50, Roman: "L"},
	{Arabic: 90, Roman: "XC"},
	{Arabic: 100, Roman: "C"},
	{Arabic: 400, Roman: "CD"},
	{Arabic: 500, Roman: "D"},
	{Arabic: 900, Roman: "CM"},
	{Arabic: 1000, Roman: "M"},
	{Arabic: 1984, Roman: "MCMLXXXIV"},
	{Arabic: 3999, Roman: "MMMCMXCIX"},
	{Arabic: 2014, Roman: "MMXIV"},
	{Arabic: 1006, Roman: "MVI"},
	{Arabic: 798, Roman: "DCCXCVIII"},
}

func TestConvertToRoman(t *testing.T) {
	t.Run("convert arabic number to roman number", func(t *testing.T) {
		for _, c := range testCases {
			got := ConvertToRoman(c.Arabic)
			want := c.Roman

			if got != want {
				t.Errorf("want %q got %q", want, got)
			}
		}
	})

	t.Run("convert roman number to arabic number", func(t *testing.T) {
		for _, c := range testCases {
			got := ConvertToArabic(c.Roman)
			want := c.Arabic

			if got != want {
				t.Errorf("want %d got %d", want, got)
			}
		}
	})
}

func BenchmarkCalculate(b *testing.B) {
	b.Run("convert arabic number to roman number", func(b *testing.B) {
		for _, c := range testCases {
			got := ConvertToRoman(c.Arabic)
			want := c.Roman

			if got != want {
				b.Errorf("want %q got %q", want, got)
			}
		}
	})

	b.Run("convert roman number to arabic number", func(b *testing.B) {
		for _, c := range testCases {
			got := ConvertToArabic(c.Roman)
			want := c.Arabic

			if got != want {
				b.Errorf("want %d got %d", want, got)
			}
		}
	})
}
