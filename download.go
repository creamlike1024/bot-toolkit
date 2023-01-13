package toolkit

import (
	"errors"
	"io"
	"net/http"
	"os"
)

func Download(filepath, url string, sizelimit int64) error {
	// 创建文件
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	// 获取文件大小
	headResp, err := http.Head(url)
	if err != nil {
		return err
	}
	defer headResp.Body.Close()
	if headResp.ContentLength > sizelimit {
		return errors.New("file size exceeds limit")
	}
	// 获取文件
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// 写入文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}
