package config

// YAMLConfigLoader is the loader of YAML file configuration.
type YAMLConfigLoader struct {
	fileLocation string
}

// NewYamlConfigLoader return the YAML Configuration loader.
func NewYamlConfigLoader(fileLocation string) *YAMLConfigLoader {
	return &YAMLConfigLoader{
		fileLocation: fileLocation,
	}
}
