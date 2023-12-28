package main

import (
	"encoding/json"
	"net/http"
)

// 定義一個結構表示API的數據
type Response struct {
	Message string `json:"message"`
}

// 處理API請求的處理程序
func handleAPIRequest(w http.ResponseWriter, r *http.Request) {
	// 創建一個Response結構的實例
	response := Response{
		Message: "Hello, this is a JSON API!",
	}

	// 將結構序列化為JSON格式
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// 設置HTTP標頭並返回JSON數據
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
