package services

import (
	"reflect"

	"github.com/PacoDw/introduction_to_GORM/db"
	"github.com/PacoDw/introduction_to_GORM/models"
	s "github.com/PacoDw/introduction_to_GORM/models/structs"
)

// Update updates a model
func Update(m models.Model) s.Response {
	// Valdiate if the fields are correctly set
	if res := s.ValidateFields(m); !reflect.DeepEqual(res, s.Response{}) {
		return res
	}

	// Get the database connection
	db := db.GetConnection()
	defer db.Close()

	updatedModel, err := m.Update(db)
	if err != nil {
		return s.Response{
			Success: false,
			Message: err.Error(),
		}
	}

	return s.Response{
		Success: true,
		Data:    updatedModel,
	}
}
