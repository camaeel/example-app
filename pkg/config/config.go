package config

import (
	"github.com/spf13/viper"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"
)

type RawDatasourceConfig struct {
	Jdbc_Url string
	Username string
	Password string
}

type Config struct {
	DatasourceUrl *url.URL
}

var mutex sync.RWMutex
var config *Config

const dbUrlEnv = "DATABASE_URL"

func LoadConfig(file *string) {
	mutex.Lock()

	dsCfg := RawDatasourceConfig{}

	defer mutex.Unlock()
	if file != nil && *file != "" {
		log.Printf("Trying to load config from %s", *file)
		viper.SetConfigName(*file)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Error reading config file %s: %s", *file, err)
		}

		err = viper.Unmarshal(&dsCfg)
		if err != nil {
			log.Fatalf("Unable to decode datasource config into struct: %s", err)
		}

		//schemes := strings.Split(u.Scheme, ":")
		//driverName := schemes[len(schemes)-1]

	} else {
		log.Print("No config file specified. Reading ENV variable: DATABASE_URL")
		dsCfg.Jdbc_Url = os.Getenv(dbUrlEnv)
	}

	// normalize
	dsCfg.Jdbc_Url = strings.Replace(dsCfg.Jdbc_Url, "jdbc:", "", 1)

	dsUrl, err := url.Parse(dsCfg.Jdbc_Url)
	if err != nil {
		log.Fatalf("Unable to parse datasource url: %s", err)
	}
	// put user & password if provided
	if dsCfg.Username != "" && dsCfg.Password != "" {
		dsUrl.User = url.UserPassword(dsCfg.Username, dsCfg.Password)
	}

	config = &Config{
		DatasourceUrl: dsUrl,
	}

	log.Print("Config loaded")
}

func GetConfig() *Config {
	mutex.RLock()
	defer mutex.RUnlock()
	return config
}
