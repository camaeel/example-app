package middleware

import (
	"database/sql"
	"fmt"
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

func InsertDB(setupDB func() (*sql.DB, error)) gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO: This is bad solution for handling database reloads.
		//It would be better to change the database for future requests,
		// but when/how to close the old connection???

		db, err := setupDB()
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
