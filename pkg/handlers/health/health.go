package health

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

func Healthz(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func Readyz(c *gin.Context) {
	value, exists := c.Get("db")
	if !exists {
		c.Abort()
	}
	db := value.(*sql.DB)
	_, err := db.Exec("SELECT 1 FROM notes")
	if err != nil {
		c.AbortWithError(503, err)
	} else {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	}
}
