package main

import (
	"net/http"
	"todo/api/todo"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router = attachCORS(router)

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	router.GET("todo", todo.GetAll)

	router.Run()
}

// attachApi CORSの設定
func attachCORS(router *gin.Engine) *gin.Engine {
	// CORSの設定
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowHeaders = []string{"*"}
	router.Use(cors.New(corsConfig))
	return router
}
