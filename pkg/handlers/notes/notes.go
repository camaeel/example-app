package notes

import (
	"database/sql"

	"github.com/camaeel/example-app/pkg/models/notes"
	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {

	value, exists := c.Get("db")
	if !exists {
		c.Abort()
	}
	db := value.(*sql.DB)
	notes, err := notes.List(db)
	if err != nil {
		c.AbortWithError(503, err)
	} else {
		c.JSON(200, notes)
	}
}

func Get(c *gin.Context) {
	value, exists := c.Get("db")
	if !exists {
		c.Abort()
	}
	db := value.(*sql.DB)
	id := c.Params.ByName("id")
	note, err := notes.Get(db, id)
	if err != nil {
		c.AbortWithError(503, err)
	} else {
		c.JSON(200, note)
	}
}

func Delete(c *gin.Context) {
	value, exists := c.Get("db")
	if !exists {
		c.Abort()
	}
	db := value.(*sql.DB)
	id := c.Params.ByName("id")
	err := notes.Delete(db, id)
	if err != nil {
		c.AbortWithError(503, err)
	} else {
		c.JSON(200, gin.H{"result": "ok"})
	}
}

func Create(c *gin.Context) {
	value, exists := c.Get("db")
	if !exists {
		c.Abort()
	}
	db := value.(*sql.DB)

	note := notes.Note{}

	if err := c.ShouldBindJSON(&note); err != nil {
		c.AbortWithError(503, err)
		return
	}

	err := notes.Create(db, note)

	if err != nil {
		c.AbortWithError(503, err)
		return
	}
	c.JSON(200, gin.H{"result": "ok"})

}

func Update(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
