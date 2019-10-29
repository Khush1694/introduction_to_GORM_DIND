package services

import (
	"github.com/PacoDw/introduction_to_GORM/db"
	"github.com/PacoDw/introduction_to_GORM/models"
	s "github.com/PacoDw/introduction_to_GORM/models/structs"
)

// FindOneByID gets a mode by id
func FindOneByID(m models.Model) s.Response {
	// Get the database connection
	db := db.GetConnection()
	defer db.Close()

	m, err := m.FindOneByID(db)
	if err != nil {
		return s.Response{
			Success: false,
			Message: err.Error(),
		}
	}

	return s.Response{
		Success: true,
		Data:    m,
	}
}
