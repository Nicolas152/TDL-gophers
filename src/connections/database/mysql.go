package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
	model := MySQLConnection {
		Host: "localhost",
		Port: 3306,
		Username: "yisus",
		Password: "46139282",
		Database: "gochat",
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