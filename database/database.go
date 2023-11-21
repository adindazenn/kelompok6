package database

import (
    "gorm.io/gorm"
    "gorm.io/driver/postgres"
	"fmt"
	"os"
	"github.com/adindazenn/kelompok6/model"
)

var (
	host		= os.Getenv("PGHOST")
	user		= os.Getenv("PGUSER")
	password	= os.Getenv("PGPASSWORD")
	port		= os.Getenv("PGPORT")
	dbname		= os.Getenv("PGDATABASE")
	db		= *gorm.DB
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

