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
	// TODO: Investigar si no hay otro forma de setear la configuracion por defecto
	cfg := DefaultConfig()

	// TODO: No es correcto que cada vez que se quiera acceder a la config, se tenga que leer el archivo
	// Leo el archivo de configuracion
	file, err := os.ReadFile("./src/config/local.yaml")
	if err != nil {
		log.Printf("Error al leer el archivo de configuracion: %v", err.Error())
	}

	// Parseo el archivo de configuracion
	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		log.Printf("Error al parsear el archivo de configuracion: %v", err)
	}

	return cfg
}
