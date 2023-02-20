package repository

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"github.com/Jeswyrne/pig-latin/models"
)

type DatabaseInterface interface {
	Save(obj models.SaveObject) (models.SaveObject, error)
	List() ([]models.SaveObject, error)
}

type Database struct {
	DB *gorm.DB
}

var _ DatabaseInterface = &Database{}

func NewDatabase() DatabaseInterface {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to open the SQLite database.")
	}
	db.AutoMigrate(&models.SaveObject{})
	
	return &Database{DB: db}
}

func (db *Database) Save(obj models.SaveObject) (models.SaveObject, error) {
	var result models.SaveObject
	db.DB.Create(&models.SaveObject{
		Input:  obj.Input,
		Output: obj.Output,
	})
	db.DB.Last(&result)

	return result, nil
}

func (db *Database) List() ([]models.SaveObject, error) {
	var objects []models.SaveObject
	db.DB.Find(&objects)
	return objects, nil
}
