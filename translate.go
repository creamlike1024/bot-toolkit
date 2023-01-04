package toolkit

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"net/http"
	"time"
)

// 百度翻译开放平台信息
type BaiduTranslateInfo struct {
	AppID     string
	Salt      string
	SecretKey string
	From      string
	To        string
	Text      string
}

// 返回结果
type TranslateResult struct {
	From      string    `json:"from"`
	To        string    `json:"to"`
	Result    [1]Result `json:"trans_result"`
	ErrorCode string    `json:"error_code"`
	ErrorMsg  string    `json:"error_msg"`
}
type Result struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

func BaiduTranslate(bi *BaiduTranslateInfo) (string, error) {
	bi.Salt = salt(10)

	// 百度翻译接口
	url := "http://api.fanyi.baidu.com/api/trans/vip/translate?q=" + bi.Text + "&from=" + bi.From + "&to=" + bi.To + "&appid=" + bi.AppID + "&salt=" + bi.Salt + "&sign=" + sign(bi)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var tr TranslateResult
	err = json.Unmarshal(body, &tr)
	if err != nil {
		return tr.ErrorMsg, errors.New(tr.ErrorCode)
	}
	return tr.Result[0].Dst, nil
}

// 生成随机字符串
func salt(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 生成32位MD5
func sign(bi *BaiduTranslateInfo) string {
	text := bi.AppID + bi.Text + bi.Salt + bi.SecretKey
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}
