package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/schoolboybru/urlShortener/internal/database"
	"github.com/schoolboybru/urlShortener/internal/routes"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	urlStore, err := database.New()

	if err != nil {
		panic(err)
	}

	urlHandler := routes.NewHandler(urlStore)

	r.GET("/url", urlHandler.GetUrl)
	r.POST("/url", urlHandler.AddUrl)

	r.Run(":8080")
}
