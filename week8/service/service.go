package main

import (
	"github.com/labstack/echo"
	"net/http"
)

var ArticleTitle string = "尚未設定文案內容"
var ArticleContent string

func main() {
	e := echo.New()
	e.GET("/get/article", getArtcle)
	e.POST("/post/article", postArtcle)

	e.Logger.Fatal(e.Start(":1323"))
}

// 取文案內容
func getArtcle(c echo.Context) error {
	response := "文案標題: " + ArticleTitle + "\r\n" + ArticleContent

	return c.String(http.StatusOK, response)
}

// 設定文案內容
func postArtcle(c echo.Context) error {
	ArticleTitle = c.FormValue("title")
	ArticleContent = c.FormValue("content")

	return c.String(http.StatusOK, "成功設定文案內容")
}

// 執行service方式: go run service.go
