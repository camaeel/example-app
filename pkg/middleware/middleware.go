package middleware

import (
	"fmt"
	"github.com/camaeel/example-app/pkg/database"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// before request
		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func InsertDB() gin.HandlerFunc {
	return func(c *gin.Context) {
		db, err := database.SetupDriver()
		defer db.Close()

		if err != nil {
			log.Printf("ERROR: connecting to database: %v", err)
			c.AbortWithError(500, fmt.Errorf("Couldn't connect to database"))
		}
		c.Set("db", db)

		// before request
		c.Next()
	}
}
