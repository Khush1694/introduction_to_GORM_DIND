package services

import (
	"github.com/PacoDw/introduction_to_GORM/db"
	"github.com/PacoDw/introduction_to_GORM/models"
	s "github.com/PacoDw/introduction_to_GORM/models/structs"
)

// DeleteOneByID removes a model of the database
func DeleteOneByID(m models.Model) s.Response {
	// Get the database connection
	db := db.GetConnection()
	defer db.Close()

	err := m.DeleteOneByID(db)
	if err != nil {
		return s.Response{
			Success: false,
			Message: err.Error(),
		}
	}

	return s.Response{
		Success: true,
	}
}
