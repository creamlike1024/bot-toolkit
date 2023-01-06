package toolkit

import "time"

func TimestampToTime(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	normal := t.UTC().Format("2006-01-02 15:04:05")
	rfc1123 := t.UTC().Format(time.RFC1123)
	rc3399 := t.UTC().Format(time.RFC3339)
	text := MergeText("以 UTC 时区输出", "普通格式：", normal, "\nRFC1123 格式: ", rfc1123, "\nRFC3339 格式: ", rc3399)
	return text
}
