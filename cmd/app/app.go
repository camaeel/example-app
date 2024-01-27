package main

import (
	"fmt"

	"github.com/camaeel/example-app/pkg/database"
	"github.com/camaeel/example-app/pkg/handlers/health"
	"github.com/camaeel/example-app/pkg/handlers/notes"
	"github.com/camaeel/example-app/pkg/middleware"
	notesModel "github.com/camaeel/example-app/pkg/models/notes"
	"github.com/gin-gonic/gin"
)

func main() {

	port := 8080

	db, err := database.SetupDriver()
	defer db.Close()
	if err != nil {
		panic(err)
	}
	err = notesModel.InitializeTable(db)
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.Use(middleware.Logger())
	r.Use(middleware.InsertDB(db))

	r.GET("/healthz", health.Healthz)
	r.GET("/notes", notes.List)
	r.GET("/notes/:id", notes.Get)
	r.POST("/notes", notes.Create)
	r.PUT("/notes/:id", notes.Update)
	r.DELETE("/notes/:id", notes.Delete)

	//start server
	err = r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
}
