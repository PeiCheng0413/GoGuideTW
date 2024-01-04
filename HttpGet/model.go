package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// 定義一個結構表示API的數據
type Response struct {
	Message string `json:"message"`
}

type waterLevel struct {
	Info []waterLavelInfo `json:"data"`
}

type waterLavelInfo struct {
	StationName string  `json:"stationName"`
	LevelOut    float32 `json:"levelOut"`
}

// 處理API請求的處理程序
func testApi(w http.ResponseWriter, r *http.Request) {
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

// 對指定網址request之後把response存成檔案
func urlGetAndSave(url string) {
	//發出requset後儲存response data
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}

	data, err := io.ReadAll(res.Body)
	// 設置API路由和處理程序

	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("httpGetTest.html", data, 0666)

	if err != nil {
		log.Fatal(err)
	}
}

func getRiverLevel() {
	resp, err := http.Get("https://wic.heo.taipei/OpenData/API/Water/Get?stationNo=&loginId=river&dataKey=9E2648AA")
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var wl waterLevel

	err = json.Unmarshal(data, &wl)

	if err != nil {
		log.Fatal(err)
	}

	for _, w := range wl.Info {
		fmt.Println("地點", w.StationName, "水高", w.LevelOut)
	}
}
