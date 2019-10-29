package models

import (
	"github.com/jinzhu/gorm"
)

// Model : you can create your own "gorm.Model" and embebed to use it
type Model interface {
	Save(db *gorm.DB) error
	DeleteOneByID(db *gorm.DB) error
	Update(db *gorm.DB) (Model, error)
	FindAll(db *gorm.DB) ([]Model, error)
	FindOneByID(db *gorm.DB) (Model, error)
}

// MyModel implements my custom attrs for all my models
type MyModel struct {
	CreatedAt string
	UpdatedAt string
	// it gets an error due to the internal casting by Gorm
	// DeletedAt *time.Time ` sql:"index"`
}
