package toolkit

import "encoding/base64"

func Base64Encode(data []byte) string {
	enc := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	return enc.EncodeToString(data)
}

func Base64Decode(data string) ([]byte, error) {
	enc := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
	return enc.DecodeString(data)
}
