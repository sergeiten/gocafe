package util

import (
	"regexp"
	"strings"
)

// ReplaceHanhulPipe replaces all Korean pipe char (ㅣ) occurrences with Latin pipe char (|)
// They are look similar but has different unicode
func ReplaceHangulPipe(str string) string {
	runes := []rune(str)
	for k, r := range runes {
		if r == 0x3163 { // hangul char
			runes[k] = 0x7C // replace with latin pipe
		}
	}

	return string(runes)
}

func GetBrand(str string) string {
	return submatch(`\[(\W+)\]`, str)
}

func GetName(str string) string {
	return submatch(`\[\W+\]\|?([\p{L}\d_\s\(\)\.]+)`, str)
}

func GetYear(str string) string {
	return submatch(`([\d]+)년`, str)
}

func GetDistance(str string) string {
	return submatch(`([\d\,\s]+)km`, str)
}

func GetPrice(str string) string {
	return submatch(`([\d\,\s]+)만원`, str)
}

func submatch(reg string, str string) string {
	r := regexp.MustCompile(reg)

	f := r.FindStringSubmatch(str)

	if len(f) > 1 {
		return strings.Trim(f[1], " ")
	}

	return ""
}
