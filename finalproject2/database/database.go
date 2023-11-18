package database

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
	"fmt"
	"github.com/adindazenn/kelompok6/finalproject2/model"
)

const (
	host		= "localhost"
	port		= 5432
	user		= "postgres"
	password	= "root"
	dbname		= "postgres"
)

func InitDB() (*gorm.DB, error) {
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    
	db.Debug().AutoMigrate(model.User{}, model.Category{}, model.Task{})
    
    return db, nil
}

