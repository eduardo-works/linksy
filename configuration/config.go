package configuration

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/eduardo-works/linksy/utils"
)

const (
	httpPort = 8080
	dbHost   = "localhost"
	dbPort   = 3306
	dbName   = "linksy"
)

var fileName = "~/.linksy"

type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type HTTP struct {
	Port int
}

type Config struct {
	Database Database
	HTTP     HTTP
}

func LoadOrCreate() (Config, error) {
	config, err := loadFile()

	if errors.Is(err, os.ErrNotExist) {
		return createFile()
	}

	return config, err
}

func (c Config) Update() error {
	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}

	file, err := utils.ExpandPath(fileName)
	if err != nil {
		return err
	}

	return os.WriteFile(file, data, 0600)
}

func createFile() (Config, error) {
	config := Config{
		Database: Database{
			Host:     dbHost,
			Port:     dbPort,
			User:     "",
			Password: "",
			Name:     dbName,
		},
		HTTP: HTTP{
			Port: httpPort,
		},
	}
	return config, config.Update()
}

func loadFile() (Config, error) {
	file, err := utils.ExpandPath(fileName)
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(file)
	if err != nil {
		return Config{}, err
	}

	var config Config

	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
