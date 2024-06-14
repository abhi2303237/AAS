package utils

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2/log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

var lock = &sync.Mutex{}

type Config struct {
	Port            string `env:"PORT" envDefault:"8080"`
	ApplicationName string `env:"APPLICATION_NAME"`
	LogLevel        string `env:"LOG_LEVEL"`
	EnableMetricsUi bool   `env:"ENABLE_METRICS_UI" envDefault:"false"`
	Backend         string `env:"BACKEND" envDefault:"memmory"`

	MongoConnectionString string `env:"MONGO_CON_STRING"`

	DBServer   string `env:"MSSQL_SERVER"`
	DBPort     int    `env:"MSSQL_PORT"`
	DBUser     string `env:"MSSQL_USER"`
	DBPassword string `env:"MSSQL_PASS"`
	DBDatabase string `env:"MSSQL_DATABASE"`
}

var configInstance *Config

func GetConfig() *Config {
	if configInstance == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal(err)
		}
		lock.Lock()
		defer lock.Unlock()
		if configInstance == nil {
			config := Config{}
			opts := env.Options{RequiredIfNoDef: true}
			if err := env.ParseWithOptions(&config, opts); err != nil {
				fmt.Printf("%+v\n", err)
			}
			fmt.Println("Creating Config instance now.")
			configInstance = &config
		}
	}
	return configInstance
}

func GetTestInstance(path string) *Config {

	err := godotenv.Load(path)
	if err != nil {
		log.Fatal(err)
	}

	config := Config{}
	opts := env.Options{RequiredIfNoDef: true}
	if err := env.ParseWithOptions(&config, opts); err != nil {
		fmt.Printf("%+v\n", err)
	}
	return &config
}
