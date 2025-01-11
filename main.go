package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

// リダイレクト先のURL
var URL = "https://docs.google.com/forms/d/e/1FAIpQLScpiOMWtDtsZE_dapZGfHNcMlGfpdv2SQpDUVsB5nbM1G2E9w/viewform?usp=pp_url"

func handler(c echo.Context) error {
	log.Println("Request received")

	// `X-Forwarded-User` ヘッダーを取得
	user := c.Request().Header.Get("X-Forwarded-User")
	if user == "" {
		// ユーザー名が取得できなかった場合は "unknown" とする
		user = "unknown"
	}

	var userURL = URL + "&entry.574283270=" + user

	// リダイレクト
	return c.Redirect(http.StatusFound, userURL)
}

func main() {
	// サーバーのルートハンドラーを設定
	e := echo.New()
	e.GET("/", handler)

	// サーバーをポート8080で開始
	e.Logger.Fatal(e.Start(":8080"))
}
