package main

import "net/http"

func main() {
	// 設置API路由和處理程序
	http.HandleFunc("/api", testApi)

	// 啟動HTTP服務器，監聽在本地的8080端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
