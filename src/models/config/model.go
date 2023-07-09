package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// MySQL Config
type MySQLConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

// General Config
type Config struct {
	MySQL MySQLConfig `yaml:"mysql"`
}

func DefaultConfig() *Config {
	return &Config{
		MySQL: MySQLConfig{
			Host:     "localhost",
			Port:     3306,
			Username: "root",
			Password: "admin1234",
			Database: "gochat",
		},
	}
}

func GetConfig() *Config {
	cfg := DefaultConfig()

	// Read the config file
	file, err := os.ReadFile("src/config/local.yaml")
	if err != nil {
		log.Printf("Error al leer el archivo de configuracion: %v", err.Error())
		return cfg
	}

	// Parse the config file
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		log.Printf("Error al parsear el archivo de configuracion: %v", err)
		return cfg
	}

	return cfg
}
