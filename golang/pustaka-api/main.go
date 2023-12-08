package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOk, gin.H{
					"name": "Rizky Ziaul Haq",
					"bio": "Pelajar",
			})
	})

	router.Run()
}