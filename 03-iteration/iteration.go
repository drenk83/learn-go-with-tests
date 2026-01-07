package iteration

import "strings"

func Repeat(str string, repeatCount int) string {
	var out strings.Builder
	out.Grow(repeatCount)

	for range repeatCount {
		out.WriteString(str)
	}
	return out.String()
}