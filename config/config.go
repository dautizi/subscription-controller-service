package config

// Viper is helpful to configure by yaml
import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	. "github.com/Accedo-Products/subscription-controller-service/repository"
	. "github.com/Accedo-Products/subscription-controller-service/service"
)

const (
	ServerPort = "server.port"

	MongoDbURI          = "mongodb.uri"
	MongoDbDatabaseName = "mongodb.database-name"
)

// Database connection
type MongoConfig struct {
	Server   string
	Database string
}

// Read and parse the configuration file
func (c *MongoConfig) Read() {

	// Load database connection
	if _, err := toml.DecodeFile("config.toml", &c); err != nil {
		log.Fatal(err)
	}
}

// Initialise reads in the viper configuration file.
func Initialise() {

	viper.SetConfigName("application")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Debugf("error while reading in the configuration: %s", err)
	}

	// Load services
	ReadServicesFromConfig()

	// Load database config
	ReadDBFromConfig()
}
