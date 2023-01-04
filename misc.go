package toolkit

func MergeText(args ...string) string {
	// 用换行连接参数
	var text string
	for _, arg := range args {
		text += arg + "\n"
	}
	return text
}
