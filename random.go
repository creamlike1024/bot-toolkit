package toolkit

import (
	"math/rand"
	"time"
)

func GenRandomInt(a, b int) int {
	rand.Seed(time.Now().UnixNano())
	// 生成指定范围的随机整数
	if a > b {
		a, b = b, a
	}
	return rand.Intn(b-a) + a
}

func GenRandomKey(l int) string {
	rand.Seed(time.Now().UnixNano())
	// 生成指定长度的随机字符串
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, l)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
