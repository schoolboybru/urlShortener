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

	urlCache, err := database.New()

	if err != nil {
		panic(err)
	}

	urlHandler := routes.NewHandler(urlCache)

	r.GET("/url", urlHandler.Get)
	r.POST("/url", urlHandler.Add)

	r.Run(":8080")
}
