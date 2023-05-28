package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// MySQL Config
type MySQLConfig struct {
	Host 		string	`yaml:"host"`
	Port 		int		`yaml:"port"`
	Username 	string	`yaml:"user"`
	Password 	string	`yaml:"password"`
	Database 	string	`yaml:"database"`
}

// General Config
type Config struct {
	MySQL 	MySQLConfig	`yaml:"mysql"`
}

func (cfg *Config) GetConfig() *Config {
	file, err := os.ReadFile("./src/config/local.yaml")
	if err != nil {
		log.Printf("Error al leer el archivo de configuracion: ", err)
	}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		log.Printf("Error al leer el archivo de configuracion: ", err)
	}

	return cfg
}
