package main

import (
	"log"
	"net/http"
	"os"
	"github.com/labstack/echo/v4"
	"github.com/joho/godotenv"
)

var URL, entry string

func handler(c echo.Context) error {
	log.Println("Request received")

	// `X-Forwarded-User` ヘッダーを取得
	user := c.Request().Header.Get("X-Forwarded-User")
	if user == "" {
		// ユーザー名が取得できなかった場合は "unknown" とする
		user = "unknown"
	}

	var userURL = URL + "&entry." + entry + "=" + user

	// リダイレクト
	return c.Redirect(http.StatusFound, userURL)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	URL = os.Getenv("URL")
	entry = os.Getenv("ENTRY")
	if URL == "" || entry == "" {
		log.Fatal("URL or ENTRY_ID not set in .env file")
	}

	// サーバーのルートハンドラーを設定
	e := echo.New()
	e.GET("/", handler)

	// サーバーをポート8080で開始
	e.Logger.Fatal(e.Start(":8080"))
}
