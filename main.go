package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/images", "./images")
	router.Use(notFound)
	router.GET("/", maintenance)
	router.Run(":9070")
}

func notFound(c *gin.Context) {
	if c.Writer.Status() == 404 {
		maintenance(c)
	}
}

func maintenance(c *gin.Context) {
	data := map[string]interface{}{
		"header": "Сайт на обслуживании",
	}
	c.HTML(200, "maintenance.html", data)
}
