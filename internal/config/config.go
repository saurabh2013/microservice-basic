package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/saurabh2013/microservice-basic/internal/constant"
	"gopkg.in/yaml.v2"
)

var config struct {
	settings map[string]interface{}
}

const configDir = "../internal/config"

//Init load config settings for service
func Init(env string) error {
	// Open config file
	f := getConfigfiile(env)
	fmt.Println("config:", f)
	file, err := os.Open(f)
	if err != nil {
		return err
	}
	defer file.Close()

	// Start YAML decoding from file
	if err := yaml.NewDecoder(file).Decode(&config.settings); err != nil {
		return err
	}

	return nil
}

func getConfigfiile(env string) string {
	var file string
	switch env {
	case constant.Prod:
		file = "prod.yml"
	case constant.Dev:
		file = "dev.yml"
	default:
		file = "local.yml"
	}
	envConfigDir := os.Getenv(constant.ConfigDir)
	if envConfigDir != "" {
		return filepath.Join(envConfigDir, file)
	}
	return filepath.Join(configDir, file)
}

//Get func returns a given config entry
func Get(s string) (string, error) {
	v, err := getValue(s)
	if err != nil {
		return "", err
	}
	if val, k := v.(string); k {
		return val, nil
	}
	return "", fmt.Errorf("error, not a valid config string value %v for %s", v, s)
}

func getValue(s string) (interface{}, error) {
	if v, k := config.settings[s]; k {
		return v, nil
	}
	return nil, fmt.Errorf("error while getting config value for %v", s)
}
