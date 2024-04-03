package userlib_server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func postHandler_datastoreGet(w http.ResponseWriter, r *http.Request) {
	// 仅处理 POST 请求
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 读取请求体中的数据
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// 处理 POST 数据
	log.Printf("Received POST data: %s", string(body))
	// 反序列化
	// 创建结构体变量，用于存储解码后的数据
	var key UUID
	// 使用解码器将 JSON 数据反序列化为多个结构体对象
	err = json.Unmarshal(body, &key)
	if err != nil {
		log.Println("Error decoding person:", err)
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}
	value, ok := datastoreGet(key)

	if !ok {
		log.Println("The value does not exist")
		http.Error(w, "The value does not exist", http.StatusBadRequest)
		return
	}
	// 返回响应
	w.WriteHeader(http.StatusOK)
	w.Write(value)
}
