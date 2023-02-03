package kitchen_sink

import (
	"os"

	"gopkg.in/yaml.v3"
)

func ReadConfig(config interface{}) error {
	configFilePath := os.Getenv("CONFIG_FILE_PATH")
	if configFilePath == "" {
		configFilePath = "config.yaml"
	}

	err := ReadConfigFile(configFilePath, config)
	if err != nil {
		return err
	}

	return nil
}

func ReadConfigFile(file string, config interface{}) error {
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return err
	}
	return nil
}

func WriteConfigFile(file string, config interface{}) error {
	d, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	err = os.WriteFile(file, d, 0644)
	if err != nil {
		return err
	}
	return nil
}
