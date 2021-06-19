package util

import "os"

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

func AddOSEnviron(envs map[string]string) []string {
	osEnvs := os.Environ()
	for k := range envs {
		osEnvs = append(osEnvs, k+"="+envs[k])
	}
	return osEnvs
}
