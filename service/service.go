package service

import (
	"fmt"
	"log"
	"net/url"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	Services = "services"

	// Web Internal
	WebInternalServiceId = "web-internal"
)

var services = map[string]*Service{}

// Service definition
type Service struct {
	Id  string
	Url url.URL
}

// Lookup looks for the given service name
func LookupService(name string) (*Service, error) {

	if svc, ok := services[name]; ok {
		return svc, nil
	}
	return nil, fmt.Errorf("service '%s' does not exist in registry", name)

}

// Read and parse configuration file to extract services
func ReadServicesFromConfig() error {

	svcs := viper.GetStringMap(Services)

	for key := range svcs {

		// Not using UnmarshalKey because of its lack of environment variable override support
		sUrl := viper.GetString(fmt.Sprintf("%s.%s.url", Services, key))
		log.Printf("Service url %s", sUrl)

		u, err := url.Parse(sUrl)
		if err != nil {
			return fmt.Errorf("unable to parse '%s' service URL: %s", key, err)
		}

		s := Service{
			Id:  key,
			Url: *u,
		}
		services[key] = &s

		logrus.Debugf("read in config for service '%s' (url: %s)", key, sUrl)
	}

	return nil
}
