package health

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
)

func Healthz(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func Readyz(c *gin.Context) {
	value, exists := c.Get("db")
	if !exists {
		log.Printf("Db not found in the context")
		c.Abort()
	}
	db := value.(*sql.DB)
	_, err := db.Exec("SELECT 1 FROM notes")
	if err != nil {
		log.Printf("ERROR: Database query failed with: %v", err)
		c.AbortWithError(503, err)
	} else {
		log.Printf("Readiness probe ok")
		c.JSON(200, gin.H{
			"message": "ok",
		})
	}
}
