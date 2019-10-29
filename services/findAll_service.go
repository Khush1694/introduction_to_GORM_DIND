package services

import (
	"github.com/PacoDw/introduction_to_GORM/db"
	"github.com/PacoDw/introduction_to_GORM/models"
	s "github.com/PacoDw/introduction_to_GORM/models/structs"
)

// FindAll gets all the found models
func FindAll(m models.Model) s.Response {
	// Get the database connection
	db := db.GetConnection()
	defer db.Close()

	ms, err := m.FindAll(db)
	if err != nil {
		return s.Response{
			Success: false,
			Message: err.Error(),
		}
	}

	return s.Response{
		Success: true,
		Data:    ms,
	}
}
