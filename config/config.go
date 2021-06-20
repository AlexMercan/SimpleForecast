package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	Location string `json:"location"`
}

func LoadConfiguration() (*Config, error) {
	configurationDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}
	//Create the config directory if it doesn't exist
	configFileDirectory := configurationDir + "/weatherapi"
	err = os.MkdirAll(configFileDirectory, 0755)
	if err != nil {
		return nil, err
	}

	configFilePath := configFileDirectory + "/config.json"
	//If the file doesn't exist, create it and write some default configuration to it
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		return createConfigAndWriteDefault(configFilePath)
	}

	configFile, err := os.Open(configFilePath)
	defer configFile.Close()
	if err != nil {
		return nil, err
	}

	fileContents, err := io.ReadAll(configFile)
	if err != nil {
		return nil, err
	}

	var configuration Config
	err = json.Unmarshal(fileContents, &configuration)
	if err != nil {
		return nil, err
	}
	return &configuration, nil
}

func createConfigAndWriteDefault(configFilePath string) (*Config, error) {
	configuration := Config{"auto:ip"}
	jsonConfigText, err := json.MarshalIndent(configuration, "", " ")
	if err != nil {
		return nil, err
	}
	err = os.WriteFile(configFilePath, jsonConfigText, 0666)
	if err != nil {
		return nil, err
	}
	return &configuration, nil
}
