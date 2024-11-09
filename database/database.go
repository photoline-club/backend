package database

import (
	"fmt"

	"github.com/photoline-club/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

func (c *DBConfig) ConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DBName,
	)
}

func InitialiseDB(cfg DBConfig) *gorm.DB {
	fmt.Printf("cfg.ConnectionString(): %v\n", cfg.ConnectionString())
	db, err := gorm.Open(mysql.Open(cfg.ConnectionString()), &gorm.Config{})
	if err != nil {
		panic("Unable to connect to database")
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.FriendLink{})
	db.AutoMigrate(&models.Event{})
	db.AutoMigrate(&models.EventAsset{})
	db.AutoMigrate(&models.EventParticipant{})
	db.AutoMigrate(&models.Session{})
	return db
}
