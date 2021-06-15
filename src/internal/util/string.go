package util

func SubString(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || end < 0 || start > end {
		return ""
	}

	if start == 0 && end >= length {
		return source
	}

	return string(r[start:end]) + "..."
}
