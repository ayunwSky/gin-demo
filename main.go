package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"

	"gin_demo/config"
)

func main() {
	// Get all configs profile
	config, err := config.LoadConfig("./config/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Println(config.Oss)
	fmt.Println(config.Oss.Endpoint)
	fmt.Println(config.Oss.AccessKeyID)
	fmt.Println(config.Oss.AccessKeySecret)
	fmt.Println(config.Oss.BucketName)

	fmt.Println("----------")

	/*
		如果时间定义在gin路由的外面，那么每次访问 /health 接口，看到的都是一个静态的时间，就是说这个时间一直
		都是你第一次访问这个接口看到的时间，不会随着每次访问接口，这个时间就发现变化，返回当前访问接口时的时间。
		
		这是因为在 Gin 框架中，路由处理函数是在一个协程（goroutine）中执行的，而多次请求的处理函数可能在同一
		个协程中执行。因此，如果你在处理函数中只是返回一个静态的时间值，那么多次请求这个处理函数得到的时间值都是
		一样的，因为它们在同一个协程中执行。

		解决该问题的方法就是在处理函数中动态生成时间值。也就是定义到 /health 接口的里面
	*/
	// currentTime := time.Now().Format("2006-01-02 15:04:05")
	
	// Create a new gin router
	route := gin.Default()
	// Define a health check endpoint
	route.GET("/health", func(c *gin.Context) {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		c.JSON(http.StatusOK, gin.H{
			"statusCode": 200,
			"status":     "up",
			"currentTime": currentTime,
			"message":    "Health check OK!",
		})
	})

	// Start Gin Server
	route.Run(":8080")
}
