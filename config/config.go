package config

import (
	"os"
)

/* Default values of configuration */
const (
	DefaultEnv        = "Dev"
	DefaultPort       = ":8080"
	DefaultDBHost     = "vm.cgjr1jzt2ktg.ap-south-1.rds.amazonaws.com"
	DefaultDBPort     = "3306"
	DefaultDBUser     = "root"
	DefaultDBPassword = "password"
	DefaultDBName     = "instances"
)

type Config struct {
	Env         string
	Port        string
	DB_Host     string
	DB_Port     string
	DB_User     string
	DB_Password string
	DB_Name     string
}

var config Config

func init() {
	config.Env = os.Getenv("ENV")
	if config.Env == "" {
		config.Env = DefaultEnv
	}

	config.Port = ":" + os.Getenv("PORT")
	if config.Port == "" {
		config.Port = DefaultPort
	}

	config.DB_Host = os.Getenv("DB_HOST")
	if config.DB_Host == "" {
		config.DB_Host = DefaultDBHost
	}

	config.DB_Port = os.Getenv("DB_PORT")
	if config.DB_Port == "" {
		config.DB_Port = DefaultDBPort
	}

	config.DB_User = os.Getenv("DB_USER")
	if config.DB_User == "" {
		config.DB_User = DefaultDBUser
	}

	config.DB_Password = os.Getenv("DB_PASSWORD")
	if config.DB_Password == "" {
		config.DB_Password = DefaultDBPassword
	}

	config.DB_Name = os.Getenv("DB_NAME")
	if config.DB_Name == "" {
		config.DB_Name = DefaultDBName
	}
}

// Builds DB connection URI
// Example - "root:password@tcp(127.0.0.1:3306)/vm?charset=utf8mb4&parseTime=True"
func GetMysqlURI() string {
	ConnString :=
		config.DB_User +
			":" +
			config.DB_Password +
			"@tcp(" +
			config.DB_Host +
			":" +
			config.DB_Port +
			")/" +
			config.DB_Name +
			"?charset=utf8mb4&parseTime=True"

	return ConnString
}

//return configurartion
func Get() Config {
	return config
}
