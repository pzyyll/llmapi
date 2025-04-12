package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Set up the router for the application

	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.Use(func(ctx *gin.Context) {
		// ctx.Header("Cache-Control", "no-cache")
		// ctx.Header("Access-Control-Allow-Origin", "*")
		// ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// ctx.Next()
		fmt.Println("M0N")
	})

	router.Use(func(ctx *gin.Context) {
		// ctx.Header("Cache-Control", "no-cache")
		// ctx.Header("Access-Control-Allow-Origin", "*")
		// ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// ctx.Next()
		fmt.Println("M1N")
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	login := router.Group("/login")
	login.Use(func(ctx *gin.Context) {
		fmt.Println("login middleware")
		ctx.Next()
		fmt.Println("login middleware end")
	})
	login.Use(func(ctx *gin.Context) {
		fmt.Println("login middleware 2")
	})
	{
		login.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "login success",
			})
		})
		login.GET("/test", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "login test success",
			})
		})
	}

	router.Use(func(ctx *gin.Context) {
		// ctx.Header("Cache-Control", "no-cache")
		// ctx.Header("Access-Control-Allow-Origin", "*")
		// ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// ctx.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// ctx.Next()
		fmt.Println("M2N")
	})

	// router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.NoRoute(func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache")
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World! NoRoute",
		})
	})

	if err := router.Run("localhost:13142"); err != nil {
		panic(err)
	}
}
