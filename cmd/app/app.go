package main

import (
	"flag"
	"fmt"
	"github.com/camaeel/example-app/pkg/config"
	"github.com/camaeel/example-app/pkg/database"
	"github.com/camaeel/example-app/pkg/handlers/health"
	"github.com/camaeel/example-app/pkg/handlers/notes"
	"github.com/camaeel/example-app/pkg/middleware"
	notesModel "github.com/camaeel/example-app/pkg/models/notes"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

func main() {

	port := 8080

	configFile := flag.String("configFile", "config/database.yml", "Path to config file")
	flag.Parse()
	config.LoadConfig(configFile)
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
		config.LoadConfig(configFile)
	})

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
	r.Use(middleware.InsertDB())

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
