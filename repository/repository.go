package repository

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	Repository = "mongodb"
)

// Database definition
type Database struct {
	Server string
	Name   string
}

var databases = map[string]*Database{}

func LookupDatabase(name string) (*Database, error) {

	if svc, ok := databases[name]; ok {
		return svc, nil
	}
	return nil, fmt.Errorf("database '%s' does not exist in registry", name)
}

func ReadDBFromConfig() error {

	svcs := viper.GetStringMap(Repository)

	for key := range svcs {

		dServer := viper.GetString(fmt.Sprintf("%s.%s.server", Repository, key))
		dName := viper.GetString(fmt.Sprintf("%s.%s.database", Repository, key))

		s := Database{
			Server: dServer,
			Name:   dName,
		}
		databases[key] = &s

		logrus.Debugf("Database '%s' (name: %s, url: %s) loaded", key, dName, dServer)
	}

	return nil
}
