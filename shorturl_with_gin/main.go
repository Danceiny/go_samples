package main

import (
	"github.com/gin-gonic/gin"
	"samples/shorturl_with_gin/controler"
	"fmt"
	"github.com/gin-gonic/gin/json"
)


func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	defer r.Handle("GET", "/api/v1/shorten", func(c *gin.Context) {
		c.JSON(200, "panic")
	})
	r.GET("/api/v1/shorten", func(c *gin.Context) {
		long_url := c.Query("long_url")
		if long_url == "" {
			//panic(nil)
			println("无参数")
		}
		res := controler.ShortenControler(long_url)
		println(long_url)
		fmt.Print(res)

		//c.JSON(200, gin.H{"short_url": res.UrlShort,
		//						"long_url": res.UrlLong})

		v, e := json.Marshal(res)
		if e != nil {
			println(v);
		}
		c.JSON(200, res)
	})
	r.Run()
}
