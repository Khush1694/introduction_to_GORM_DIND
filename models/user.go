package models

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

// User Model
type User struct {
	// we need to use our own custom model because the id is different of the "gorm.Model"
	ID      string `gorm:"PRIMARY KEY; UNIQUE" json:"id"`
	Name    string `gorm:"type:varchar(255); NOT NULL" json:"name" validate:"required"`
	Email   string `gorm:"type:varchar(255)" json:"email"`
	Phone   string `gorm:"type:varchar(100); NOT NULL; UNIQUE; UNIQUE_INDEX" json:"phone" validate:"required"`
	Address string `gorm:"type:text" json:"address"`
	MyModel
}

// Check if Users Model implements/match Model interface
var _ Model = &User{}

// Save saves a User inside the database
func (u *User) Save(db *gorm.DB) error {
	uuidResult, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}

	u.ID = uuidResult.String()
	u.CreatedAt = time.Now().String()
	u.UpdatedAt = time.Now().String()

	err = db.Save(&u).Error
	if err != nil {
		return err
	}
	return nil
}

// FindAll gets all the users
func (u *User) FindAll(db *gorm.DB) ([]Model, error) {
	users := []*User{}
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	usersModel := []Model{}
	for _, user := range users {
		usersModel = append(usersModel, user)
	}

	return usersModel, nil
}

// FindOneByID gets a user by id
func (u *User) FindOneByID(db *gorm.DB) (Model, error) {
	err := db.Find(&u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

// Update updates the User and saves it
func (u *User) Update(db *gorm.DB) (Model, error) {
	var updatedUser User = *u

	_, err := u.FindOneByID(db)
	if err != nil {
		return nil, err
	}

	updatedUser.CreatedAt = u.CreatedAt
	updatedUser.UpdatedAt = time.Now().String()

	err = db.Save(&updatedUser).Error
	if err != nil {
		return nil, err
	}
	return &updatedUser, nil
}

// DeleteOneByID deletes the User by id
func (u *User) DeleteOneByID(db *gorm.DB) error {
	_, err := u.FindOneByID(db)
	if err != nil {
		return err
	}

	err = db.Delete(&u, "id = ? ", u.ID).Error
	if err != nil {
		return err
	}
	return nil
}
