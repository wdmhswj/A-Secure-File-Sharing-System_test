package userlib_client

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func client() {
	// 向指定 URL 发送 HTTP GET 请求
	resp, err := http.Get("http://example.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// 输出服务器响应的状态码和内容
	fmt.Println("Status Code:", resp.Status)
	fmt.Println("Response Body:")
	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
	}
}
