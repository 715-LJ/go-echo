package check

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Path(path string) bool {
	_, err := os.Stat(path)
	if path != "" && err != nil {
		dir := filepath.Dir(path)

		err = os.Mkdir(dir, os.ModePerm)
		IfError(err)

		_, err := os.Create(path)
		IfError(err)
		return false
	}
	return true
}

// 下载文件到本地
func DownloadFile(url string, localFilePath string) error {
	// 发送 HTTP GET 请求
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to make GET request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// 创建本地文件
	out, err := os.Create(localFilePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	//defer out.Close()

	// 将下载的文件内容写入本地文件
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save file: %v", err)
	}

	fmt.Println("File downloaded successfully:", localFilePath)
	return nil
}
