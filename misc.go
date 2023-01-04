package toolkit

import "strings"

func MergeText(args ...string) string {
	// 用换行连接参数
	var text string
	for _, arg := range args {
		text += arg + "\n"
	}
	return text
}

func StringSliceToText(stringSlice []string) string {
	text := strings.Join(stringSlice, "\n")
	return text
}
