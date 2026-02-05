package service

import (
	"unicode"

	"github.com/Jorka2099/homework_jan/pkg/morse"
)

var StringIsMorse bool

// HasLetters определяет есть ли буквы в строке
func HasLetters(s string) bool {
	for _, r := range s {
		upper := unicode.ToUpper(r)

		if (upper >= 'А' && upper <= 'Я') || upper == 'Ё' {
			return true
		}

		if upper >= 'A' && upper <= 'Z' {
			return true
		}
	}

	return false
}

// TextOrMorse определяет является ли строка текстом или кодом Морзе
func TextOrMorse(parsedStr string) string {
	if HasLetters(parsedStr) {
		return morse.ToMorse(parsedStr)
	}

	return morse.ToText(parsedStr)
}
