package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gochat/src/models/config"
	"log"
)

type MySQLConnection struct {
	Host string
	Port int
	Username string
	Password string
	Database string
}

func GetConnection() *sql.DB {
	// Obtengo la configuracion
	cfg := config.GetConfig()

	// Obtengo la configuracion de MySQL
	MySQLConfig := cfg.MySQL
	model := MySQLConnection {
		Host: MySQLConfig.Host,
		Port: MySQLConfig.Port,
		Username: MySQLConfig.Username,
		Password: MySQLConfig.Password,
		Database: MySQLConfig.Database,
	}

	return model.Connect()
}

func (mysql MySQLConnection) resolveDataSourceName() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", mysql.Username, mysql.Password, mysql.Host, mysql.Port, mysql.Database)
}

func (mysql MySQLConnection) Connect() *sql.DB {
	// Conecto a la base de datos
	db, err := sql.Open("mysql", mysql.resolveDataSourceName())
	if err != nil {
		log.Println("Error al conectarse a la base de datos: ", err)
	}
	return db
}