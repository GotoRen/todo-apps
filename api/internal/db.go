package internal

import (
	"fmt"

	"github.com/GotoRen/todo-apps/api/logger"
	"github.com/GotoRen/todo-apps/api/model"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConn() *gorm.DB {
	return DB
}

func DBClose() error {
	db, err := DBConn().DB()
	if err != nil {
		logger.LogErr("Failed to close database connection", "error", err)
	}

	return db.Close()
}

type dbConfig struct {
	Host     string `required:"true"`
	User     string `required:"true"`
	Password string `required:"true"`
	Database string `required:"true"`
}

// dbClient create a connection with the database.
func dbClient(dns string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// DBConnecter establish a database connection.
func DBConnecter() {
	dbConf := dbConfig{}
	err := envconfig.Process("DB", &dbConf)
	if err != nil {
		logger.LogErr("Failed to load environment variables", "error", err)
	} else {
		logger.LogDebug("Loaded environment variables", "Gorm v2", dbConf)
		EnvShow("DB_HOST")
		EnvShow("DB_DATABASE")
		EnvShow("DB_USER")
		EnvShow("DB_PASSWORD")
	}

	dns := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConf.User, dbConf.Password, dbConf.Host, dbConf.Database)
	db, err := dbClient(dns)
	if err != nil {
		logger.LogErr("Failed to create database connection", "error", err)
	} else {
		DB = db
		logger.LogDebug("Database connection established", "Gorm v2", DB)
	}
}

// DBMigrate migrate the database.
func DBMigrate() {
	if err := DBConn().AutoMigrate(&model.Todo{}); err != nil {
		logger.LogErr("Failed to database migration.", "error", err)
	} else {
		logger.LogDebug("Migrated the database", "Gorm v2", DB)
	}
}
