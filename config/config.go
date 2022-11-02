package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cpatsonakis/goa-calc-example/helpers"
	"github.com/go-playground/validator/v10"
)

type Service interface {
	GetConfig() Config
}

type Config struct {
	SwaggerFile      string `json:"swagger_file" validate:"required,file"`
	ExternalEndpoint string `json:"external_endpoint" validate:"required"`
}

type configService struct {
	conf Config
}

func NewConfig(configFilePath string) (Service, error) {
	var err error
	if !helpers.FileExists(configFilePath) {
		return nil, fmt.Errorf("config file %s does not exist", configFilePath)
	}
	configFile, err := os.Open(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("os.Open() for config file returned error: %w", err)
	}
	var conf Config
	if err = json.NewDecoder(configFile).Decode(&conf); err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll() for config file returned error: %w", err)
	}
	validator := validator.New()
	if err = validator.Struct(conf); err != nil {
		return nil, fmt.Errorf("validator.Struct() for configuration returned error: %w", err)
	}
	return configService{
		conf: conf,
	}, nil
}

func (c configService) GetConfig() Config {
	return c.conf
}
