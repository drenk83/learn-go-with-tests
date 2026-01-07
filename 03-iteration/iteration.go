package iteration

import "strings"

func Repeat(str string, repeatCount int) string {
	var out strings.Builder
	out.Grow(len(str) * repeatCount)

	for i := 0; i < repeatCount; i++ {
		out.WriteString(str)
	}
	return out.String()
}