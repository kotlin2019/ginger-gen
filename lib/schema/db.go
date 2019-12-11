package schema

import (
	"database/sql"
	"github.com/didi/gendry/manager"
	_ "github.com/go-sql-driver/mysql"
)

// DBConfig holds the basic configuration of database
type DBConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Port     int    `json:"port"`
	DBName   string `json:"database"`
}

func GetDBInstance(conf *DBConfig) (*sql.DB, error) {
	option := manager.New(conf.DBName, conf.User, conf.Password, conf.Host)
	return option.Port(conf.Port).Open(true)
}