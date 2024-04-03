package userlib_server

import (
	"fmt"
	"net/http"
)

func main() {
	// 注册处理函数
	http.HandleFunc("/datastoreGet", postHandler_datastoreGet)

	// 启动 HTTP 服务器，监听在 8080 端口
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
