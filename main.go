package main

import (
	"log"
	"net/http"
)

// リダイレクト先のURL
var URL = "https://docs.google.com/forms/d/e/1FAIpQLScpiOMWtDtsZE_dapZGfHNcMlGfpdv2SQpDUVsB5nbM1G2E9w/viewform?usp=pp_url"

func handler(w http.ResponseWriter, r *http.Request) {
	// `X-Forwarded-User` ヘッダーを取得
	user := r.Header.Get("X-Forwarded-User")
	if user == "" {
		http.Error(w, "X-Forwarded-User header is missing", http.StatusBadRequest)
		return
	}

	var userURL = URL + "&entry.574283270=" + user

	// リダイレクト
	http.Redirect(w, r, userURL, http.StatusFound)
}

func main() {
	// サーバーのルートハンドラーを設定
	http.HandleFunc("/", handler)

	// サーバーをポート8080で開始
	log.Println("Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
