package common

import (
	"fmt"
	"gorm.io/gorm"
)

var db *gorm.DB

// DBConfig represents db configuration
type dbConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func buildDBConfig(host, dbName, user, password string) *dbConfig {
	config := dbConfig{
		Host:     host,
		Port:     3306,
		DBName:   dbName,
		User:     user,
		Password: password,
	}
	return &config
}

func dbURL(config *dbConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)
}

func GetDb() *gorm.DB {
	return db
}

func ConnectDb(host, dbName, user, password string) {
	//connect, err := gorm.Open(mysql.Open(dbURL(buildDBConfig(host, dbName, user, password))), &gorm.Config{})
	//db = connect
	//if err != nil {
	//	fmt.Println("status: ", err)
	//}
}
