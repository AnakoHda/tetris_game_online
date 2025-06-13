package main

import (
	"auth-service/storage"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	if err := storage.InitPostgres(); err != nil {
		log.Fatal("DB connect error:", err)
	}

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	r.Run(":8080")
}
