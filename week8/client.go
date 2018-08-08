package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	println("====取文案內容====")
	// 取文案內容
	getArticle()

	println("====設定文案====")
	// 設定文案內容
	setArticle()

	println("====取文案內容====")
	// 取文案內容
	getArticle()
}

// 取文案內容
func getArticle() {
	client := &http.Client{}
	//向服务端发送get请求
	request, _ := http.NewRequest("GET", "http://localhost:1323/get/article", nil)
	//接收服务端返回给客户端的信息
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		str, _ := ioutil.ReadAll(response.Body)
		bodystr := string(str)
		fmt.Println(bodystr)
	}
}

// 設定文案內容
func setArticle() {
	client := &http.Client{}
	//post请求
	postValues := url.Values{}
	postValues.Add("title", "烈 女 操")
	postValues.Add("content", "梧 桐 相 待 老 ，鴛 鴦 會 雙 死 。貞 婦 貴 殉 夫 ，捨 生 亦 如 此 。波 瀾 誓 不 起 ，妾 心 古 井 水 。")

	resp, err := client.PostForm("http://localhost:1323/post/article", postValues)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
	}
}
