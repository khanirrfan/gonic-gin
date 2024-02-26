package database

import (
	"database/sql"
	"log"
	"time"

	"docxa/main.go/config"

	"github.com/allegro/bigcache"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	db          *gorm.DB // add gorm config
	GlobalCache *bigcache.BigCache
)

func OpenDBConnection() {
	var err error
	dsn := config.AppConfig.GetString("database.username") + ":" +
		config.AppConfig.GetString("database.password") + "@tcp(" +
		config.AppConfig.GetString("database.host") + ":" +
		config.AppConfig.GetString("database.port") + ")/" +
		config.AppConfig.GetString("database.name") +
		"?charset=utf8&parseTime=True&loc=Local"
	dbConnection, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Use the database connection to initialize GORM
	db, err = gorm.Open(mysql.New(mysql.Config{
		Conn: dbConnection,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Error initializing GORM:", err)
	}
	GlobalCache, err = bigcache.NewBigCache(bigcache.DefaultConfig(30 * time.Minute))
	if err != nil {
		log.Fatal("error connecting database", err.Error())
	}
}

func GetDb() *gorm.DB {
	return db
}
