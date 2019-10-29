package services

import (
	"fmt"
	"reflect"

	"github.com/PacoDw/introduction_to_GORM/db"
	"github.com/PacoDw/introduction_to_GORM/models"
	s "github.com/PacoDw/introduction_to_GORM/models/structs"
)

// DeleteByIDS deletes users by id
func DeleteByIDS(ids *s.IDs, m models.Model) s.Response {
	// Validate if the fields are correctly set
	if res := s.ValidateFields(ids); !reflect.DeepEqual(res, s.Response{}) {
		return res
	}

	// Get the database connection
	db := db.GetConnection()
	defer db.Close()

	/* We can't use this because it don't return errors/error
	 * but it looks nice */
	// errs := db.Where("id IN (?)", ids.IDs).Find(&users).Error
	// db.Where("id in (?)", idsFound).Delete(m).GetErrors()

	msn := []string{}
	existsOneRemoved := false
	for _, id := range ids.IDs {
		if err := db.Where("id = ?", id).Find(m).Error; err != nil {
			msn = append(msn, fmt.Sprintf("The id %s is not found", id))
		} else {
			if err := db.Where("id = ?", id).Delete(m).Error; err == nil {
				msn = append(msn, fmt.Sprintf("The id %s is removed", id))
				existsOneRemoved = true
			} else {
				msn = append(msn, err.Error())
			}
		}
	}

	if !existsOneRemoved {
		return s.Response{
			Success:  false,
			Messages: msn,
		}
	}

	return s.Response{
		Success:  true,
		Messages: msn,
	}
}
