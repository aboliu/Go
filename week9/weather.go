package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"time"
)

type Weather struct {
	Create_time string `gorm:"type:time.Time;"`
	Information string `gorm:"type:text;"`
}

func main() {
	e := echo.New()
	e.GET("/weather", getWeather)

	e.Logger.Fatal(e.Start(":1111"))
}

// 顯示資料來源用
var source string

var now string

// 取天氣資料（作業要求: redis -> DB -> api）
func getWeather(c echo.Context) error {
	// 今天日期
	now = time.Now().Format("2006-01-02")

	// redis
	source = "redis"
	response := getRedis()

	// 若Redis沒資料，改取DB
	if response == "false" {
		// 取DB
		source = "DB"
		response = getDB()

		// 若DB沒資料，改取API
		if response == "false" {
			// 取API
			source = "API"
			response = getAPI()

			// 寫進DB
			setDB(response)
		}

		// 寫進Redis
		setRedis(response)
	}

	return c.String(http.StatusOK, "資料來源:"+source+"\n"+response)
}

func getRedis() string {
	// Redis 建立連線
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return "Connect to redis error"
	}
	// 執行結束時關閉連線
	defer c.Close()

	// 確認Redis是否存在該key
	isKeyExit, err := redis.Bool(c.Do("EXISTS", now))
	if err != nil {
		return "error"
	}

	// 如果存在，取Redis的資料
	if isKeyExit {
		// 讀取Redis
		weather, err := redis.String(c.Do("GET", now))
		if err != nil {
			return "redis get failed"
		} else {
			return weather + "\n"
		}
	}

	return "false"
}

// Redis補資料
func setRedis(data string) string {
	// Redis 建立連線
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return "Connect to redis error"
	}
	// 執行結束時關閉連線
	defer c.Close()

	// 寫進Redis，保存時間: 5
	_, err = c.Do("SET", now, data, "EX", "5")
	if err != nil {
		return "redis set failed"
	}

	return "secess"
}

func getDB() string {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/weather")
	if err != nil {
		return "failed to connect database"
	}
	defer db.Close()

	// 讀DB資料
	var weather Weather

	// db.First(&weather, 1) // find product with id 1
	db.First(&weather, "create_time = ?", now) // find product with code l1212

	if weather.Information != "" {
		return weather.Information
	}

	return "false"
}

// 寫進DB
func setDB(data string) string {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/weather")
	if err != nil {
		return "failed to connect database"
	}
	defer db.Close()

	// 要寫進DB的資料設定
	weather := &Weather{
		Create_time: now,
		Information: data,
	}

	if err := db.Create(&weather).Error; err != nil {
		return "db set err"
	}

	return "secess"
}

func getAPI() string {
	client := &http.Client{}
	//向服务端发送get请求
	request, _ := http.NewRequest("GET", "http://weather.json.tw/api", nil)
	//接收服务端返回给客户端的信息
	response, _ := client.Do(request)
	if response.StatusCode == 200 {
		str, _ := ioutil.ReadAll(response.Body)
		bodystr := string(str)
		return bodystr
	}

	return ""
}

// 執行service方式: go run weather.go
