package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// ServiceConfig stores the whole configuration for service.
type ServiceConfig struct {
	Address     string `yaml:"address"`
	EndPoint    string `yaml:"endpoint"`
	ProgramPath string `yaml:"program_path"`
	DirPath     string `yaml:"dir_path"`
	FilePath    string `yaml:"file_path"`
}

func getRawConfig(fileLocation string) (*ServiceConfig, error) {
	configByte, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		return nil, err
	}
	config := &ServiceConfig{}
	err = yaml.Unmarshal(configByte, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

// GetServiceConfig parse the configuration from YAML file.
func (c *YAMLConfigLoader) GetServiceConfig() (*ServiceConfig, error) {
	config, err := getRawConfig(c.fileLocation)
	if err != nil {
		return nil, fmt.Errorf("Unable to get raw config content: %v", err)
	}
	return config, nil
}
